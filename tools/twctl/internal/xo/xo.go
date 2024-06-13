package xo

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"
	"unicode"

	"github.com/templwind/templwind/tools/twctl/internal/utils"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ProcessFiles(inputPath, outputPath string, ignoredTypes map[string]bool, baseImportPath string) error {
	// Initialize the process with the provided options
	process := New(
		WithInputPath(inputPath),
		WithOutputPath(outputPath),
		WithIgnoredTypes(ignoredTypes),
		WithBaseImportPath(baseImportPath),
	)

	// Resolve the absolute paths of input and output directories
	absInputPath, err := filepath.Abs(inputPath)
	if err != nil {
		return fmt.Errorf("error resolving absolute path of input directory: %v", err)
	}

	absOutputPath, err := filepath.Abs(outputPath)
	if err != nil {
		return fmt.Errorf("error resolving absolute path of output directory: %v", err)
	}

	// Find the go.mod file for the source file
	goModPath, _ := utils.FindGoModPath(filepath.Dir(absInputPath))
	// Extract module name from go.mod
	moduleName, _ := utils.GetModuleName(goModPath)

	err = process.generateStaticFiles(moduleName, absOutputPath)
	if err != nil {
		return fmt.Errorf("error generating static files: %v", err)
	}

	fmt.Println("Parsing .xo.go files in:", absInputPath)
	err = filepath.Walk(absInputPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".xo.go") {
			if strings.Contains(path, "/sf_") {
				fmt.Println("Not attempting to parse sf file:", path)
				return nil
			}
			err := process.ProcessXoFile(path)
			if err != nil {
				fmt.Println("Error parsing file:", path)
				return err
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error parsing files: %v", err)
	}

	return nil
}

func (p *Process) ProcessXoFile(filename string) error {

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	// Skip processing if the file contains an enum declaration
	if p.containsEnumDeclaration(node) {
		fmt.Printf("Skipping file %s as it contains enum declaration\n", filename)
		return nil

	}

	// Skip processing if the file contains an ignored type.
	if p.containsIgnoredType(node, p.IgnoredTypes) {
		fmt.Printf("Skipping file %s as it contains ignored types\n", filename)
		return nil
	}

	// Initialize a map to track function names
	functionNameMap := make(map[string]struct{})

	var functions []FunctionInfo
	var functionSignatures []string // Store function signatures as strings
	var tableName string

	// Loop through the declarations in the file
	for _, f := range node.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}

		// Check for duplicate function names
		if _, exists := functionNameMap[fn.Name.Name]; exists {
			// rename the function name
			fn.Name.Name = fn.Name.Name + p.DuplicateFunctionMarker
		}

		// Add the function name to the map
		functionNameMap[fn.Name.Name] = struct{}{}

		isReceiver := fn.Recv != nil && len(fn.Recv.List) > 0
		functions = append(functions, FunctionInfo{
			Decl:       fn,
			IsReceiver: isReceiver,
		})
	}

	// Directly inspect all comments
	for _, commentGroup := range node.Comments {
		for _, comment := range commentGroup.List {
			// Look for a comment that starts with "represents a row from"
			if strings.Contains(comment.Text, "represents a row from") {
				// Extract the table name from the comment
				parts := strings.Split(comment.Text, "'")
				if len(parts) >= 3 {
					tableName = parts[1] // The table name is between the single quotes
					break
				}
			}
		}

		// Break the outer loop if table name is found
		if tableName != "" {
			break
		}
	}

	// Use the filename or struct name to generate modelName
	// Get the struct name from the AST
	modelName, err := p.getStructNameFromAstFile(node)
	if err != nil {
		return err
	}
	fieldNames, rows, rowsExpectAutoSet, rowsWithPlaceHolder, rowsWithNamedPlaceHolder := p.extractFieldData(node)

	inputPackageName := p.getLastPathComponent(p.InputPath)

	// Find the go.mod file for the source file
	goModPath, err := utils.FindGoModPath(filepath.Dir(filename))
	if err != nil {
		return fmt.Errorf("error finding go.mod: %v", err)
	}

	// Extract module name from go.mod
	moduleName, err := utils.GetModuleName(goModPath)
	if err != nil {
		return fmt.Errorf("error getting module name from go.mod: %v", err)
	}

	// Get the relative path from the module root to the source file's directory
	moduleDir := filepath.Dir(goModPath)
	relativePath, err := filepath.Rel(moduleDir, filepath.Dir(filename))
	if err != nil {
		return fmt.Errorf("error finding relative path: %v", err)
	}

	// Replace OS-specific path separators with slashes for consistent package naming
	relativePath = strings.ReplaceAll(relativePath, string(os.PathSeparator), "/")

	// Concatenate the module name and the relative path to form the full package name
	fullPackageName := moduleName
	if relativePath != "." {
		fullPackageName += "/" + relativePath
	}

	// fmt.Println("Full Package Name: ", fullPackageName)
	return p.generateModelCode(
		NewModelCodeOptions(
			WithModelName(modelName),
			WithFullPackageName(fullPackageName),
			WithNode(node),
			WithFunctions(functions),
			WithFunctionSignatures(functionSignatures),
			WithOriginalPackageName(inputPackageName),
			WithFieldNames(fieldNames),
			WithRows(rows),
			WithRowsExpectAutoSet(rowsExpectAutoSet),
			WithRowsWithPlaceHolder(rowsWithPlaceHolder),
			WithTableName(tableName),
			WithRowsWithNamedPlaceHolder(rowsWithNamedPlaceHolder),
		),
	)

	// return generateModelCode(
	// 	modelName,
	// 	fullPackageName,
	// 	node,
	// 	functions,
	// 	functionSignatures,
	// 	inputPackageName,
	// 	fieldNames,
	// 	rows,
	// 	rowsExpectAutoSet,
	// 	rowsWithPlaceHolder,
	// 	rowsWithNamedPlaceHolder,
	// 	tableName,
	// )
}

