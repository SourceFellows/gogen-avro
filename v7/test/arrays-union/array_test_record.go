// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     arrays.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

type ArrayTestRecord struct {
	IntField []*UnionNullInt `json:"IntField"`
}

const ArrayTestRecordAvroCRC64Fingerprint = "t\x06\x9e\xc8\u0088\xa0\xcb"

func NewArrayTestRecord() *ArrayTestRecord {
	return &ArrayTestRecord{}
}

func DeserializeArrayTestRecord(r io.Reader) (*ArrayTestRecord, error) {
	t := NewArrayTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeArrayTestRecordFromSchema(r io.Reader, schema string) (*ArrayTestRecord, error) {
	t := NewArrayTestRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writeArrayTestRecord(r *ArrayTestRecord, w io.Writer) error {
	var err error
	err = writeArrayUnionNullInt(r.IntField, w)
	if err != nil {
		return err
	}
	return err
}

func (r *ArrayTestRecord) Serialize(w io.Writer) error {
	return writeArrayTestRecord(r, w)
}

func (r *ArrayTestRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"IntField\",\"type\":{\"items\":[\"null\",\"int\"],\"type\":\"array\"}}],\"name\":\"ArrayTestRecord\",\"type\":\"record\"}"
}

func (r *ArrayTestRecord) SchemaName() string {
	return "ArrayTestRecord"
}

func (_ *ArrayTestRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *ArrayTestRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *ArrayTestRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *ArrayTestRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *ArrayTestRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *ArrayTestRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *ArrayTestRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ *ArrayTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ArrayTestRecord) Get(i int) types.Field {
	switch i {
	case 0:
		r.IntField = make([]*UnionNullInt, 0)

		return &ArrayUnionNullIntWrapper{Target: &r.IntField}
	}
	panic("Unknown field index")
}

func (r *ArrayTestRecord) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ArrayTestRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *ArrayTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *ArrayTestRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *ArrayTestRecord) Finalize()                        {}

func (_ *ArrayTestRecord) AvroCRC64Fingerprint() []byte {
	return []byte(ArrayTestRecordAvroCRC64Fingerprint)
}
