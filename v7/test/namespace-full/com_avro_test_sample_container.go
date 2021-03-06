// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/container"
	"github.com/actgardner/gogen-avro/v7/vm"
)

func NewComAvroTestSampleWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewComAvroTestSample()
	return container.NewWriter(writer, codec, recordsPerBlock, str.AvroRecordSchema())
}

// container reader
type ComAvroTestSampleReader struct {
	r io.Reader
	p *vm.Program
}

func NewComAvroTestSampleReader(r io.Reader) (*ComAvroTestSampleReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewComAvroTestSample()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.AvroRecordSchema()))
	if err != nil {
		return nil, err
	}

	return &ComAvroTestSampleReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r ComAvroTestSampleReader) Read() (*ComAvroTestSample, error) {
	t := NewComAvroTestSample()
	err := vm.Eval(r.r, r.p, t)
	return t, err
}