// Check if the file contains an enum declaration.
func (p *Process) containsEnumDeclaration(node *ast.File) bool {
	for _, f := range node.Decls {
		genDecl, ok := f.(*ast.GenDecl)
		if !ok {
			continue
		}
		if genDecl.Tok == token.CONST {
			return true
		}
	}
	return false
}

// Check if a file contains any ignored types.
func (p *Process) containsIgnoredType(node *ast.File, ignoredTypes map[string]bool) bool {
	for _, f := range node.Decls {
		genDecl, ok := f.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			if _, ok := ignoredTypes[typeSpec.Name.Name]; ok {
				return true
			}
		}
	}
	return false
}

func (p *Process) getLastPathComponent(path string) string {
	parts := strings.Split(filepath.Clean(path), string(os.PathSeparator))
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return ""
}

func (p *Process) getStructNameFromAstFile(node *ast.File) (string, error) {
	// Loop through the declarations to find the struct
	for _, f := range node.Decls {
		if genDecl, ok := f.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if _, ok := typeSpec.Type.(*ast.StructType); ok {
						return typeSpec.Name.Name, nil // Return struct name
					}
				}
			}
		}
	}

	return "", fmt.Errorf("no struct found in file")
}

func (p *Process) exprToString(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.SelectorExpr:
		return p.exprToString(e.X) + "." + e.Sel.Name
	case *ast.StarExpr:
		return "*" + p.exprToString(e.X)
	case *ast.ArrayType:
		return "[]" + p.exprToString(e.Elt)
	default:
		return ""
	}
}

func (p *Process) extractFieldData(file *ast.File) ([]string, string, string, string, string) {
	var fieldNames []string

	// Loop through the declarations in the file
	for _, decl := range file.Decls {
		// Check if the declaration is a GenDecl (general declaration)
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		// Loop through the specifications in the declaration
		for _, spec := range genDecl.Specs {
			// Check if the spec is a TypeSpec (type specification)
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			// Check if the type spec is a struct type
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			// Loop through the fields in the struct type
			for _, field := range structType.Fields.List {
				// Extract the tag, assuming `db` tag is used
				if field.Tag != nil {
					tag := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
					dbName := tag.Get("db")
					if dbName != "" {
						fieldNames = append(fieldNames, dbName)
					}
				}
			}
		}
	}

	fieldNamesStr := strings.Join(fieldNames, ",")
	rowsExpectAutoSet := strings.Join(stringx.Remove(fieldNames, "id"), ",")
	rowsWithPlaceHolder := builder.PostgreSqlJoin(stringx.Remove(fieldNames, "id"))
	// Logic for named placeholders
	var namedPlaceholders []string
	for _, fieldName := range stringx.Remove(fieldNames, "id") {
		namedPlaceholders = append(namedPlaceholders, fieldName+" = :"+fieldName)
	}
	rowsWithNamedPlaceHolder := strings.Join(namedPlaceholders, ", ")

	return fieldNames, fieldNamesStr, rowsExpectAutoSet, rowsWithPlaceHolder, rowsWithNamedPlaceHolder
}

