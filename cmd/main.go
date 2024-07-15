package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Component struct {
	Name string
}

var components = []Component{
	{Name: "button"}, {Name: "link"}, {Name: "button-group"}, {Name: "dropdown"}, {Name: "tab"}, {Name: "speed-dial"},
	{Name: "alert"}, {Name: "avatar"}, {Name: "badge"}, {Name: "card"}, {Name: "carousel"},
	{Name: "device-mockups"}, {Name: "file-dropzone"}, {Name: "gallery"}, {Name: "indicator"},
	{Name: "keylabel"}, {Name: "list-group"}, {Name: "pagination"}, {Name: "progress"},
	{Name: "progressbar"}, {Name: "progressradial"}, {Name: "rating"}, {Name: "skeleton"},
	{Name: "spinner"}, {Name: "table"}, {Name: "timeline"}, {Name: "toast"}, {Name: "tooltip"},
	{Name: "typography"}, {Name: "autocomplete"}, {Name: "checkbox"}, {Name: "datepicker"},
	{Name: "filebutton"}, {Name: "file-input"}, {Name: "floating-label"}, {Name: "inputchip"},
	{Name: "input-field"}, {Name: "radio"}, {Name: "range"}, {Name: "search-input"},
	{Name: "select"}, {Name: "slider"}, {Name: "textarea"}, {Name: "toggle"},
	{Name: "alert"}, {Name: "toast"}, {Name: "app-bar"}, {Name: "app-header"}, {Name: "app-shell"},
	{Name: "app-rail"}, {Name: "drawer"}, {Name: "footer"}, {Name: "header"}, {Name: "sidebar"},
	{Name: "device-mockups"}, {Name: "breadcrumb"}, {Name: "bottom-navigation"}, {Name: "dropdown"},
	{Name: "mega-menu"}, {Name: "navbar"}, {Name: "sidenav"}, {Name: "tabs"},
	{Name: "blockquote"}, {Name: "heading"}, {Name: "hr"}, {Name: "image"}, {Name: "link"},
	{Name: "list"}, {Name: "paragraph"}, {Name: "text"},
}

func toCamelCase(input string) string {
	words := strings.FieldsFunc(input, func(r rune) bool {
		return r == '-' || r == '_'
	})
	for i := range words {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

func createFile(path string, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

func createStructure(basePath string) error {
	for _, comp := range components {
		componentPath := filepath.Join(basePath, "components", comp.Name)
		err := os.MkdirAll(componentPath, os.ModePerm)
		if err != nil {
			return err
		}

		// Create props.go
		propsContent := fmt.Sprintf(`package %s

import (
	"github.com/a-h/templ"
	"github.com/templwind/templwind"
)

// Props for the %s component
type Props struct {
}

// New creates a new component
func New(props ...templwind.OptFunc[Props]) templ.Component {
	return templwind.New(defaultProps, tpl, props...)
}

// NewWithProps creates a new component with the given props
func NewWithProps(props *Props) templ.Component {
	return templwind.NewWithProps(tpl, props)
}

// WithProps builds the props with the given options
func WithProps(props ...templwind.OptFunc[Props]) *Props {
	return templwind.WithProps(defaultProps, props...)
}

func defaultProps() *Props {
	return &Props{}
}`, strings.ReplaceAll(comp.Name, "-", ""), comp.Name)
		err = createFile(filepath.Join(componentPath, fmt.Sprintf("%s.go", comp.Name)), propsContent)
		if err != nil {
			return err
		}

		// Create <component>.templ
		templContent := fmt.Sprintf(`package %s

templ tpl(props *Props) {
	<tw-%s>
		<div>component</div>
	</tw-%s>
}`, strings.ReplaceAll(comp.Name, "-", ""), comp.Name, comp.Name)
		err = createFile(filepath.Join(componentPath, fmt.Sprintf("%s.templ", comp.Name)), templContent)
		if err != nil {
			return err
		}

		// Create <component>.ts
		tsContent := fmt.Sprintf(`class %s extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("%s connected");
	}
}

customElements.define("tw-%s", %s);`,
			toCamelCase(comp.Name), toCamelCase(comp.Name), comp.Name, toCamelCase(comp.Name))
		err = createFile(filepath.Join(componentPath, fmt.Sprintf("%s.ts", comp.Name)), tsContent)
		if err != nil {
			return err
		}

		// Create <component>.scss
		scssContent := fmt.Sprintf(`.%s {
  // Add your styles here
}`, comp.Name)
		err = createFile(filepath.Join(componentPath, fmt.Sprintf("%s.scss", comp.Name)), scssContent)
		if err != nil {
			return err
		}

		// Create index.ts
		indexContent := fmt.Sprintf(`import './%s.scss';
import './%s';`, comp.Name, comp.Name)
		err = createFile(filepath.Join(componentPath, "index.ts"), indexContent)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	os.Exit(0)
	basePath := ""
	err := createStructure(basePath)
	if err != nil {
		fmt.Println("Error creating structure:", err)
		return
	}
	fmt.Println("Directory structure created successfully!")
}
