package spec

import (
	"fmt"
)

// Define the constants used in the spec
const RoutePrefixKey = "prefix"

// Define the types used in the spec
type (
	// Doc represents documentation strings
	Doc []string

	// Annotation defines key-value properties for annotations
	Annotation struct {
		Properties map[string]string
	}

	// SiteSpec describes a Site file
	SiteSpec struct {
		Name    string
		Types   []Type
		Servers []Server
		Modules []Module
	}

	// Module describes an external module
	Module struct {
		Name   string
		Source string
		Prefix string
		Attr   map[string]string
	}

	// Server describes a server block with its services
	Server struct {
		Annotation Annotation
		Services   []Service
	}

	// Service describes a Site service with its handlers
	Service struct {
		Name     string
		Handlers []Handler
	}

	// Handler describes a Site handler
	Handler struct {
		Name           string
		Method         string
		Route          string
		RequestType    Type
		ResponseType   Type
		Page           *Page
		Doc            *DocNode
		HandlerDoc     Doc
		HandlerComment Doc
		DocAnnotation  Annotation
	}

	// Page represents a page in a handler
	Page struct {
		Annotation Annotation
	}

	// DocNode represents a doc block in a handler
	DocNode struct {
		Annotation Annotation
	}

	// Type defines the types used in the Site spec
	Type interface {
		GetName() string
		GetFields() []Field
		GetComments() []string
		GetDocuments() []string
	}

	// StructType describes a structure type
	StructType struct {
		Name    string
		Fields  []Field
		Docs    Doc
		Comment Doc
	}

	// Field describes the field of a structure
	Field struct {
		Name    string
		Type    string
		Tag     string
		Comment string
		Docs    Doc
	}

	// PrimitiveType describes a primitive type
	PrimitiveType struct {
		Name string
	}

	// MapType describes a map type
	MapType struct {
		Name  string
		Key   string
		Value Type
	}

	// ArrayType describes an array type
	ArrayType struct {
		Name  string
		Value Type
	}

	// PointerType describes a pointer type
	PointerType struct {
		Name string
		Type Type
	}

	// InterfaceType describes an interface type
	InterfaceType struct {
		Name string
	}
)

// NewAnnotation creates a new annotation
func NewAnnotation(properties map[string]string) Annotation {
	return Annotation{
		Properties: properties,
	}
}

// NewDocNode creates a new doc node
func NewDocNode(annotation Annotation) *DocNode {
	return &DocNode{
		Annotation: annotation,
	}
}

// NewPage creates a new page node
func NewPage(annotation Annotation) *Page {
	return &Page{
		Annotation: annotation,
	}
}

// NewServer creates a new server node
func NewServer(annotation Annotation) *Server {
	return &Server{
		Annotation: annotation,
		Services:   []Service{},
	}
}

// NewService creates a new service node
func NewService(name string) *Service {
	return &Service{
		Name:     name,
		Handlers: []Handler{},
	}
}

// NewModule creates a new module node
func NewModule(name string, attr map[string]string) Module {
	return Module{
		Name: name,
		Attr: attr,
	}
}

// NewHandler creates a new handler node
func NewHandler(name, method, route string, requestType, responseType interface{}, page *Page, doc *DocNode) *Handler {

	var (
		reqType Type
		resType Type
	)

	if requestType != nil {
		reqType = requestType.(Type)
	}

	if responseType != nil {
		resType = responseType.(Type)
	}

	return &Handler{
		Name:         name,
		Method:       method,
		Route:        route,
		RequestType:  reqType,
		ResponseType: resType,
		Page:         page,
		Doc:          doc,
	}
}

// NewStructType creates a new struct type
func NewStructType(name string, fields []Field, docs, comment Doc) *StructType {
	return &StructType{
		Name:    name,
		Fields:  fields,
		Docs:    docs,
		Comment: comment,
	}
}

// NewPrimitiveType creates a new primitive type
func NewPrimitiveType(name string) *PrimitiveType {
	return &PrimitiveType{
		Name: name,
	}
}

// NewMapType creates a new map type
func NewMapType(key string, value Type) *MapType {
	return &MapType{
		Key:   key,
		Value: value,
	}
}

// NewArrayType creates a new array type
func NewArrayType(value Type) *ArrayType {
	return &ArrayType{
		Value: value,
	}
}

// NewPointerType creates a new pointer type
func NewPointerType(t Type) *PointerType {
	return &PointerType{
		Type: t,
	}
}

// NewInterfaceType creates a new interface type
func NewInterfaceType(name string) *InterfaceType {
	return &InterfaceType{
		Name: name,
	}
}

// annotation methods
// GetAnnotation returns the value by specified key from @server
func (s Server) GetAnnotation(key string) string {
	if s.Annotation.Properties == nil {
		return ""
	}

	return s.Annotation.Properties[key]
}

// Methods to implement the Type interface for StructType
func (t *StructType) GetName() string {
	return t.Name
}

func (t *StructType) GetComments() []string {
	return []string(t.Comment)
}

func (t *StructType) GetDocuments() []string {
	return []string(t.Docs)
}

func (t *StructType) GetFields() []Field {
	return t.Fields
}

// Methods to implement the Type interface for Field
func (t *Field) GetName() string {
	return t.Name
}

func (t *Field) GetComments() []string {
	return []string{t.Comment}
}

func (t *Field) GetDocuments() []string {
	return []string(t.Docs)
}

func (t *Field) GetFields() []Field {
	return nil
}

// Methods to implement the Type interface for PrimitiveType
func (t *PrimitiveType) GetName() string {
	return t.Name
}

func (t *PrimitiveType) GetComments() []string {
	return nil
}

func (t *PrimitiveType) GetDocuments() []string {
	return nil
}

func (t *PrimitiveType) GetFields() []Field {
	return nil
}

// Methods to implement the Type interface for MapType
func (t *MapType) GetName() string {
	return fmt.Sprintf("map[%s]%s", t.Key, t.Value.GetName())
}

func (t *MapType) GetComments() []string {
	return nil
}

func (t *MapType) GetDocuments() []string {
	return nil
}

func (t *MapType) GetFields() []Field {
	return nil
}

// Methods to implement the Type interface for ArrayType
func (t *ArrayType) GetName() string {
	return fmt.Sprintf("[]%s", t.Value.GetName())
}

func (t *ArrayType) GetComments() []string {
	return nil
}

func (t *ArrayType) GetDocuments() []string {
	return nil
}

func (t *ArrayType) GetFields() []Field {
	return nil
}

// Methods to implement the Type interface for PointerType
func (t *PointerType) GetName() string {
	return fmt.Sprintf("*%s", t.Type.GetName())
}

func (t *PointerType) GetComments() []string {
	return nil
}

func (t *PointerType) GetDocuments() []string {
	return nil
}

func (t *PointerType) GetFields() []Field {
	return nil
}

// Methods to implement the Type interface for InterfaceType
func (t *InterfaceType) GetName() string {
	return t.Name
}

func (t *InterfaceType) GetComments() []string {
	return nil
}

func (t *InterfaceType) GetDocuments() []string {
	return nil
}

func (t *InterfaceType) GetFields() []Field {
	return nil
}
