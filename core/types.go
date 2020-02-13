package core

import pb "github.com/go-sif/sif/core/internal/rpc"

// A DataFrame is a tool for constructing a chain of
// transformations and actions applied to columnar data
type DataFrame interface {
	GetSchema() *Schema                          // GetSchema returns the Schema of a DataFrame
	GetDataSource() DataSource                   // GetDataSource returns the DataSource of a DataFrame
	GetParser() DataSourceParser                 // GetParser returns the DataSourceParser of a DataFrame
	To(...DataFrameOperation) (DataFrame, error) // To is a "functional operations" factory method for DataFrames, chaining operations onto the current one(s).
}

// An executableDataFrame adds methods specific to cluster execution of DataFrames
type executableDataFrame interface {
	DataFrame
	getParent() DataFrame                                                  // getParent returns the parent DataFrame of a DataFrame
	optimize() *plan                                                       // optimize splits the DataFrame chain into stages which each share a schema. Each stage's execution will be blocked until the completion of the previous stage
	analyzeSource() (PartitionMap, error)                                  // analyzeSource returns a PartitionMap for the source data for this DataFrame
	workerExecuteTask(previous OperablePTition) ([]OperablePTition, error) // workerExecuteTask runs this DataFrame's task against the previous Partition, returning the modified Partition (or a new one(s) if necessary). The previous Partition may be nil.
}

// A PTition is a portion of a columnar dataset, consisting of multiple Rows.
// Partitions are not generally interacted with directly, instead being
// manipulated in parallel by DataFrame Tasks.
type PTition interface {
	ID() string             // ID retrieves the ID of this Partition
	GetMaxRows() int        // GetMaxRows retrieves the maximum number of rows in this Partition
	GetNumRows() int        // GetNumRows retrieves the number of rows in this Partition
	GetRow(rowNum int) *Row // GetRow retrieves a specific row from this Partition
}

// A BuildablePTition can be built
type BuildablePTition interface {
	PTition
	canInsertRowData(row []byte) error                                                                                                             // canInsertRowData checks if a Row can be inserted into this Partition
	AppendEmptyRowData() (*Row, error)                                                                                                             // AppendEmptyRowData is a convenient way to add an empty Row to the end of this Partition, returning the Row so that Row methods can be used to populate it
	AppendRowData(row []byte, meta []byte, varData map[string]interface{}, serializedVarRowData map[string][]byte) error                           // AppendRowData adds a Row to the end of this Partition, if it isn't full and if the Row fits within the schema
	appendKeyedRowData(row []byte, meta []byte, varData map[string]interface{}, serializedVarRowData map[string][]byte, key uint64) error          // appendKeyedRowData appends a keyed Row to the end of this Partition
	InsertRowData(row []byte, meta []byte, varRowData map[string]interface{}, serializedVarRowData map[string][]byte, pos int) error               // InsertRowData inserts a Row at a specific position within this Partition, if it isn't full and if the Row fits within the schema. Other Rows are shifted as necessary.
	insertKeyedRowData(row []byte, meta []byte, varData map[string]interface{}, serializedVarRowData map[string][]byte, key uint64, pos int) error // insertKeyedRowData inserts a keyed Row into this Partition
}

// A CloneablePTition is cloneable
type CloneablePTition interface {
	getRowMeta(rowNum int) []byte                         // getRowMeta retrieves specific row metadata from this Partition
	getRowMetaRange(start int, end int) []byte            // getRowMetaRange retrieves an arbitrary range of bytes from the row meta
	getRowData(rowNum int) []byte                         // getRowData retrieves a specific row from this Partition
	getRowDataRange(start int, end int) []byte            // getRowDataRange retrieves an arbitrary range of bytes from the row data
	getVarRowData(rowNum int) map[string]interface{}      // getVarRowData retrieves the variable-length data for a given row from this Partition
	getSerializedVarRowData(rowNum int) map[string][]byte // getSerializedVarRowData retrieves the serialized variable-length data for a given row from this Partition
	getCurrentSchema() *Schema                            // getCurrentSchema retrieves the Schema from the most recent task that manipulated this Partition
	getWidestSchema() *Schema                             // getWidestSchema retrieves the widest Schema from the stage that produced this Partition, which is equal to the size of a row
	getIsKeyed() bool                                     // GetIsKeyed returns true iff this Partition has been keyed with KeyRows
	getKey(rowNum int) (uint64, error)                    // getKey returns the shuffle key for a row, as generated by KeyRows
	getKeyRange(rowNum int, numRows int) []uint64         // getKeyRange returns a range of shuffle keys for a row, as generated by KeyRows, starting at rowNum
}

// A KeyablePTition can be keyed
type KeyablePTition interface {
	KeyRows(kfn KeyingOperation) (OperablePTition, error) // KeyRows generates hash keys for a row from a key column. Attempts to manipulate partition in-place, falling back to creating a fresh partition if there are row errors
}

// A TreeablePTition can be stored in a pTree
type TreeablePTition interface {
	BuildablePTition
	CloneablePTition
	KeyablePTition
	buildRow(idx int) *Row                                                         // buildRow constructs a Row from this Partition using the given index
	findFirstKey(key uint64) (int, error)                                          // PRECONDITION: Partition must already be sorted by key
	findFirstRowKey(keyBuf []byte, key uint64, keyfn KeyingOperation) (int, error) // PRECONDITION: Partition must already be sorted by key
	averageKeyValue() (uint64, error)                                              // value of key within this sorted, keyed Partition
	split(pos int) (TreeablePTition, TreeablePTition, error)                       // split splits a Partition into two Partitions. Split position ends up in right Partition.
	balancedSplit() (uint64, TreeablePTition, TreeablePTition, error)              // Split position ends up in right Partition.
	toBytes() ([]byte, error)                                                      // toBytes serializes a Partition to a byte array suitable for persistance to disk
}

// An TransferrablePTition can be transferred and cloned
type TransferrablePTition interface {
	BuildablePTition
	CloneablePTition
	toMetaMessage() *pb.MPartitionMeta                                                                               // toMetaMessage serializes metadata about this Partition to a protobuf message
	receiveStreamedData(stream pb.PartitionsService_TransferPartitionDataClient, incomingSchema *Schema) error       // receiveStreamedData loads data from a protobuf stream into this Partition
	partitionFromMetaMessage(m *pb.MPartitionMeta, widestSchema *Schema, currentSchema *Schema) TransferrablePTition // PartitionFromMetaMessage deserializes a Partition from a protobuf message
}

// An OperablePTition can be operated on
type OperablePTition interface {
	PTition
	KeyablePTition
	UpdateCurrentSchema(currentSchema *Schema)                  // Sets the current schema of a Partition
	MapRows(fn MapOperation) (OperablePTition, error)           // MapRows runs a MapOperation on each row in this Partition, manipulating them in-place. Will fall back to creating a fresh partition if PartitionRowErrors occur.
	FlatMapRows(fn FlatMapOperation) ([]OperablePTition, error) // FlatMapRows runs a FlatMapOperation on each row in this Partition, creating new Partitions
	FilterRows(fn FilterOperation) (OperablePTition, error)     // FilterRows filters the Rows in the current Partition, creating a new one
	Repack(newSchema *Schema) (OperablePTition, error)          // Repack repacks a Partition according to a new Schema
}

// A CollectedPTition has been collected
type CollectedPTition interface {
	PTition
	MapRows(fn MapOperation) (OperablePTition, error) // MapRows runs a MapOperation on each row in this Partition, manipulating them in-place. Will fall back to creating a fresh partition if PartitionRowErrors occur.
}
