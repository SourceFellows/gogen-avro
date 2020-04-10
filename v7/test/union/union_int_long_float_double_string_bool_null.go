// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     union.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
)

type UnionIntLongFloatDoubleStringBoolNullTypeEnum int

const (
	UnionIntLongFloatDoubleStringBoolNullTypeEnumInt UnionIntLongFloatDoubleStringBoolNullTypeEnum = 0

	UnionIntLongFloatDoubleStringBoolNullTypeEnumLong UnionIntLongFloatDoubleStringBoolNullTypeEnum = 1

	UnionIntLongFloatDoubleStringBoolNullTypeEnumFloat UnionIntLongFloatDoubleStringBoolNullTypeEnum = 2

	UnionIntLongFloatDoubleStringBoolNullTypeEnumDouble UnionIntLongFloatDoubleStringBoolNullTypeEnum = 3

	UnionIntLongFloatDoubleStringBoolNullTypeEnumString UnionIntLongFloatDoubleStringBoolNullTypeEnum = 4

	UnionIntLongFloatDoubleStringBoolNullTypeEnumBool UnionIntLongFloatDoubleStringBoolNullTypeEnum = 5
)

type UnionIntLongFloatDoubleStringBoolNull struct {
	Int       int32
	Long      int64
	Float     float32
	Double    float64
	String    string
	Bool      bool
	Null      *types.NullVal
	UnionType UnionIntLongFloatDoubleStringBoolNullTypeEnum
}

func writeUnionIntLongFloatDoubleStringBoolNull(r *UnionIntLongFloatDoubleStringBoolNull, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(6, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionIntLongFloatDoubleStringBoolNullTypeEnumInt:
		return vm.WriteInt(r.Int, w)
	case UnionIntLongFloatDoubleStringBoolNullTypeEnumLong:
		return vm.WriteLong(r.Long, w)
	case UnionIntLongFloatDoubleStringBoolNullTypeEnumFloat:
		return vm.WriteFloat(r.Float, w)
	case UnionIntLongFloatDoubleStringBoolNullTypeEnumDouble:
		return vm.WriteDouble(r.Double, w)
	case UnionIntLongFloatDoubleStringBoolNullTypeEnumString:
		return vm.WriteString(r.String, w)
	case UnionIntLongFloatDoubleStringBoolNullTypeEnumBool:
		return vm.WriteBool(r.Bool, w)
	}
	return fmt.Errorf("invalid value for *UnionIntLongFloatDoubleStringBoolNull")
}

func NewUnionIntLongFloatDoubleStringBoolNull() *UnionIntLongFloatDoubleStringBoolNull {
	return &UnionIntLongFloatDoubleStringBoolNull{}
}

func (_ *UnionIntLongFloatDoubleStringBoolNull) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionIntLongFloatDoubleStringBoolNull) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionIntLongFloatDoubleStringBoolNull) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionIntLongFloatDoubleStringBoolNull) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionIntLongFloatDoubleStringBoolNull) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionIntLongFloatDoubleStringBoolNull) SetString(v string)  { panic("Unsupported operation") }
func (r *UnionIntLongFloatDoubleStringBoolNull) SetLong(v int64) {
	r.UnionType = (UnionIntLongFloatDoubleStringBoolNullTypeEnum)(v)
}
func (r *UnionIntLongFloatDoubleStringBoolNull) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.Int{Target: (&r.Int)}
	case 1:
		return &types.Long{Target: (&r.Long)}
	case 2:
		return &types.Float{Target: (&r.Float)}
	case 3:
		return &types.Double{Target: (&r.Double)}
	case 4:
		return &types.String{Target: (&r.String)}
	case 5:
		return &types.Boolean{Target: (&r.Bool)}
	case 6:
		return r.Null
	}
	panic("Unknown field index")
}
func (_ *UnionIntLongFloatDoubleStringBoolNull) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionIntLongFloatDoubleStringBoolNull) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionIntLongFloatDoubleStringBoolNull) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionIntLongFloatDoubleStringBoolNull) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ *UnionIntLongFloatDoubleStringBoolNull) Finalize() {}

func (r *UnionIntLongFloatDoubleStringBoolNull) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	switch r.UnionType {
	case UnionIntLongFloatDoubleStringBoolNullTypeEnumInt:
		return json.Marshal(map[string]interface{}{"int": r.Int})
	case UnionIntLongFloatDoubleStringBoolNullTypeEnumLong:
		return json.Marshal(map[string]interface{}{"long": r.Long})
	case UnionIntLongFloatDoubleStringBoolNullTypeEnumFloat:
		return json.Marshal(map[string]interface{}{"float": r.Float})
	case UnionIntLongFloatDoubleStringBoolNullTypeEnumDouble:
		return json.Marshal(map[string]interface{}{"double": r.Double})
	case UnionIntLongFloatDoubleStringBoolNullTypeEnumString:
		return json.Marshal(map[string]interface{}{"string": r.String})
	case UnionIntLongFloatDoubleStringBoolNullTypeEnumBool:
		return json.Marshal(map[string]interface{}{"boolean": r.Bool})
	}
	return nil, fmt.Errorf("invalid value for *UnionIntLongFloatDoubleStringBoolNull")
}

func (r *UnionIntLongFloatDoubleStringBoolNull) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if value, ok := fields["int"]; ok {
		r.UnionType = 0
		return json.Unmarshal([]byte(value), &r.Int)
	}
	if value, ok := fields["long"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Long)
	}
	if value, ok := fields["float"]; ok {
		r.UnionType = 2
		return json.Unmarshal([]byte(value), &r.Float)
	}
	if value, ok := fields["double"]; ok {
		r.UnionType = 3
		return json.Unmarshal([]byte(value), &r.Double)
	}
	if value, ok := fields["string"]; ok {
		r.UnionType = 4
		return json.Unmarshal([]byte(value), &r.String)
	}
	if value, ok := fields["boolean"]; ok {
		r.UnionType = 5
		return json.Unmarshal([]byte(value), &r.Bool)
	}
	return fmt.Errorf("invalid value for *UnionIntLongFloatDoubleStringBoolNull")
}
