package parsexo

const XoExtensionTemplate = `{{- $modelName := .ModelName -}}
{{- $originalPkg := .OriginalPackageName -}}
{{- $functionSignatures := .FunctionSignatures -}}
package {{$originalPkg}}

import (
	"context"
	"fmt"
	"strings"

	"{{.BaseImportPath}}/buildsql"
	"{{.BaseImportPath}}/types"
)

var (
	{{$modelName}}TableName                = "{{RawTableName .TableName}}"
	{{$modelName}}FieldNames               = {{FormatFieldNames .FieldNames}}
	{{$modelName}}Rows                     = "{{.Rows}}"
	{{$modelName}}RowsExpectAutoSet        = "{{.RowsExpectAutoSet}}"
	{{$modelName}}RowsWithPlaceHolder      = "{{.RowsWithPlaceHolder}}"
	{{$modelName}}RowsWithNamedPlaceHolder = "{{.RowsWithNamedPlaceHolder}}"
)

func FindAll{{.ModelName}}s(ctx context.Context, db SqlxDB, page int, pageSize int) ([]*{{.ModelName}}, error) {
    var query string
    if pageSize == 0{
        query = fmt.Sprintf({{WrapInBackticks "SELECT %s FROM %s"}}, {{.ModelName}}Rows, {{$modelName}}TableName)
    } else {
        offset := (page - 1) * pageSize
        query = fmt.Sprintf({{WrapInBackticks "SELECT %s FROM %s LIMIT %d OFFSET %d"}}, {{.ModelName}}Rows, {{$modelName}}TableName, pageSize, offset)
    }

    var results []*{{.ModelName}}
    err := db.SelectContext(ctx, &results, query)
    if err != nil {
        return nil, err
    }
    return results, nil
}

// response type
type Search{{.ModelName}}Response struct {
	{{.ModelName}}s []{{.ModelName}}
	PagingStats    types.PagingStats
}

func Search{{.ModelName}}s(ctx context.Context, db SqlxDB, currentPage, pageSize int64, filter string) (res *Search{{.ModelName}}Response, err error) {
	var builder = buildsql.NewQueryBuilder()
	where, orderBy, namedParamMap, err := builder.Build(filter, map[string]interface{}{
		"t1": {{.ModelName}}{},
	})
	if err != nil {
		return nil, err
	}

	if where != "" {
		where = fmt.Sprintf("WHERE 1 = 1 %s", where)
	}

	// set a default order by
	if orderBy == "" {
		orderBy = "ORDER BY t1.id DESC"
	}
	limit := fmt.Sprintf("LIMIT %d, %d", currentPage*pageSize, pageSize)

	// field names
	var fieldNames []string
	for _, fieldName := range {{.ModelName}}FieldNames {
		fieldNames = append(fieldNames, fmt.Sprintf("t1.%s as \"%s.%s\"", fieldName, {{$modelName}}TableName, fieldName))
	}

	// fmt.Println("fieldNames:", fieldNames)
	// fmt.Println("tableNameNoTicks:", {{$modelName}}TableName)

	sql := fmt.Sprintf({{InsertBacktick}}
		SELECT
			-- {{RawTableName .TableName}}
			%s,
			-- stats
			COUNT(*) OVER() AS "pagingstats.total_records"
		FROM {{RawTableName .TableName}} t1
		%s
		%s
		%s
	{{InsertBacktick}}, strings.Join(fieldNames, ", "), where, orderBy, limit)

	nstmt, err := db.PrepareNamedContext(ctx, sql)

	if err != nil {
		return nil, fmt.Errorf("error::Search::Prepared::%s", err.Error())
	}

	var result []struct {
		{{.ModelName}} {{.ModelName}}    {{InsertBacktick}}db:"{{RawTableName .TableName}}"{{InsertBacktick}}
		PagingStats   types.PagingStats  {{InsertBacktick}}db:"pagingstats"{{InsertBacktick}}
	}

	namedParamMap["offset"] = currentPage * pageSize
	namedParamMap["limit"] = pageSize

	err = nstmt.Select(&result, namedParamMap)
	if err != nil {
		return nil, fmt.Errorf("error::Search::Select::%s", err.Error())
	}

	records := []{{.ModelName}}{}

	var stats *types.PagingStats = &types.PagingStats{}
	for i, r := range result {
		if i == 0 {
			stats = r.PagingStats.Calc(pageSize)
		}
		records = append(records, r.{{.ModelName}})
	}

	out := &Search{{.ModelName}}Response{
		{{.ModelName}}s: records,
		PagingStats:    *stats,
	}

	return out, err
}
`
