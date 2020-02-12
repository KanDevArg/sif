package core

import (
	"log"
)

// A DataFrame is a factory for constructing a chain of
// transformations and actions applied to columnar data
type DataFrame struct {
	parent   *DataFrame       // the parent DataFrame. Nil if this is the root.
	task     Task             // the task represented by this DataFrame, executed to produce the next one
	taskType string           // a unique name for the type of task this DataFrame represents
	source   DataSource       // the source of the data
	parser   DataSourceParser // the parser for the source data
	schema   *Schema          // the current schema of the data at this task
}

// CreateDataFrame is a factory for DataFrames. This function is not intended to be used directly,
// as DataFrames are returned by DataSource packages.
func CreateDataFrame(source DataSource, parser DataSourceParser, schema *Schema) *DataFrame {
	return &DataFrame{
		parent:   nil,
		task:     &noOpTask{},
		taskType: "extract",
		source:   source,
		parser:   parser,
		schema:   schema,
	}
}

// GetParent returns the parent DataFrame of a DataFrame
func (df *DataFrame) GetParent() *DataFrame {
	return df.parent
}

// GetSchema returns the Schema of a DataFrame
func (df *DataFrame) GetSchema() *Schema {
	return df.schema
}

// GetDataSource returns the DataSource of a DataFrame
func (df *DataFrame) GetDataSource() DataSource {
	return df.source
}

// GetParser returns the DataSourceParser of a DataFrame
func (df *DataFrame) GetParser() DataSourceParser {
	return df.parser
}

// To is a "functional operations" factory method for DataFrames,
// chaining operations onto the current one(s).
func (df *DataFrame) To(ops ...DataFrameOperation) (*DataFrame, error) {
	next := df
	// See https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis for details of approach
	for _, op := range ops {
		nextTask, nextTaskType, nextSchema, err := op(next)
		if err != nil {
			return nil, err
		}
		next = &DataFrame{
			parent:   next,
			source:   df.source,
			task:     nextTask,
			taskType: nextTaskType,
			parser:   df.parser,
			schema:   nextSchema,
		}
	}
	return next, nil
}

// optimize splits the DataFrame chain into stages which each share a schema.
// Each stage's execution will be blocked until the completion of the previous stage
func (df *DataFrame) optimize() *plan {
	// create a slice of frames, in order of execution, by following parent links
	frames := []*DataFrame{}
	for next := df; next != nil; next = next.parent {
		frames = append([]*DataFrame{next}, frames...)
	}
	// split into stages at reductions and repacks, discovering incoming and outgoing schemas for the stage
	nextID := 0
	stages := []*stage{createStage(nextID)}
	nextID++
	for _, f := range frames {
		currentStage := stages[len(stages)-1]
		currentStage.frames = append(currentStage.frames, f)
		if len(stages) > 1 {
			currentStage.incomingSchema = stages[len(stages)-2].outgoingSchema
		}
		// the outgoing schema is always the last schema
		currentStage.outgoingSchema = f.schema
		// if this is a reduce, this is the end of the Stage
		if f.taskType == "reduce" {
			rTask, ok := f.task.(reductionTask)
			if !ok {
				log.Panicf("taskType is reduce but Task is not a reductionTask")
			}
			currentStage.setKeyingOperation(rTask.GetKeyingOperation())
			currentStage.setReductionOperation(rTask.GetReductionOperation())
			stages = append(stages, createStage(nextID))
			nextID++
		} else if f.taskType == "repack" {
			// repack should never be the first frame. Throw error if that is the case
			if len(currentStage.frames) == 0 {
				log.Panicf("Repack cannot be the first Task in a DataFrame")
			}
		} else if f.taskType == "collect" {
			break // no tasks can come after a collect
		}
	}
	return &plan{stages, df.parser, df.source}
}

// analyzeSource returns a PartitionMap for the source data for this DataFrame
func (df *DataFrame) analyzeSource() (PartitionMap, error) {
	return df.source.Analyze()
}

// workerExecuteTask runs this DataFrame's task against the previous Partition,
// returning the modified Partition (or a new one(s) if necessary).
// The previous Partition may be nil.
func (df *DataFrame) workerExecuteTask(previous *Partition) ([]*Partition, error) {
	res, err := df.task.RunWorker(previous)
	if err != nil {
		return nil, err
	}
	previous.currentSchema = df.schema // update current schema
	return res, err
}