type FunctionInfo struct {
	Decl       *ast.FuncDecl
	IsReceiver bool
}

// Create a function to generate the function return types map
func (p *Process) makeFunctionReturnTypesMap(functions []FunctionInfo) map[string]string {
	functionReturnTypes := make(map[string]string)
	for _, fn := range functions {
		functionReturnTypes[fn.Decl.Name.Name] = p.getFunctionReturnType(fn.Decl, "")
	}
	return functionReturnTypes
}

type defaultInterfaceTemplateData struct {
	FullPackageName          string
	OriginalPackageName      string
	IncludePackageName       string
	UsesSqlNull              bool
	UsesResourcesTypes       bool
	ModelName                string
	FunctionSignatures       []string
	Functions                []FunctionInfo
	FieldNames               []string
	Rows                     string
	RowsExpectAutoSet        string
	RowsWithPlaceHolder      string
	RowsWithNamedPlaceHolder string
	FunctionReturnTypes      map[string]string
	TypeMethods              map[string]bool
	TableName                string
	BaseImportPath           string
}

func (p *Process) generateStaticFiles(packageName, outputPath string) error {
	// Extract the package name from the output path
	outputPathBase := filepath.Base(outputPath)
	includePackageName := cases.Lower(language.English).String(outputPathBase) // Ensure the package name is in lowercase

	// Ensure the output directory exists
	err := os.MkdirAll(outputPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Create the data for the template
	data := defaultInterfaceTemplateData{
		// FullPackageName:          fullPackageName,
		// OriginalPackageName:      originalPackageName,
		IncludePackageName: includePackageName,
		BaseImportPath:     p.BaseImportPath,
	}

	funcMap := template.FuncMap{
		"GetFunctionParams":             p.getFunctionParams,
		"GetFunctionCleanParams":        p.getFunctionCleanParams,
		"RemoveDuplicateFunctionMarker": p.removeDuplicateFunctionMarker,
		"GetFunctionReturnType":         p.getFunctionReturnType,
		"FirstToLower":                  utils.FirstToLower,
		"FormatTableName":               p.formatTableName,
		"FormatFieldNames":              p.formatFieldNames,
		"WrapInBackticks":               p.wrapInBackticks,
		"InsertBacktick":                p.insertBacktick,
		"RawTableName":                  p.rawTableName,
	}

	filePath := filepath.Join(p.InputPath, "transactions.ext.go")
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create transaction file: %v", err)
	}
	defer file.Close()

	// Parse the template
	tmpl, err := template.New("transaction").Funcs(funcMap).ParseFiles("templates/xo/transaction.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse template for extension: %v", err)
	}
	// Execute the template with the data struct
	err = tmpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("failed to execute template for transactions: %v", err)
	}
	return nil
}

