package {{.pkgName}}

import (
	{{.templImports}}
)

templ {{.templName}}(props *Props){
    @{{.controllerLayout}}.New(
		{{.controllerLayout}}.WithRequest(props.Request),
		{{.controllerLayout}}.WithConfig(props.Config),
	){
		<div>
            <h1>{{.templName}}</h1>
        </div>
	}
}