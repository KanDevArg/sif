package core

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"math"
	"testing"
	"time"

	types "github.com/go-sif/sif/v0.0.1/columntype"
	"github.com/stretchr/testify/require"
)

func TestGetUint64(t *testing.T) {
	row := Row{
		schema: &Schema{
			schema: map[string]*column{
				"col1": &column{0, 0, &types.Uint64ColumnType{}},
			},
		},
		data: make([]byte, 16),
		meta: make([]byte, 1),
	}
	binary.LittleEndian.PutUint64(row.data, math.MaxUint64)
	data, _ := row.GetUint64("col1")
	if data != math.MaxUint64 {
		t.FailNow()
	}
}

func TestTime(t *testing.T) {
	row := Row{
		schema: &Schema{
			schema: map[string]*column{
				"col1": &column{0, 0, &types.TimeColumnType{}},
			},
		},
		data: make([]byte, 15),
		meta: make([]byte, 1),
	}
	v := time.Now()
	err := row.SetTime("col1", v)
	require.Nil(t, err)
	v2, err := row.GetTime("col1")
	require.Nil(t, err)
	require.EqualValues(t, v.UnixNano(), v2.UnixNano())
}

func TestDeserialization(t *testing.T) {
	// When partition is transferred over a network, all variable-length data is Gob-encoded and deserialized on-demand on the other side.
	serialized := make(map[string][]byte)
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)
	err := e.Encode("world")
	serialized["hello"] = b.Bytes()
	require.Nil(t, err)
	row := Row{
		schema: &Schema{
			schema: map[string]*column{
				"hello": &column{0, 0, &types.VarStringColumnType{}},
			},
		},
		data:              make([]byte, 16),
		varData:           make(map[string]interface{}),
		serializedVarData: serialized,
		meta:              make([]byte, 1),
	}
	val, err := row.GetVarString("hello")
	require.Nil(t, err)
	require.Equal(t, "world", val)
}
