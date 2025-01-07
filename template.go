//go:build ignore

// go generate
package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"os"
	"strings"
	"text/template"

	"github.com/fatih/color"
	"github.com/iancoleman/strcase"
)

var (
	red    = color.New(color.FgRed)
	green  = color.New(color.FgGreen)
	blue   = color.New(color.FgBlue)
	yellow = color.New(color.FgYellow)
	grey   = color.New(color.FgHiBlack)
)

type typeOfResource int
type typeOfTemplate int

func (t typeOfResource) String() string {
	switch t {
	case simpleType:
		return "simple"
	case arrayType:
		return "array"
	case nestedType:
		return "nested"
	case objectType:
		return "object"
	case singleNestedType:
		return "singleNested"
	}
	return "unknown"
}

const (
	simpleType typeOfResource = iota
	arrayType
	nestedType
	objectType
	singleNestedType
)

const (
	templateType typeOfTemplate = iota
	templateValue
	templateTypeTest
	templateValueTest
)

var (
	templateStructs = map[typeOfResource]templateStruct{
		simpleType: {
			tOResource:     simpleType,
			tOTemplate:     templateType,
			files:          templateFilesMap[simpleType],
			terraformTypes: terraformTypesMap[simpleType],
		},
		arrayType: {
			tOResource:     arrayType,
			tOTemplate:     templateType,
			files:          templateFilesMap[arrayType],
			terraformTypes: terraformTypesMap[arrayType],
		},
		nestedType: {
			tOResource:     nestedType,
			tOTemplate:     templateType,
			files:          templateFilesMap[nestedType],
			terraformTypes: terraformTypesMap[nestedType],
		},
		objectType: {
			tOResource:     objectType,
			tOTemplate:     templateType,
			files:          templateFilesMap[objectType],
			terraformTypes: terraformTypesMap[objectType],
		},
		singleNestedType: {
			tOResource:     singleNestedType,
			tOTemplate:     templateType,
			files:          templateFilesMap[singleNestedType],
			terraformTypes: terraformTypesMap[singleNestedType],
		},
	}
	templateFilesMap = map[typeOfResource]templateFiles{
		simpleType: {
			templateType:  templateSimpleType,
			templateValue: templateSimpleValue,
		},
		arrayType: {
			templateType:  templateArrayType,
			templateValue: templateArrayValue,
		},
		nestedType: {
			templateType:  templateNestedType,
			templateValue: templateNestedValue,
		},
		objectType: {
			templateType:  templateObjectType,
			templateValue: templateObjectValue,
		},
		singleNestedType: {
			templateType:  templateSingleNestedType,
			templateValue: templateSingleNestedValue,
		},
	}
	terraformTypesMap = map[typeOfResource][]string{
		simpleType:       {"string", "bool", "float64", "int64", "number"},
		arrayType:        {"list", "set", "map"},
		nestedType:       {"list_nested", "set_nested", "map_nested"},
		objectType:       {"object"},
		singleNestedType: {"single_nested"},
	}
)

type (
	templateStruct struct {
		tOResource     typeOfResource
		tOTemplate     typeOfTemplate
		files          templateFiles
		terraformTypes []string
	}

	templateFiles map[typeOfTemplate]string

	templateInfos struct {
		TypeName string
		fileName string
	}
)

//go:embed templates/common.go.tmpl
var templateCommon string

// * Simple
//
//go:embed templates/simple_type.go.tmpl
var templateSimpleType string

//go:embed templates/simple_value.go.tmpl
var templateSimpleValue string

// * Array
//
//go:embed templates/array_type.go.tmpl
var templateArrayType string

//go:embed templates/array_value.go.tmpl
var templateArrayValue string

// * Nested
//
//go:embed templates/nested_type.go.tmpl
var templateNestedType string

//go:embed templates/nested_value.go.tmpl
var templateNestedValue string

// * Object
//
//go:embed templates/object_type.go.tmpl
var templateObjectType string

//go:embed templates/object_value.go.tmpl
var templateObjectValue string