func (p *Process) generateModelCode(opts *ModelCodeOptions) error {

	// modelName, fullPackageName string, node *ast.File, functions []FunctionInfo, functionSignatures []string, originalPackageName string, fieldNames []string, rows, rowsExpectAutoSet, rowsWithPlaceHolder, rowsWithNamedPlaceHolder, p.OutputPath, tableName string

	// Extract the package name from the output path
	outputPathBase := filepath.Base(p.OutputPath)
	includePackageName := cases.Lower(language.English).String(outputPathBase) // Ensure the package name is in lowercase

	// Ensure the output directory exists
	err := os.MkdirAll(p.OutputPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Check for SQL null types in the AST node
	usesSqlNull := p.containsSqlNullTypes(opts.Node)
	usesResourcesTypes := p.containsResourcesTypes(opts.Node)

	// Create the data for the template
	data := defaultInterfaceTemplateData{
		FullPackageName:          opts.FullPackageName,
		OriginalPackageName:      opts.OriginalPackageName,
		IncludePackageName:       includePackageName,
		UsesSqlNull:              usesSqlNull,
		UsesResourcesTypes:       usesResourcesTypes,
		ModelName:                opts.ModelName,
		FunctionSignatures:       opts.FunctionSignatures,
		Functions:                opts.Functions,
		FieldNames:               opts.FieldNames,
		Rows:                     opts.Rows,
		RowsExpectAutoSet:        opts.RowsExpectAutoSet,
		RowsWithPlaceHolder:      opts.RowsWithPlaceHolder,
		RowsWithNamedPlaceHolder: opts.RowsWithNamedPlaceHolder,
		FunctionReturnTypes:      p.makeFunctionReturnTypesMap(opts.Functions),
		TableName:                opts.TableName,
		BaseImportPath:           p.BaseImportPath,
	}

	funcMap := template.FuncMap{
		"GetFunctionParams":             p.getFunctionParams,
		"GetFunctionCleanParams":        p.getFunctionCleanParams,
		"RemoveDuplicateFunctionMarker": p.removeDuplicateFunctionMarker,
		"GetFunctionReturnType":         p.getFunctionReturnType,
		"FirstToLower":                  utils.FirstToLower,
		"FormatTableName":               p.formatTableName,
		"FormatFieldNames":              p.formatFieldNames,
		"WrapInBackticks":               p.wrapInBackticks,
		"InsertBacktick":                p.insertBacktick,
		"RawTableName":                  p.rawTableName,
	}

	fmt.Println("Generating model code for:", opts.ModelName)
	{

		if p.XoOnly {

			extensionFilePath := filepath.Join(p.InputPath, fmt.Sprintf("%s.ext.go", strings.ToLower(opts.ModelName)))
			extensionFile, err := os.Create(extensionFilePath)
			if err != nil {
				return fmt.Errorf("failed to create extension file: %v", err)
			}
			defer extensionFile.Close()

			// Parse the template
			extTmpl, err := template.New("model").Funcs(funcMap).ParseFiles("templates/xo/xo_extension.tpl")
			if err != nil {
				return fmt.Errorf("failed to parse template for extension: %v", err)
			}
			// Execute the template with the data struct
			err = extTmpl.Execute(extensionFile, data)
			if err != nil {
				return fmt.Errorf("failed to execute template for extension: %v", err)
			}
			return nil
		}

		// Create the output file
		outputFilePath := filepath.Join(p.OutputPath, fmt.Sprintf("%s_model.gen.go", utils.ToSnake(opts.ModelName)))
		outputFile, err := os.Create(outputFilePath)
		if err != nil {
			return fmt.Errorf("failed to create output file: %v", err)
		}
		defer outputFile.Close()

		// Parse the template
		tmpl, err := template.New("model").Funcs(funcMap).ParseFiles("templates/xo/default_interface.tpl")
		if err != nil {
			return fmt.Errorf("failed to parse template: %v", err)
		}

		// Execute the template with the data struct
		err = tmpl.Execute(outputFile, data)
		if err != nil {
			return fmt.Errorf("failed to execute template: %v", err)
		}

		fmt.Printf("Generated model code at %s\n", outputFilePath)
	}

	// Generate custom interface template here
	{
		// Create the output file
		outputFilePath := filepath.Join(p.OutputPath, fmt.Sprintf("%s_model.go", utils.ToSnake(opts.ModelName)))
		// if the outputFilePath already exists then skip
		if _, err := os.Stat(outputFilePath); !os.IsNotExist(err) {
			fmt.Printf("Skipping generation of custom model code at %s\n", outputFilePath)
			return nil
		}

		outputFile, err := os.Create(outputFilePath)
		if err != nil {
			return fmt.Errorf("failed to create output file: %v", err)
		}
		defer outputFile.Close()

		// Parse the template
		tmpl, err := template.New("model").Funcs(funcMap).ParseFiles("templates/xo/custom_interface.tpl")
		if err != nil {
			return fmt.Errorf("failed to parse template: %v", err)
		}
		// Execute the template with the data struct
		err = tmpl.Execute(outputFile, data)
		if err != nil {
			return fmt.Errorf("failed to execute template: %v", err)
		}

		fmt.Printf("Generated model code at %s\n", outputFilePath)
	}
	return nil
}

func (p *Process) getFunctionParams(fn *ast.FuncDecl, originalPkg string, isReceiver bool, modelName string) string {
	params := make([]string, 0)
	for _, item := range fn.Type.Params.List {
		paramType := p.exprToStringWithPkg(item.Type, originalPkg)

		// Check if the parameter type is a pointer
		if _, ok := item.Type.(*ast.StarExpr); ok {
			paramType = "*" + paramType
		}

		for _, name := range item.Names {
			if !strings.Contains(paramType, "DB") {
				params = append(params, name.Name+" "+paramType)
			}
		}
	}

	if isReceiver {
		// Include the model object as the first parameter for receiver functions
		receiverType := modelName
		if originalPkg != "" {
			receiverType = originalPkg + "." + modelName
		}

		// make the modelName have a lowercase first letter
		modelName = utils.FirstToLower(modelName)
		params = append(params, fmt.Sprintf("%s *%s", modelName, receiverType))
	}

	return strings.Join(params, ", ")
}

func (p *Process) getFunctionCleanParams(fn *ast.FuncDecl, useAdapter bool) string {
	var params []string
	for _, item := range fn.Type.Params.List {
		for _, name := range item.Names {
			paramType := p.exprToStringWithPkg(item.Type, "")
			if strings.Contains(paramType, "DB") {
				if useAdapter {
					params = append(params, "db")
				} else {
					params = append(params, "m.transaction.db")
				}
			} else {
				params = append(params, name.Name)
			}
		}
	}
	return strings.Join(params, ", ")
}

func (p *Process) removeDuplicateFunctionMarker(fn *ast.FuncDecl) string {
	return strings.ReplaceAll(fn.Name.Name, p.DuplicateFunctionMarker, "")
}

func (p *Process) getFunctionReturnType(fn *ast.FuncDecl, originalPkg string) string {
	if fn.Type.Results == nil {
		return ""
	}

	var packageName string
	if originalPkg != "" {
		packageName = originalPkg + "."
	}

	results := make([]string, 0)
	for _, r := range fn.Type.Results.List {
		resultType := p.exprToStringWithPkg(r.Type, packageName)

		// replace .. with .
		resultType = strings.ReplaceAll(resultType, "..", ".")

		results = append(results, resultType)
	}

	// Check the number of return types
	numResults := len(results)
	if numResults == 1 {
		// Single return type
		return results[0]
	} else if numResults > 1 {
		// Multiple return types
		return "(" + strings.Join(results, ", ") + ")"
	} else {
		// No return types
		return ""
	}
}

func (p *Process) exprToStringWithPkg(expr ast.Expr, packageName string) string {
	switch e := expr.(type) {
	case *ast.Ident:
		typeName := e.Name
		// If it's a native Go type or an ignored type, return as is.
		if p.isNativeType(typeName) || p.IgnoredTypes[typeName] {
			return typeName
		}
		// Prepend the package name if it's a non-native type and the package name is provided.
		if packageName != "" {
			return packageName + "." + typeName
		}
		return typeName
	case *ast.SelectorExpr:
		// Handle selector expressions.
		fullTypeName := p.exprToString(e.X) + "." + e.Sel.Name
		if p.IgnoredTypes[fullTypeName] {
			return fullTypeName
		}
		// Check if it already represents a package name.
		if xIdent, ok := e.X.(*ast.Ident); ok {
			if p.isPackageName(xIdent.Name) {
				return fullTypeName
			}
		}
		// If not a package identifier, prepend the package name if provided.
		if packageName != "" {
			return packageName + "." + fullTypeName
		}
		return fullTypeName
	case *ast.StarExpr:
		// For pointer types, add '*' and process the element type.
		return "*" + p.exprToStringWithPkg(e.X, packageName)
	case *ast.ArrayType:
		// For array types, add '[]' and process the element type.
		return "[]" + p.exprToStringWithPkg(e.Elt, packageName)
	default:
		return ""
	}
}

// isPackageName checks if the provided name is a known package name
func (p *Process) isPackageName(name string) bool {
	// Check if the name contains a dot, indicating it might be a package or type name from a package
	if strings.Contains(name, ".") {
		return true
	}

	// Check if the name is not a native type and if it begins with a lowercase letter
	if !p.isNativeType(name) && unicode.IsLower(rune(name[0])) {
		return true
	}

	return false
}

// isNativeType checks if the provided type name is a native Go type
func (p *Process) isNativeType(typeName string) bool {
	nativeTypes := map[string]bool{
		"string": true, "int": true, "int8": true, "int16": true,
		"int32": true, "int64": true, "uint": true, "uint8": true,
		"uint16": true, "uint32": true, "uint64": true, "uintptr": true,
		"byte": true, "rune": true, "float32": true, "float64": true,
		"complex64": true, "complex128": true, "bool": true, "error": true,
		"interface{}": true,
		// Standard library types
		"context.Context": true,
		// Add other standard library types as needed
	}
	_, exists := nativeTypes[typeName]
	return exists
}

func (p *Process) containsSqlNullTypes(node *ast.File) bool {
	// Named recursive function
	var checkType func(ast.Expr) bool
	checkType = func(expr ast.Expr) bool {
		switch e := expr.(type) {
		case *ast.SelectorExpr:
			if pkgIdent, ok := e.X.(*ast.Ident); ok {
				// Check for "sql" package and "Null" prefix in type
				if pkgIdent.Name == "sql" && strings.HasPrefix(e.Sel.Name, "Null") {
					return true
				}
			}
		case *ast.StarExpr:
			// Recursively check the element type for *ast.StarExpr
			return checkType(e.X)
		case *ast.ArrayType:
			// Recursively check the element type for *ast.ArrayType
			return checkType(e.Elt)
		}
		return false
	}

	// Define checkSqlNullType using the named recursive function
	checkSqlNullType := func(expr ast.Expr) bool {
		return checkType(expr)
	}

	// Check each function declaration for sql.Null* types
	for _, decl := range node.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		// Check parameters for sql.Null* types
		if fn.Type.Params != nil {
			for _, p := range fn.Type.Params.List {
				if checkSqlNullType(p.Type) {
					return true
				}
			}
		}

		// Check return types for sql.Null* types
		if fn.Type.Results != nil {
			for _, r := range fn.Type.Results.List {
				if checkSqlNullType(r.Type) {
					return true
				}
			}
		}
	}
	return false
}

