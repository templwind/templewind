{{- $modelName := .ModelName -}}
{{- $includePkg := .IncludePackageName -}}
{{- $originalPkg := .OriginalPackageName -}}
{{- $functionSignatures := .FunctionSignatures -}}
package {{.IncludePackageName}}

// Code generated. DO NOT EDIT.

import (
	"context"
{{- if .UsesSqlNull }}
    "database/sql"
{{- end }}
{{- if or .UsesResourcesTypes $includePkg -}}
{{ "\n" }}
{{- end -}}
{{- if .UsesResourcesTypes }}
    "{{.BaseImportPath}}/types"
{{- end }}
{{- if $includePkg }}
    "{{.FullPackageName}}"
{{- end }}
	"github.com/jmoiron/sqlx"  
)

var (
	{{$modelName}}TableName                = "{{RawTableName .TableName}}"
	{{$modelName}}FieldNames               = {{FormatFieldNames .FieldNames}}
	{{$modelName}}Rows                     = "{{.Rows}}"
	{{$modelName}}RowsExpectAutoSet        = "{{.RowsExpectAutoSet}}"
	{{$modelName}}RowsWithPlaceHolder      = "{{.RowsWithPlaceHolder}}"
	{{$modelName}}RowsWithNamedPlaceHolder = "{{.RowsWithNamedPlaceHolder}}"
)

// Definition of the model interface and its implementation.
// It includes transaction management and methods for operating on the model.
type (
	{{FirstToLower $modelName}}Model interface {
		{{- range .Functions}}
		{{.Decl.Name.Name}}({{ GetFunctionParams .Decl $originalPkg .IsReceiver $modelName}}) {{ GetFunctionReturnType .Decl $originalPkg }}
		{{- end }}
	}

	default{{.ModelName}}Model struct {
		{{- if $includePkg }}
		*{{$originalPkg}}.{{$modelName}}
		{{- else }}
		*{{$modelName}}
		{{- end }}
		transaction transactions
		table string
	}
)

func new{{$modelName}}Model(db *sqlx.DB) *default{{$modelName}}Model {
	return &default{{$modelName}}Model{
		transaction: transactions{db: db},
		table: {{FormatTableName .TableName}},
	}
}

/////////////////////////
// For Insert, Update, Delete, Save, and Upsert use the adapter for transaction handling
// Other functions remain as-is
/////////////////////////

{{- range .Functions}}
    {{- if eq .Decl.Name.Name "Insert" "Update" "Delete" "Save" "Upsert"}} 

func (m *default{{$modelName}}Model) {{.Decl.Name.Name}}({{ GetFunctionParams .Decl $originalPkg .IsReceiver $modelName}}) {{ GetFunctionReturnType .Decl $originalPkg }} {
	return m.transaction.adapter(ctx, func(ctx context.Context, db xo.DB) error {
		return {{FirstToLower $modelName}}.{{.Decl.Name.Name}}({{ GetFunctionCleanParams .Decl true }})
	})
}
    {{- else}}

func (m *default{{$modelName}}Model) {{.Decl.Name.Name}}({{ GetFunctionParams .Decl $originalPkg .IsReceiver $modelName}}) {{ GetFunctionReturnType .Decl $originalPkg }} {
	{{- if .IsReceiver }}
		return {{FirstToLower $modelName}}.{{RemoveDuplicateFunctionMarker .Decl}}({{ GetFunctionCleanParams .Decl false }})
	{{- else}}
		{{- if $includePkg }}
			return {{$originalPkg}}.{{.Decl.Name.Name}}({{ GetFunctionCleanParams .Decl false }})
		{{- else }}
			return {{RemoveDuplicateFunctionMarker .Decl}}({{ GetFunctionCleanParams .Decl false }})
		{{- end }}
	{{- end }}
}
    {{- end}}
{{- end}}