// * Single nested
//
//go:embed templates/single_nested_type.go.tmpl
var templateSingleNestedType string

//go:embed templates/single_nested_value.go.tmpl
var templateSingleNestedValue string

var templateFuncs = template.FuncMap{
	"golangType": func(typeName string) string {
		switch strings.ToLower(typeName) {
		case "string":
			return "string"
		case "bool":
			return "bool"
		case "float64":
			return "float64"
		case "int64":
			return "int64"
		case "number":
			return "*big.Float"
		case "object":
			return "any"
		case "list", "list_nested":
			return "[]interface{}"
		case "set", "set_nested":
			return "[]interface{}"
		case "map", "map_nested":
			return "map[string]interface{}"
		case "single_nested":
			return "map[string]interface{}"
		default:
			return ""
		}
	},
}

func main() {
	yellow.Println("Starting generation of supertypes")
	yellow.Println("=================================")

	for _, t := range templateStructs {
		blue.Printf("\n-> Generating %s type\n", t.tOResource.String())
		for _, f := range t.terraformTypes {
			generateTemplates(f, t)
		}
	}
}

func generateTemplates(typeName string, tStruct templateStruct) {
	grey.Printf("  * Generating %s\n", typeName)
	infos := templateInfos{
		TypeName: strcase.ToCamel(typeName),
		fileName: typeName,
	}

	if strings.HasSuffix(typeName, "_nested") {
		infos.TypeName = strcase.ToCamel(strings.TrimSuffix(typeName, "_nested"))
	}

	// * Generate Type file
	if err := generateTemplate(infos, templateType, tStruct); err != nil {
		red.Printf("\nerror from template execute : %v\n", err)
		os.Exit(1)
	}

	// * Generate Value file
	if err := generateTemplate(infos, templateValue, tStruct); err != nil {
		red.Printf("\nerror from template execute : %v\n", err)
		os.Exit(1)
	}

	// * Generate Type test file
	if _, ok := tStruct.files[templateTypeTest]; ok {
		if err := generateTemplate(infos, templateTypeTest, tStruct); err != nil {
			red.Printf("\nerror from template execute : %v\n", err)
			os.Exit(1)
		}
	}

	// * Generate Value test file
	if _, ok := tStruct.files[templateValueTest]; ok {
		if err := generateTemplate(infos, templateValueTest, tStruct); err != nil {
			red.Printf("\nerror from template execute : %v\n", err)
			os.Exit(1)
		}
	}
}

func generateTemplate(tInfos templateInfos, tOTemplate typeOfTemplate, t templateStruct) error {
	var tpl bytes.Buffer

	grey.Printf("    Processing ")
	// parse the template
	tmpl, err := template.New(fmt.Sprintf("template-%s", tInfos.fileName)).Funcs(templateFuncs).Parse(concatTemplate(t.files[tOTemplate]))
	if err != nil {
		return fmt.Errorf("error from template parse : %v\n", err)
	}

	// execute the template
	if err := tmpl.Execute(&tpl, tInfos); err != nil {
		return fmt.Errorf("error from template execute : %v\n", err)
	}

	// format the code
	formattedTmpl, err := format.Source(tpl.Bytes())
	if err != nil {
		return fmt.Errorf("error from format code : %v\n", err)
	}

	var filename string

	switch tOTemplate {
	case templateType:
		filename = tInfos.fileName + "_type.go"
	case templateValue:
		filename = tInfos.fileName + "_value.go"
	case templateTypeTest:
		filename = tInfos.fileName + "_type_test.go"
	case templateValueTest:
		filename = tInfos.fileName + "_value_test.go"
	default:
		// No Default case because we want to be sure to have a filename and const contraint the type.
		return fmt.Errorf("unknown type of template")
	}

	filename = strings.ToLower(filename)

	yellow.Printf(filename)

	// write the file
	if err := os.WriteFile(filename, formattedTmpl, 0644); err != nil {
		return fmt.Errorf("write to file error : %v\n", err)
	}

	green.Println(" ✔")

	return nil
}

func concatTemplate(template string) string {
	return templateCommon + template
}
