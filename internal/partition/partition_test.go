package partition

import (
	"testing"

	"github.com/go-sif/sif"
	errors "github.com/go-sif/sif/errors"
	"github.com/go-sif/sif/schema"
	"github.com/stretchr/testify/require"
)

func createPartitionTestSchema() sif.Schema {
	schema := schema.CreateSchema()
	schema.CreateColumn("col1", &sif.Uint8ColumnType{})
	return schema
}

func TestCreatePartitionImpl(t *testing.T) {
	schema := createPartitionTestSchema()
	part := createPartitionImpl(4, schema, schema)
	require.Equal(t, part.GetMaxRows(), 4)
	require.Equal(t, part.GetNumRows(), 0)
	require.Nil(t, part.CanInsertRowData(make([]byte, 1)))
	require.NotNil(t, part.CanInsertRowData(make([]byte, 4)))
	require.False(t, part.GetIsKeyed())
}

func TestAppendRowData(t *testing.T) {
	// make partition
	schema := createPartitionTestSchema()
	part := createPartitionImpl(4, schema, schema)
	require.Equal(t, part.GetNumRows(), 0)
	r := []byte{byte(uint8(1))}
	// append and validate row
	err := part.AppendRowData(r, []byte{0}, make(map[string]interface{}), make(map[string][]byte))
	require.Nil(t, err)
	require.Equal(t, part.GetNumRows(), 1)
	val, err := part.GetRow(0).GetUint8("col1")
	require.Nil(t, err)
	require.Equal(t, val, uint8(1))
	// append and validate another row
	r = []byte{byte(uint8(2))}
	err = part.AppendRowData(r, []byte{0}, make(map[string]interface{}), make(map[string][]byte))
	require.Nil(t, err)
	require.Equal(t, part.GetNumRows(), 2)
	val, err = part.GetRow(1).GetUint8("col1")
	require.Nil(t, err)
	require.Equal(t, val, uint8(2))
}

func TestInsertRowData(t *testing.T) {
	// create partition
	schema := createPartitionTestSchema()
	part := createPartitionImpl(4, schema, schema)
	require.Equal(t, part.GetNumRows(), 0)
	// append and validate row
	r := []byte{byte(uint8(1))}
	err := part.AppendRowData(r, []byte{0}, make(map[string]interface{}), make(map[string][]byte))
	require.Nil(t, err)
	require.Equal(t, part.GetNumRows(), 1)
	val, err := part.GetRow(0).GetUint8("col1")
	require.Nil(t, err)
	require.Equal(t, val, uint8(1))
	// insert and validate row
	r = []byte{byte(uint8(2))}
	err = part.InsertRowData(r, []byte{0}, make(map[string]interface{}), make(map[string][]byte), 0)
	require.Nil(t, err)
	require.Equal(t, part.GetNumRows(), 2)
	val, err = part.GetRow(0).GetUint8("col1")
	require.Nil(t, err)
	require.Equal(t, val, uint8(2))
}

func TestPartitionFullError(t *testing.T) {
	// create partition with max 1 row
	schema := createPartitionTestSchema()
	part := createPartitionImpl(1, schema, schema)
	require.Equal(t, part.GetNumRows(), 0)
	// append and validate row
	r := []byte{byte(uint8(1))}
	err := part.AppendRowData(r, []byte{0}, make(map[string]interface{}), make(map[string][]byte))
	require.Nil(t, err)
	require.Equal(t, part.GetNumRows(), 1)
	val, err := part.GetRow(0).GetUint8("col1")
	require.Nil(t, err)
	require.Equal(t, val, uint8(1))
	// attempt to append row again
	err = part.AppendRowData(r, []byte{0}, make(map[string]interface{}), make(map[string][]byte))
	require.NotNil(t, err)
	_, ok := err.(errors.PartitionFullError)
	require.True(t, ok)
}

