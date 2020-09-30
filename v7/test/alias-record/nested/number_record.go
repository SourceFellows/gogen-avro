// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     nested.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

type NumberRecord struct {
	IntField int32 `json:"IntField"`

	LongField int64 `json:"LongField"`

	FloatField float32 `json:"FloatField"`

	DoubleField float64 `json:"DoubleField"`
}

const NumberRecordAvroCRC64Fingerprint = "\xf4Zu\xd5Nt'~"

func NewNumberRecord() *NumberRecord {
	return &NumberRecord{}
}

func DeserializeNumberRecord(r io.Reader) (*NumberRecord, error) {
	t := NewNumberRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.AvroRecordSchema()), []byte(t.AvroRecordSchema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeNumberRecordFromSchema(r io.Reader, schema string) (*NumberRecord, error) {
	t := NewNumberRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.AvroRecordSchema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writeNumberRecord(r *NumberRecord, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.IntField, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.LongField, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.FloatField, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.DoubleField, w)
	if err != nil {
		return err
	}
	return err
}

func (r *NumberRecord) Serialize(w io.Writer) error {
	return writeNumberRecord(r, w)
}

func (r *NumberRecord) AvroRecordSchema() string {
	return "{\"fields\":[{\"name\":\"IntField\",\"type\":\"int\"},{\"name\":\"LongField\",\"type\":\"long\"},{\"name\":\"FloatField\",\"type\":\"float\"},{\"name\":\"DoubleField\",\"type\":\"double\"}],\"name\":\"NumberRecord\",\"type\":\"record\"}"
}

func (r *NumberRecord) SchemaName() string {
	return "NumberRecord"
}

func (_ *NumberRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *NumberRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *NumberRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *NumberRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *NumberRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *NumberRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *NumberRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ *NumberRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NumberRecord) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.Int{Target: &r.IntField}
	case 1:
		return &types.Long{Target: &r.LongField}
	case 2:
		return &types.Float{Target: &r.FloatField}
	case 3:
		return &types.Double{Target: &r.DoubleField}
	}
	panic("Unknown field index")
}

func (r *NumberRecord) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *NumberRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *NumberRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *NumberRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *NumberRecord) Finalize()                        {}

func (_ *NumberRecord) AvroCRC64Fingerprint() []byte {
	return []byte(NumberRecordAvroCRC64Fingerprint)
}