package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

func main() {
    schemaLoader := gojsonschema.NewReferenceLoader("file:///home/francis/go/src/github.com/turntabl/validation of GoLang/user_schema.json")
    documentLoader := gojsonschema.NewReferenceLoader("file:///home/francis/go/src/github.com/turntabl/validation of GoLang/document.json")

    result, err := gojsonschema.Validate(schemaLoader, documentLoader)
    if err != nil {
        panic(err.Error())
    }

    if result.Valid() {
        fmt.Printf("The document is valid\n")
    } else {
        fmt.Printf("The document is not valid. see errors :\n")
        for _, desc := range result.Errors() {
            fmt.Printf("- %s\n", desc)
        }
    }
}