func TestIncompatibleRowError(t *testing.T) {
	// create partition with max 1 row
	schema := createPartitionTestSchema()
	part := createPartitionImpl(1, schema, schema)
	require.Equal(t, part.GetNumRows(), 0)
	// append and validate row
	r := []byte{byte(uint8(1))}
	err := part.AppendRowData(r, []byte{0}, make(map[string]interface{}), make(map[string][]byte))
	require.Nil(t, err)
	require.Equal(t, part.GetNumRows(), 1)
	val, err := part.GetRow(0).GetUint8("col1")
	require.Nil(t, err)
	require.Equal(t, val, uint8(1))
	// attempt to append incompatible row
	r = []byte{byte(uint8(1)), byte(uint8(2))}
	err = part.AppendRowData(r, []byte{0}, make(map[string]interface{}), make(map[string][]byte))
	require.NotNil(t, err)
	_, ok := err.(errors.IncompatibleRowError)
	require.True(t, ok)
}

func TestMapRows(t *testing.T) {
	// create partition
	schema := createPartitionTestSchema()
	part := createPartitionImpl(4, schema, schema)
	require.Equal(t, part.GetNumRows(), 0)
	// append rows
	for i := 0; i < 4; i++ {
		r := []byte{byte(uint8(i))}
		err := part.AppendRowData(r, []byte{0}, make(map[string]interface{}), make(map[string][]byte))
		require.Nil(t, err)
	}
	sum := 0
	_, err := part.MapRows(func(row sif.Row) error {
		val, err := row.GetUint8("col1")
		sum += int(val)
		return err
	})
	require.Nil(t, err)
	require.Equal(t, sum, 6)
}

func TestKeyRows(t *testing.T) {
	// create partition
	schema := createPartitionTestSchema()
	part := createPartitionImpl(8, schema, schema)
	// append rows
	for i := 0; i < 7; i++ {
		r := []byte{uint8(i)}
		meta := []byte{0}
		err := part.AppendRowData(r, meta, make(map[string]interface{}), make(map[string][]byte))
		require.Nil(t, err)
	}
	// add in a single duplicate row for good measure.
	err := part.AppendRowData([]byte{6}, []byte{0}, make(map[string]interface{}), make(map[string][]byte))
	// shouldn't be able to get keys before we key a partition
	_, err = part.GetKey(0)
	require.NotNil(t, err)
	// key rows
	_, err = part.KeyRows(func(row sif.Row) ([]byte, error) {
		val, err := row.GetUint8("col1")
		if err != nil {
			return nil, err
		}
		return []byte{byte(val)}, nil
	})
	require.Nil(t, err)
	require.True(t, part.GetIsKeyed())
	// compare keys for identical rows
	key1, err := part.GetKey(6)
	require.Nil(t, err)
	key2, err := part.GetKey(7)
	require.Nil(t, err)
	require.EqualValues(t, key1, key2)
	// even though the key appears twice, FindFirstKey should always return the first occurrence
	idx, err := part.FindFirstKey(key2)
	require.Nil(t, err)
	require.Equal(t, 6, idx)
	// keys that don't exist
	_, err = part.FindFirstKey(uint64(1234))
	require.NotNil(t, err)
}

func TestSplit(t *testing.T) {
	// create partition
	schema := createPartitionTestSchema()
	part := createPartitionImpl(8, schema, schema)
	// append rows
	for i := 0; i < 8; i++ {
		r := []byte{byte(uint8(i))}
		err := part.AppendRowData(r, []byte{0}, make(map[string]interface{}), make(map[string][]byte))
		require.Nil(t, err)
	}
	left, right, err := part.Split(4)
	require.Nil(t, err)
	// verify values
	val, err := left.GetRow(0).GetUint8("col1")
	require.Nil(t, err)
	require.Equal(t, val, uint8(0))
	val, err = right.GetRow(0).GetUint8("col1")
	require.Nil(t, err)
	require.Equal(t, val, uint8(4))
	// key rows
	_, err = part.KeyRows(func(row sif.Row) ([]byte, error) {
		val, err := row.GetUint8("col1")
		if err != nil {
			return nil, err
		}
		return []byte{byte(val)}, nil
	})
	// split again and verify keys
	left, right, err = part.Split(4)
	key, err := part.GetKey(0)
	require.Nil(t, err)
	lkey, err := left.GetKey(0)
	require.Nil(t, err)
	require.Equal(t, key, lkey)
	key, err = part.GetKey(4)
	require.Nil(t, err)
	rkey, err := right.GetKey(0)
	require.Nil(t, err)
	require.Equal(t, rkey, key)
}

