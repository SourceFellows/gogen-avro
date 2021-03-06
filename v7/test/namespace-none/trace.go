// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

// Trace
type Trace struct {
	// Trace Identifier
	TraceId *UnionNullUUID `json:"traceId"`
}

const TraceAvroCRC64Fingerprint = "\x83<\x8e\xd5T\xfc\x8d\x94"

func NewTrace() *Trace {
	return &Trace{}
}

func DeserializeTrace(r io.Reader) (*Trace, error) {
	t := NewTrace()
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

func DeserializeTraceFromSchema(r io.Reader, schema string) (*Trace, error) {
	t := NewTrace()

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

func writeTrace(r *Trace, w io.Writer) error {
	var err error
	err = writeUnionNullUUID(r.TraceId, w)
	if err != nil {
		return err
	}
	return err
}

func (r *Trace) Serialize(w io.Writer) error {
	return writeTrace(r, w)
}

func (r *Trace) AvroRecordSchema() string {
	return "{\"doc\":\"Trace\",\"fields\":[{\"default\":null,\"doc\":\"Trace Identifier\",\"name\":\"traceId\",\"type\":[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"headerworks.datatype\",\"type\":\"record\"}]}],\"name\":\"bodyworks.Trace\",\"type\":\"record\"}"
}

func (r *Trace) SchemaName() string {
	return "bodyworks.Trace"
}

func (_ *Trace) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *Trace) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *Trace) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *Trace) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *Trace) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *Trace) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *Trace) SetString(v string)   { panic("Unsupported operation") }
func (_ *Trace) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Trace) Get(i int) types.Field {
	switch i {
	case 0:
		r.TraceId = NewUnionNullUUID()

		return r.TraceId
	}
	panic("Unknown field index")
}

func (r *Trace) SetDefault(i int) {
	switch i {
	case 0:
		r.TraceId = nil
		return
	}
	panic("Unknown field index")
}

func (r *Trace) NullField(i int) {
	switch i {
	case 0:
		r.TraceId = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ *Trace) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Trace) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *Trace) Finalize()                        {}

func (_ *Trace) AvroCRC64Fingerprint() []byte {
	return []byte(TraceAvroCRC64Fingerprint)
}
