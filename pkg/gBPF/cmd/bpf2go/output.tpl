// Code generated by bpf2go; DO NOT EDIT.
{{ with .Constraints }}//go:build {{ . }}{{ end }}

package {{ .Package }}

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"{{ .Module }}"
)

{{- if .Types }}
{{- range $type := .Types }}
{{ $.TypeDeclaration (index $.TypeNames $type) $type }}

{{ end }}
{{- end }}

// {{ .Name.Load }} returns the embedded CollectionSpec for {{ .Name }}.
func {{ .Name.Load }}() (*gbpf.CollectionSpec, error) {
	reader := bytes.NewReader({{ .Name.Bytes }})
	spec, err := gbpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load {{ .Name }}: %w", err)
	}

	return spec, err
}

// {{ .Name.LoadObjects }} loads {{ .Name }} and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*{{ .Name.Objects }}
//	*{{ .Name.Programs }}
//	*{{ .Name.Maps }}
//
// See gbpf.CollectionSpec.LoadAndAssign documentation for details.
func {{ .Name.LoadObjects }}(obj interface{}, opts *gbpf.CollectionOptions) (error) {
	spec, err := {{ .Name.Load }}()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// {{ .Name.Specs }} contains maps and programs before they are loaded into the kernel.
//
// It can be passed gbpf.CollectionSpec.Assign.
type {{ .Name.Specs }} struct {
	{{ .Name.ProgramSpecs }}
	{{ .Name.MapSpecs }}
}

// {{ .Name.Specs }} contains programs before they are loaded into the kernel.
//
// It can be passed gbpf.CollectionSpec.Assign.
type {{ .Name.ProgramSpecs }} struct {
{{- range $name, $id := .Programs }}
	{{ $id }} *gbpf.ProgramSpec `gbpf:"{{ $name }}"`
{{- end }}
}

// {{ .Name.MapSpecs }} contains maps before they are loaded into the kernel.
//
// It can be passed gbpf.CollectionSpec.Assign.
type {{ .Name.MapSpecs }} struct {
{{- range $name, $id := .Maps }}
	{{ $id }} *gbpf.MapSpec `gbpf:"{{ $name }}"`
{{- end }}
}

// {{ .Name.Objects }} contains all objects after they have been loaded into the kernel.
//
// It can be passed to {{ .Name.LoadObjects }} or gbpf.CollectionSpec.LoadAndAssign.
type {{ .Name.Objects }} struct {
	{{ .Name.Programs }}
	{{ .Name.Maps }}
}

func (o *{{ .Name.Objects }}) Close() error {
	return {{ .Name.CloseHelper }}(
		&o.{{ .Name.Programs }},
		&o.{{ .Name.Maps }},
	)
}

// {{ .Name.Maps }} contains all maps after they have been loaded into the kernel.
//
// It can be passed to {{ .Name.LoadObjects }} or gbpf.CollectionSpec.LoadAndAssign.
type {{ .Name.Maps }} struct {
{{- range $name, $id := .Maps }}
	{{ $id }} *gbpf.Map `gbpf:"{{ $name }}"`
{{- end }}
}

func (m *{{ .Name.Maps }}) Close() error {
	return {{ .Name.CloseHelper }}(
{{- range $id := .Maps }}
		m.{{ $id }},
{{- end }}
	)
}

// {{ .Name.Programs }} contains all programs after they have been loaded into the kernel.
//
// It can be passed to {{ .Name.LoadObjects }} or gbpf.CollectionSpec.LoadAndAssign.
type {{ .Name.Programs }} struct {
{{- range $name, $id := .Programs }}
	{{ $id }} *gbpf.Program `gbpf:"{{ $name }}"`
{{- end }}
}

func (p *{{ .Name.Programs }}) Close() error {
	return {{ .Name.CloseHelper }}(
{{- range $id := .Programs }}
		p.{{ $id }},
{{- end }}
	)
}

func {{ .Name.CloseHelper }}(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//go:embed {{ .File }}
var {{ .Name.Bytes }} []byte