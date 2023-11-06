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

var json = "{\"enable\":true,\"supportedModelFormats\":[{\"name\":\"pytorch\",\"version\":\"1\",\"priority\":1}],\"template\":{\"metadata\":{},\"spec\":{\"containers\":[{\"name\":\"user-container\",\"image\":\"tsz.io/t9kmirror/pytorch-serve:0.4.2-cpu\",\"args\":[\"torchserve\",\"--start\",\"--model-store=/var/lib/t9k/model\"],\"resources\":{\"limits\":{\"cpu\":\"200m\",\"memory\":\"200Mi\"}}}]}}}"

var yaml = `apiVersion: tensorstack.dev/v1beta1
kind: ClusterMLServiceRuntime
metadata:
  name: t9k-torchserve
spec:
  enable: true
  supportedModelFormats:
    - name: pytorch
      version: "1"
      priority: 1
  template:
    spec:
      containers:
      - name: user-container
        image: torchserve:latest
        args:
          - torchserve
          - --start
          - --model-store=/var/lib/t9k/model
          {{- if .MODEL_PATH}}- --models {{.MODEL_PATH}}{{end}}
        resources:
          limits:
            cpu: "200m"
            memory: 200Mi
`

func main() {

	tmpl, err := template.New("test").Parse(yaml)
	if err != nil {
		fmt.Println(err)
	}
	m := map[string]string{
		"MODEL_PATH": "mnist=model.mar",
	}
	tmpl.Execute(os.Stdout, m)
}