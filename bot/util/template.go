package util

import (
	"bytes"
	"text/template"
)

type TemplateFunctionProvider interface {
	GetTemplateFunction() template.FuncMap
}

type Parameters map[string]string

var functions = template.FuncMap{
	// creates a slice out of argument
	"makeSlice": func(args ...interface{}) []interface{} {
		return args
	},
	"slice": func(str string, start int, end int) string {
		return str[start:end]
	},
}

// GetTemplateFunctions returns a list of the currently available template functions which can be used in definedCommands or user specific commands
func GetTemplateFunctions() template.FuncMap {
	return functions
}

// RegisterFunctions will add a function to any template renderer
func RegisterFunctions(funcMap template.FuncMap) {
	for name, function := range funcMap {
		functions[name] = function
	}
}

// CompileTemplate pre compiles a template and returns an error if an function is not available etc
func CompileTemplate(temp string) (*template.Template, error) {
	return template.New(temp).Funcs(functions).Parse(temp)
}

// EvalTemplate renders the template
func EvalTemplate(temp *template.Template, params Parameters) (string, error) {
	var buf bytes.Buffer
	err := temp.Execute(&buf, params)

	return buf.String(), err
}
