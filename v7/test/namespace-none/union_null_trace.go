// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
)

type UnionNullTraceTypeEnum int

const (
	UnionNullTraceTypeEnumTrace UnionNullTraceTypeEnum = 1
)

type UnionNullTrace struct {
	Null      *types.NullVal
	Trace     *Trace
	UnionType UnionNullTraceTypeEnum
}

func writeUnionNullTrace(r *UnionNullTrace, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullTraceTypeEnumTrace:
		return writeTrace(r.Trace, w)
	}
	return fmt.Errorf("invalid value for *UnionNullTrace")
}

func NewUnionNullTrace() *UnionNullTrace {
	return &UnionNullTrace{}
}

func (_ *UnionNullTrace) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullTrace) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullTrace) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullTrace) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullTrace) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullTrace) SetString(v string)  { panic("Unsupported operation") }
func (r *UnionNullTrace) SetLong(v int64) {
	r.UnionType = (UnionNullTraceTypeEnum)(v)
}
func (r *UnionNullTrace) Get(i int) types.Field {
	switch i {
	case 0:
		return r.Null
	case 1:
		r.Trace = NewTrace()
		return r.Trace
	}
	panic("Unknown field index")
}
func (_ *UnionNullTrace) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullTrace) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullTrace) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullTrace) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullTrace) Finalize()                        {}

func (r *UnionNullTrace) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	switch r.UnionType {
	case UnionNullTraceTypeEnumTrace:
		return json.Marshal(map[string]interface{}{"Trace": r.Trace})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullTrace")
}

func (r *UnionNullTrace) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if value, ok := fields["Trace"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Trace)
	}
	return fmt.Errorf("invalid value for *UnionNullTrace")
}
