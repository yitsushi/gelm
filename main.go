package gelm

import (
	"fmt"
)

// Generate Elm type aliases with decoder functions as a module.
func Generate(moduleName string, types ...interface{}) []byte {
	output := []byte(fmt.Sprintf(`-- This file was generated with github.com/yitsushi/gelm
-- Do not edit this file unless you know what you are doing.


module %s exposing (..)

import Json.Decode as Decode
import Json.Decode.Pipeline as Dp`, moduleName))

	for _, obj := range types {
		typeDef := parse(obj)

		output = append(output, []byte{'\n', '\n', '\n'}...)
		output = append(output, generateTypeAlias(typeDef)...)
	}

	for _, obj := range types {
		typeDef := parse(obj)

		output = append(output, []byte{'\n', '\n', '\n'}...)
		output = append(output, generateDecoder(typeDef)...)
	}

	output = append(output, '\n')

	return output
}