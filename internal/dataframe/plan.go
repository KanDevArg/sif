package dataframe

import (
	"github.com/go-sif/sif"
	itypes "github.com/go-sif/sif/internal/types"
)

// planImpl is an optimized execution Plan for a DataFrame
type planImpl struct {
	stages []*stageImpl
	parser sif.DataSourceParser
	source sif.DataSource
}

// Size returns the number of stages in this Plan
func (p *planImpl) Size() int {
	return len(p.stages)
}

// GetStage returns a particular Stage in this Plan
func (p *planImpl) GetStage(idx int) itypes.Stage {
	return p.stages[idx]
}

// Parser returns this Plan's DataSourceParser
func (p *planImpl) Parser() sif.DataSourceParser {
	return p.parser
}

// Source returns this Plan's DataSource
func (p *planImpl) Source() sif.DataSource {
	return p.source
}

// Execute creates a planExecutor for this Plan
func (p *planImpl) Execute(conf *itypes.PlanExecutorConfig, statsTracker *itypes.RunStatistics) itypes.PlanExecutor {
	return CreatePlanExecutor(p, conf, statsTracker)
}
