package templates

const UnionTemplate = `
import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
)


type {{ .UnionEnumType }} int
const (
{{ range $i, $t := .ItemTypes -}}
	{{ if ne $i $.NullIndex -}}
	{{ $.UnionEnumType }}{{ .Name }} {{ $.UnionEnumType }} = {{ $i }}
        {{ end }}
{{ end -}}
)

type {{ .Name }} struct {
{{ range $i, $t := .ItemTypes -}}
	{{ .Name }} {{ .GoType }}
{{ end -}}
	UnionType {{ $.UnionEnumType }}
}

func {{ .SerializerMethod }}(r {{ .GoType }}, w io.Writer) error {
	{{ if ne .NullIndex -1 }}
	if r == nil {
		err := vm.WriteLong({{ $.NullIndex }}, w)
		return err
	}
        {{ end }}
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	{{ range $i, $t := .ItemTypes -}}
	  {{ if ne $i $.NullIndex -}}
	  case {{ $.ItemName $t }}:
		return {{ .SerializerMethod }}(r.{{ .Name }}, w)
          {{ end -}}
	{{ end -}}
	}
	return fmt.Errorf("invalid value for {{ .GoType }}")
}

func {{ .ConstructorMethod }} {{ .GoType }} {
	return &{{ .Name }}{}
}

func (_ {{ .GoType }}) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ {{ .GoType }}) SetInt(v int32) { panic("Unsupported operation") }
func (_ {{ .GoType }}) SetFloat(v float32) { panic("Unsupported operation") }
func (_ {{ .GoType }}) SetDouble(v float64) { panic("Unsupported operation") }
func (_ {{ .GoType }}) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ {{ .GoType }}) SetString(v string) { panic("Unsupported operation") }
func (r {{ .GoType }}) SetLong(v int64) { 
	r.UnionType = ({{ .UnionEnumType }})(v)
}
func (r {{ .GoType }}) Get(i int) types.Field {
	switch (i) {
	{{ range $i, $t := .ItemTypes -}}
	case {{ $i }}:
		{{ if $.ItemConstructor $t | ne "" -}}
		r.{{ .Name }} = {{ $.ItemConstructor $t }}
		{{ end -}}
		{{ if eq .WrapperType "" -}}
		return r.{{ .Name }}
		{{ else -}}
		return &{{ .WrapperType }}{Target: (&r.{{ .Name }})}
		{{ end -}}
	{{ end -}}
	}
	panic("Unknown field index")
}
func (_ {{ .GoType }}) NullField(i int) { panic("Unsupported operation") }
func (_ {{ .GoType }}) SetDefault(i int) { panic("Unsupported operation") }
func (_ {{ .GoType }}) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ {{ .GoType }}) AppendArray() types.Field { panic("Unsupported operation") }
func (_ {{ .GoType }}) Finalize()  { }

func (r {{ .GoType }}) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	switch r.UnionType{
	{{ range $i, $t := .ItemTypes -}}
	{{ if ne $i $.NullIndex -}}
	case {{ $.ItemName $t }}:
		return json.Marshal(map[string]interface{}{"{{ .UnionKey }}": r.{{ .Name }}})
        {{ end -}}
	{{ end -}}
	}
	return nil, fmt.Errorf("invalid value for {{ .GoType }}")
}

func (r {{ .GoType }}) UnmarshalJSON(data []byte) (error) {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	{{ range $i, $t := .ItemTypes -}}
	{{ if ne $i $.NullIndex -}}
	if value,  ok := fields["{{ .UnionKey }}"]; ok {
		r.UnionType = {{ $i }}
		return json.Unmarshal([]byte(value), &r.{{ .Name }})
	}
        {{ end -}}
	{{ end -}}
	return fmt.Errorf("invalid value for {{ .GoType }}")
}
`