func (p *Process) containsResourcesTypes(node *ast.File) bool {
	// Named recursive function
	var checkType func(ast.Expr) bool
	checkType = func(expr ast.Expr) bool {
		switch e := expr.(type) {
		case *ast.SelectorExpr:
			if pkgIdent, ok := e.X.(*ast.Ident); ok {
				// Check for "resources" package
				if pkgIdent.Name == "resource" {
					return true
				}
			}
		case *ast.StarExpr:
			// Recursively check the element type for *ast.StarExpr
			return checkType(e.X)
		case *ast.ArrayType:
			// Recursively check the element type for *ast.ArrayType
			return checkType(e.Elt)
		}
		return false
	}

	// Define checkResourcesType using the named recursive function
	checkResourcesType := func(expr ast.Expr) bool {
		return checkType(expr)
	}

	// Check each function declaration for "resources" package types
	for _, decl := range node.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		// Check parameters for "resources" package types
		if fn.Type.Params != nil {
			for _, p := range fn.Type.Params.List {
				if checkResourcesType(p.Type) {
					return true
				}
			}
		}

		// Check return types for "resources" package types
		if fn.Type.Results != nil {
			for _, r := range fn.Type.Results.List {
				if checkResourcesType(r.Type) {
					return true
				}
			}
		}
	}
	return false
}

