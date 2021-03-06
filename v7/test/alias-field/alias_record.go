// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     defaults.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

type AliasRecord struct {
	A string `json:"a"`

	C string `json:"c"`
}

const AliasRecordAvroCRC64Fingerprint = "B~_@rz\xdb\xc7"

func NewAliasRecord() *AliasRecord {
	return &AliasRecord{}
}

func DeserializeAliasRecord(r io.Reader) (*AliasRecord, error) {
	t := NewAliasRecord()
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

func DeserializeAliasRecordFromSchema(r io.Reader, schema string) (*AliasRecord, error) {
	t := NewAliasRecord()

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

func writeAliasRecord(r *AliasRecord, w io.Writer) error {
	var err error
	err = vm.WriteString(r.A, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.C, w)
	if err != nil {
		return err
	}
	return err
}

func (r *AliasRecord) Serialize(w io.Writer) error {
	return writeAliasRecord(r, w)
}

func (r *AliasRecord) AvroRecordSchema() string {
	return "{\"fields\":[{\"name\":\"a\",\"type\":\"string\"},{\"aliases\":[\"d\"],\"name\":\"c\",\"type\":\"string\"}],\"name\":\"AliasRecord\",\"type\":\"record\"}"
}

func (r *AliasRecord) SchemaName() string {
	return "AliasRecord"
}

func (_ *AliasRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *AliasRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *AliasRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *AliasRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *AliasRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *AliasRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *AliasRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ *AliasRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *AliasRecord) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: &r.A}
	case 1:
		return &types.String{Target: &r.C}
	}
	panic("Unknown field index")
}

func (r *AliasRecord) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *AliasRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *AliasRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *AliasRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *AliasRecord) Finalize()                        {}

func (_ *AliasRecord) AvroCRC64Fingerprint() []byte {
	return []byte(AliasRecordAvroCRC64Fingerprint)
}
