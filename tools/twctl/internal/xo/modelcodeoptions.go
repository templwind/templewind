package xo

import "go/ast"

type ModelCodeOptions struct {
	// modelName, fullPackageName string, node *ast.File, functions []FunctionInfo, functionSignatures []string, originalPackageName string, fieldNames []string, rows, rowsExpectAutoSet, rowsWithPlaceHolder, rowsWithNamedPlaceHolder, outputPath, tableName string
	ModelName                string
	FullPackageName          string
	Node                     *ast.File
	Functions                []FunctionInfo
	FunctionSignatures       []string
	OriginalPackageName      string
	FieldNames               []string
	Rows                     string
	RowsExpectAutoSet        string
	RowsWithPlaceHolder      string
	RowsWithNamedPlaceHolder string
	TableName                string
}

type ModelCodeOptionFn func(*ModelCodeOptions)

func NewModelCodeOptions(options ...ModelCodeOptionFn) *ModelCodeOptions {
	opts := &ModelCodeOptions{}
	for _, opt := range options {
		opt(opts)
	}

	return opts
}

func WithModelName(modelName string) ModelCodeOptionFn {
	return func(opts *ModelCodeOptions) {
		opts.ModelName = modelName
	}
}

func WithFullPackageName(fullPackageName string) ModelCodeOptionFn {
	return func(opts *ModelCodeOptions) {
		opts.FullPackageName = fullPackageName
	}
}

func WithNode(node *ast.File) ModelCodeOptionFn {
	return func(opts *ModelCodeOptions) {
		opts.Node = node
	}
}

func WithFunctions(functions []FunctionInfo) ModelCodeOptionFn {
	return func(opts *ModelCodeOptions) {
		opts.Functions = functions
	}
}

func WithFunctionSignatures(functionSignatures []string) ModelCodeOptionFn {
	return func(opts *ModelCodeOptions) {
		opts.FunctionSignatures = functionSignatures
	}
}

func WithOriginalPackageName(originalPackageName string) ModelCodeOptionFn {
	return func(opts *ModelCodeOptions) {
		opts.OriginalPackageName = originalPackageName
	}
}

func WithFieldNames(fieldNames []string) ModelCodeOptionFn {
	return func(opts *ModelCodeOptions) {
		opts.FieldNames = fieldNames
	}
}

func WithRows(rows string) ModelCodeOptionFn {
	return func(opts *ModelCodeOptions) {
		opts.Rows = rows
	}
}

func WithRowsExpectAutoSet(rowsExpectAutoSet string) ModelCodeOptionFn {
	return func(opts *ModelCodeOptions) {
		opts.RowsExpectAutoSet = rowsExpectAutoSet
	}
}

func WithRowsWithPlaceHolder(rowsWithPlaceHolder string) ModelCodeOptionFn {
	return func(opts *ModelCodeOptions) {
		opts.RowsWithPlaceHolder = rowsWithPlaceHolder
	}
}

func WithRowsWithNamedPlaceHolder(rowsWithNamedPlaceHolder string) ModelCodeOptionFn {
	return func(opts *ModelCodeOptions) {
		opts.RowsWithNamedPlaceHolder = rowsWithNamedPlaceHolder
	}
}

func WithTableName(tableName string) ModelCodeOptionFn {
	return func(opts *ModelCodeOptions) {
		opts.TableName = tableName
	}
}