func (p *Process) formatTableName(tableName string) string {
	parts := strings.Split(tableName, ".")
	// Check if the tableName is split into two parts (schema and table name)
	if len(parts) == 2 {
		// Format and return the table name with each part enclosed in double quotes, and the whole thing in backticks
		return fmt.Sprintf("`\"%s\"`", parts[1])
	}
	// Return the original table name if it's not in the expected format
	return tableName
}

func (p *Process) rawTableName(tableName string) string {
	parts := strings.Split(tableName, ".")
	// Check if the tableName is split into two parts (schema and table name)
	if len(parts) == 2 {
		// Format and return the table name with each part enclosed in double quotes, and the whole thing in backticks
		return fmt.Sprintf("%s", parts[1])
	}
	// Return the original table name if it's not in the expected format
	return tableName
}

func (p *Process) formatFieldNames(fieldNames []string) string {

	// format the fieldnemaes into an array of strings
	var formattedFieldNames []string
	for _, fieldName := range fieldNames {
		formattedFieldNames = append(formattedFieldNames, fmt.Sprintf("\"%s\"", fieldName))
	}

	// return the formatted fieldnames as a string
	return `[]string{` + strings.Join(formattedFieldNames, ",") + `}`
}

func (p *Process) wrapInBackticks(str string) string {
	return fmt.Sprintf("`%s`", str)
}

func (p *Process) insertBacktick() string {
	return "`"
}
