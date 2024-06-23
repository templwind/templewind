package parser

import (
	"fmt"
	"strings"

	"github.com/templwind/templwind/tools/twctl/pkg/site/ast"
	"github.com/templwind/templwind/tools/twctl/pkg/site/lexer"
	"github.com/templwind/templwind/tools/twctl/pkg/site/spec"
)

// Parser represents a parser
type Parser struct {
	lexer     *lexer.Lexer
	curToken  lexer.Token
	peekToken lexer.Token
}

// NewParser initializes a new parser
func NewParser(filename string) (*Parser, error) {
	lex, err := lexer.NewLexer(filename)
	if err != nil {
		return nil, err
	}
	p := &Parser{lexer: lex}
	p.nextToken()
	p.nextToken() // read two tokens, so curToken and peekToken are both set
	return p, nil
}

// nextToken advances to the next token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

// Parse returns the parsed AST
func (p *Parser) Parse() ast.SiteAST {
	var astTree ast.SiteAST

	for p.curToken.Type != lexer.EOF {
		// fmt.Println("TOKEN", p.curToken.Literal, p.curToken.Type)
		switch p.curToken.Type {
		case lexer.AT_TYPE:
			astTree.Structs = append(astTree.Structs, p.parseStruct())
		case lexer.AT_SERVER:
			astTree.Servers = append(astTree.Servers, p.parseServer())
		// case lexer.AT_MENUS:
		// 	astTree.Menus = append(astTree.Menus, p.parseMenus())
		case lexer.AT_MODULE:
			astTree.Modules = append(astTree.Modules, p.parseModule())
		}
		p.nextToken()
	}
	return astTree
}

func (p *Parser) parseStruct() ast.StructNode {
	node := ast.StructNode{
		BaseNode: ast.NewBaseNode(ast.NodeTypeStruct, p.curToken.Literal),
		Fields:   []ast.StructField{},
	}

	p.nextToken() // advance
	for p.curToken.Type != lexer.CLOSE_BRACE {
		field := p.parseStructField()
		node.Fields = append(node.Fields, field)
		p.nextToken()
	}
	return node
}

func (p *Parser) parseStructField() ast.StructField {
	parts := strings.Fields(p.curToken.Literal)
	if len(parts) < 2 {
		if len(parts) == 1 { // Handle embedded structs
			return ast.StructField{
				Name: parts[0],
				Type: parts[0],
				Tags: "",
			}
		}
		return ast.StructField{}
	}
	name := parts[0]
	fieldType := parts[1]
	tags := ""
	if len(parts) > 2 {
		tags = strings.Join(parts[2:], " ")
	}
	return ast.StructField{
		Name: name,
		Type: fieldType,
		Tags: tags,
	}
}

func (p *Parser) parseServer() ast.ServerNode {
	node := ast.NewServerNode(p.parseAttributes())

	// fmt.Println("STARTING SERVER TOKEN", p.curToken.Literal, p.curToken.Type)

	for p.curToken.Type != lexer.CLOSE_BRACE {
		// fmt.Println("Parsing server", p.curToken.Literal, p.curToken.Type)
		if p.curToken.Type == lexer.AT_SERVICE {
			node.Services = append(node.Services, p.parseService())

			// fmt.Println("Parsing server", node)
			// fmt.Println("TOKEN", p.curToken.Literal, p.curToken.Type)
			// os.Exit(1)
		}
	}

	return *node
}

func (p *Parser) parseService() ast.ServiceNode {
	node := ast.NewServiceNode(p.curToken.Literal)
	p.nextToken() // skip 'service'
	activeHandler := ast.HandlerNode{}

	for p.curToken.Type == lexer.AT_PAGE ||
		p.curToken.Type == lexer.AT_DOC ||
		p.curToken.Type == lexer.AT_HANDLER {
		if p.curToken.Type == lexer.AT_HANDLER {
			// fmt.Println("HANDLER", p.curToken.Literal, p.curToken.Type)
			node.Handlers = append(node.Handlers, p.parseHandler(activeHandler))
			activeHandler = ast.HandlerNode{}
		} else if p.curToken.Type == lexer.AT_PAGE {
			activeHandler.Page = p.parsePage()
			// fmt.Println("PAGE", p.curToken.Literal, p.peekToken.Literal, p.curToken.Type)
			continue
		} else if p.curToken.Type == lexer.AT_DOC {
			// fmt.Println("DOC", p.curToken.Literal, p.curToken.Type)
			activeHandler.Doc = p.parseDoc()
			continue
		}
	}
	return *node
}