func TestSerialization(t *testing.T) {
	// TODO test pre-serialized var row data
	// create partition
	schema := createPartitionTestSchema()
	schema.CreateColumn("col2", &sif.VarStringColumnType{})
	part := createPartitionImpl(8, schema, schema)
	// append rows
	for i := 0; i < 8; i++ {
		r := []byte{byte(uint8(i))}
		vr := make(map[string]interface{})
		vr["col2"] = "Hello World"
		err := part.AppendRowData(r, []byte{0}, vr, make(map[string][]byte))
		require.Nil(t, err)
	}
	// verify values
	require.Equal(t, 8, part.GetNumRows())
	for i := 0; i < 8; i++ {
		val1, err := part.GetRow(i).GetUint8("col1")
		require.Nil(t, err)
		require.Equal(t, val1, uint8(i))
		val2, err := part.GetRow(i).GetVarString("col2")
		require.Nil(t, err)
		require.Equal(t, val2, "Hello World")
	}
	// serialize and deserialize
	buff, err := part.ToBytes()
	require.Nil(t, err)
	rpart, err := FromBytes(buff, part.widestSchema, part.currentSchema)
	require.Nil(t, err)
	// verify values again
	require.Equal(t, 8, rpart.GetNumRows())
	for i := 0; i < 8; i++ {
		val1, err := rpart.GetRow(i).GetUint8("col1")
		require.Nil(t, err)
		require.Equal(t, val1, uint8(i))
		val2, err := rpart.GetRow(i).GetVarString("col2")
		require.Nil(t, err)
		require.Equal(t, val2, "Hello World")
	}
}

func TestRepack(t *testing.T) {
	// create partition
	schema := createPartitionTestSchema()
	schema.CreateColumn("col2", &sif.Float64ColumnType{})
	schema.CreateColumn("col3", &sif.VarStringColumnType{})
	part := createPartitionImpl(8, schema, schema)
	// append rows
	for i := 0; i < 8; i++ {
		row, err := part.AppendEmptyRowData()
		require.Nil(t, err)
		row.SetInt8("col1", int8(i))
		row.SetFloat64("col2", float64(i+1))
		row.SetVarString("col3", "Hello World")
	}
	// verify values
	require.Equal(t, 8, part.GetNumRows())
	for i := 0; i < 8; i++ {
		val1, err := part.GetRow(i).GetUint8("col1")
		require.Nil(t, err)
		require.Equal(t, val1, uint8(i))
		val2, err := part.GetRow(i).GetFloat64("col2")
		require.Nil(t, err)
		require.Equal(t, val2, float64(i+1))
		val3, err := part.GetRow(i).GetVarString("col3")
		require.Nil(t, err)
		require.Equal(t, val3, "Hello World")
	}
	// test repack
	newSchema := schema.Clone()
	newSchema.RemoveColumn("col1")
	newPart, err := part.Repack(newSchema)
	require.Nil(t, err)
	require.Equal(t, part.GetNumRows(), newPart.GetNumRows())
	require.Equal(t, part.GetMaxRows(), newPart.GetMaxRows())
	for i := 0; i < 8; i++ {
		origRow := part.GetRow(i)
		newRow := newPart.GetRow(i)
		val2, err := origRow.GetFloat64("col2")
		require.Nil(t, err)
		newVal2, err := newRow.GetFloat64("col2")
		require.Nil(t, err)
		require.Equal(t, val2, newVal2)
		val3, err := origRow.GetVarString("col3")
		require.Nil(t, err)
		newVal3, err := newRow.GetVarString("col3")
		require.Nil(t, err)
		require.Equal(t, val3, newVal3)
	}
}
