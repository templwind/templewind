package ast

// Define the node types
const (
	NodeTypeStruct    = "struct"
	NodeTypeServer    = "@server"
	NodeTypeService   = "service"
	NodeTypeHandler   = "@handler"
	NodeTypePage      = "@page"
	NodeTypeDoc       = "@doc"
	NodeTypeAttribute = "attribute"
	NodeTypeMenus     = "@menus"
	NodeTypeMenu      = "@menu"
	NodeTypeModules   = "@modules"
	NodeTypeModule    = "module"
)

// BaseNode represents a base AST node with a generic type.
type BaseNode struct {
	Type     string
	Name     string
	Children []BaseNode
	Attrs    map[string]string
}

// StructField represents a field in a struct.
type StructField struct {
	Name string
	Type string
	Tags string
}

// StructNode represents a custom struct.
type StructNode struct {
	BaseNode
	Fields []StructField
}

// ServerNode represents a server block with its services.
type ServerNode struct {
	BaseNode
	Services []ServiceNode
}

// ServiceNode represents a service with its handlers.
type ServiceNode struct {
	BaseNode
	Handlers []HandlerNode
}

// HandlerNode represents a handler in a service.
type HandlerNode struct {
	BaseNode
	Method       string
	Route        string
	Request      string
	RequestType  interface{}
	Response     string
	ResponseType interface{}
	Page         *PageNode
	Doc          *DocNode
}

// PageNode represents a page in a service.
type PageNode struct {
	BaseNode
}

// DocNode represents a doc block in a handler.
type DocNode struct {
	BaseNode
}

// MenuItem represents an item in a menu.
type MenuItem struct {
	Title string
	URL   string
	Icon  string
}

// MenusNode represents 1 or more menus.
type MenusNode struct {
	BaseNode
	Menus []MenuNode
}

// MenuNode represents a menu with its items.
type MenuNode struct {
	BaseNode
	Items []MenuItem
}

// ModulesNode represents one or more modules.
type ModulesNode struct {
	BaseNode
	Modules []ModuleNode
}

// ModuleNode represents a module with its configuration.
type ModuleNode struct {
	BaseNode
	Source string
	Prefix string
}

// SiteAST represents the entire Abstract Syntax Tree.
type SiteAST struct {
	Name    string
	Structs []StructNode
	Servers []ServerNode
	Menus   []MenusNode
	Modules []ModuleNode
}

// NewBaseNode creates a new base node.
func NewBaseNode(nodeType, name string) BaseNode {
	return BaseNode{
		Type:     nodeType,
		Name:     name,
		Children: []BaseNode{},
		Attrs:    map[string]string{},
	}
}

// NewDocNode creates a new doc node.
func NewDocNode(attrs map[string]string) *DocNode {
	return &DocNode{
		BaseNode: BaseNode{
			Type:  NodeTypeDoc,
			Name:  "doc",
			Attrs: attrs,
		},
	}
}

// NewPageNode creates a new page node.
func NewPageNode(attrs map[string]string) *PageNode {
	return &PageNode{
		BaseNode: BaseNode{
			Type:  NodeTypePage,
			Name:  "page",
			Attrs: attrs,
		},
	}
}

// NewServerNode creates a new server node.
func NewServerNode(attrs map[string]string) *ServerNode {
	return &ServerNode{
		BaseNode: BaseNode{
			Type:  NodeTypeServer,
			Name:  "server",
			Attrs: attrs,
		},
		Services: []ServiceNode{},
	}
}

// NewServiceNode creates a new service node.
func NewServiceNode(name string) *ServiceNode {
	return &ServiceNode{
		BaseNode: BaseNode{
			Type: NodeTypeService,
			Name: name,
		},
		Handlers: []HandlerNode{},
	}
}

// NewMenusNode creates a new menus node.
func NewMenusNode(name string) *MenusNode {
	return &MenusNode{
		BaseNode: BaseNode{
			Type: NodeTypeMenu,
			Name: name,
		},
		Menus: []MenuNode{},
	}
}

// NewMenuNode creates a new menu node.
func NewMenuNode(name string) *MenuNode {
	return &MenuNode{
		BaseNode: BaseNode{
			Type: NodeTypeMenu,
			Name: name,
		},
		Items: []MenuItem{},
	}
}

// NewModuleNode creates a new module node.
func NewModuleNode(name string, attr map[string]string) *ModuleNode {
	return &ModuleNode{
		BaseNode: BaseNode{
			Type:  NodeTypeModule,
			Name:  name,
			Attrs: attr,
		},
	}
}