func (p *Parser) parseHandler(handler ast.HandlerNode) ast.HandlerNode {
	name := p.curToken.Literal
	p.nextToken()

	literal := strings.Replace(p.curToken.Literal, "(", " (", -1)
	literal = strings.Join(strings.Fields(literal), " ")

	parts := strings.Fields(literal)

	handler.BaseNode = ast.NewBaseNode(ast.NodeTypeHandler, name)
	if len(parts) < 2 {
		return handler
	}
	handler.Method = parts[0]
	handler.Route = parts[1]

	if len(parts) > 2 {
		handler.Request = parts[2][1 : len(parts[2])-1]
		handler.RequestType = spec.NewStructType(handler.Request, nil, nil, nil)
	}
	if len(parts) > 3 {
		handler.Response = parts[4][1 : len(parts[4])-1]
		handler.ResponseType = spec.NewStructType(handler.Response, nil, nil, nil)
	}

	p.nextToken()
	// fmt.Println("CURRENT TOKEN", p.curToken.Literal)

	// fmt.Println("HANDLER", handler)
	// os.Exit(1)
	return handler
}

func (p *Parser) parsePage() *ast.PageNode {
	attrs := p.parseAttributes()
	return ast.NewPageNode(attrs)
}

func (p *Parser) parseDoc() *ast.DocNode {
	attrs := p.parseAttributes()
	return ast.NewDocNode(attrs)
}

func (p *Parser) parseAttributes() map[string]string {
	attrs := make(map[string]string)
	p.nextToken()
	for p.curToken.Type != lexer.CLOSE_PAREN {
		parts := strings.SplitN(p.curToken.Literal, ":", 2)
		if len(parts) == 2 {
			attrs[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
		// fmt.Println("PARTS", parts, p.curToken.Literal, attrs)
		p.nextToken()
	}
	p.nextToken() // skip ')'
	return attrs
}

func (p *Parser) skipBlock() {
	openBraces := 1
	for openBraces > 0 {
		p.nextToken()
		if p.curToken.Type == lexer.OPEN_BRACE {
			openBraces++
		} else if p.curToken.Type == lexer.CLOSE_BRACE {
			openBraces--
		}
	}
}

func (p *Parser) parseMenus() ast.MenusNode {
	node := ast.NewMenusNode(ast.NodeTypeMenus)
	for p.curToken.Type != lexer.CLOSE_PAREN {
		if p.curToken.Type == lexer.AT_MENU {
			node.Menus = append(node.Menus, p.parseMenu())
		} else {
			p.nextToken()
		}
	}
	return *node
}

func (p *Parser) parseMenu() ast.MenuNode {
	node := ast.NewMenuNode(ast.NodeTypeMenu)
	p.nextToken()
	for p.curToken.Type != lexer.CLOSE_BRACE {
		attrs := p.parseAttributes()
		fmt.Println("MENU ATTRS", attrs)
	}

	return *node
}

func (p *Parser) parseModule() ast.ModuleNode {
	attrs := make(map[string]string)
	for p.curToken.Type != lexer.CLOSE_PAREN {
		parts := strings.SplitN(p.curToken.Literal, ":", 2)
		if len(parts) == 2 {
			attrs[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
		p.nextToken()
	}

	// fmt.Println("MODULE ATTRS", attrs)
	node := ast.NewModuleNode(attrs["name"], attrs)
	node.Attrs = attrs
	node.Source = attrs["source"]
	node.Prefix = attrs["prefix"]
	return *node
}
