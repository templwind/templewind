package parser

import (
	"fmt"

	"github.com/templwind/templwind/tools/twctl/pkg/site/ast"
)

// PrintAST prints the entire AST with indentation
func PrintAST(ast ast.SiteAST) {
	for _, s := range ast.Structs {
		PrintStruct(s, 0)
	}
	for _, m := range ast.Modules {
		PrintModule(m, 0)
	}
	for _, s := range ast.Servers {
		PrintServer(s, 0)
	}
}

// PrintStruct prints a struct node with indentation
func PrintStruct(s ast.StructNode, indent int) {
	printIndent(indent)
	fmt.Println("Struct:", s.Name)
	for _, f := range s.Fields {
		PrintStructField(f, indent+1)
	}
}

// PrintStructField prints a struct field with indentation
func PrintStructField(f ast.StructField, indent int) {
	printIndent(indent)
	fmt.Println("Field:", f.Name, f.Type, f.Tags)
}

// PrintServer prints a server node with indentation
func PrintServer(s ast.ServerNode, indent int) {
	printIndent(indent)
	fmt.Println("Server:", s.Name)
	for _, srv := range s.Services {
		PrintService(srv, indent+1)
	}
}

func PrintModule(m ast.ModuleNode, indent int) {
	printIndent(indent)
	fmt.Println("Module:", m.Name)
	PrintMap(m.Attrs, indent+1)
}

func PrintMap(m map[string]string, indent int) {
	for key, value := range m {
		printIndent(indent)
		fmt.Println(key+":", value)
	}
}

// PrintService prints a service node with indentation
func PrintService(s ast.ServiceNode, indent int) {
	printIndent(indent)
	fmt.Println("Service:", s.Name)
	for _, h := range s.Handlers {
		PrintHandler(h, indent+1)
	}
}

// PrintHandler prints a handler node with indentation
func PrintHandler(h ast.HandlerNode, indent int) {
	printIndent(indent)
	fmt.Println("Handler:", h.Name)
	printIndent(indent + 1)
	fmt.Println("Method:", h.Method)
	printIndent(indent + 1)
	fmt.Println("Route:", h.Route)
	if h.Request != "" {
		printIndent(indent + 1)
		fmt.Println("Request:", h.Request)
	}
	if h.Response != "" {
		printIndent(indent + 1)
		fmt.Println("Response:", h.Response)
	}
	if h.Page != nil {
		PrintPage(*h.Page, indent+1)
	}
	if h.Doc != nil {
		PrintDoc(*h.Doc, indent+1)
	}
}

// PrintPage prints a page node with indentation
func PrintPage(p ast.PageNode, indent int) {
	printIndent(indent)
	fmt.Println("Page:", p.Name)
	for key, value := range p.Attrs {
		printIndent(indent + 1)
		fmt.Println(key+":", value)
	}
}

// PrintDoc prints a doc node with indentation
func PrintDoc(d ast.DocNode, indent int) {
	printIndent(indent)
	fmt.Println("Doc:", d.Name)
	for key, value := range d.Attrs {
		printIndent(indent + 1)
		fmt.Println(key+":", value)
	}
}

// printIndent prints dots for indentation
func printIndent(indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print("·  ")
	}
}
