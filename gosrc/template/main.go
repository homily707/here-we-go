package main

import (
	"fmt"
	"os"
	"text/template"
)

var parseText = `{
	format: json,
	bareValue: {{ .bareValue }},
	trueif: {{ if .condition }} "if true,{{.bareValue}}" {{ else }} if false,{{.bareValue}} {{end}},
	falseif: {{ if .falsecondition }} if true,{{.bareValue}} {{ else }} if false,{{.bareValue}} {{end}},
	range: [
		{{range $index, $element := . }}{{$index}}:{{$element}},
		{{end}}
	]
}
`

func main() {

	tmpl, err := template.New("test").Parse(parseText)
	if err != nil {
		fmt.Println(err)
	}
	m := map[string]string{
		"bareValue": "bare",
		"condition": "true",
		"falsecondition": "0",
	}
	tmpl.Execute(os.Stdout, m)
}