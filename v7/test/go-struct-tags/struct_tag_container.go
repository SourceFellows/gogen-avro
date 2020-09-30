// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     tagstest.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/container"
	"github.com/actgardner/gogen-avro/v7/vm"
)

func NewStructTagWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewStructTag()
	return container.NewWriter(writer, codec, recordsPerBlock, str.AvroRecordSchema())
}

// container reader
type StructTagReader struct {
	r io.Reader
	p *vm.Program
}

func NewStructTagReader(r io.Reader) (*StructTagReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewStructTag()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.AvroRecordSchema()))
	if err != nil {
		return nil, err
	}

	return &StructTagReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r StructTagReader) Read() (*StructTag, error) {
	t := NewStructTag()
	err := vm.Eval(r.r, r.p, t)
	return t, err
}
