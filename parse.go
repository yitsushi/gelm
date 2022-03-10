package gelm

import (
	"reflect"
	"strings"
)

const (
	tagName = "elm"
	kvParts = 2
)

func parse(obj interface{}) typeAlias {
	dataType := reflect.TypeOf(obj)

	typeDef := typeAlias{
		Name:   dataType.Name(),
		Fields: []typeAliasField{},
	}

	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		tag := strings.Split(field.Tag.Get(tagName), ",")
		elmT, decodeT, isList := goTypeToElmType(field.Type)

		if len(tag) < 1 {
			continue
		}

		elmName := tag[0]
		optional := false
		defaultValue := ""

		if len(tag) > 1 {
			parts := strings.SplitN(tag[1], "=", kvParts)
			if parts[0] == "optional" {
				optional = true
				defaultValue = parts[1]
			}
		}

		fieldDef := typeAliasField{
			GoName:     field.Name,
			ElmName:    elmName,
			ElmType:    elmT,
			DecodeType: decodeT,
			Optional:   optional,
			Default:    defaultValue,
			List:       isList,
		}

		typeDef.Fields = append(typeDef.Fields, fieldDef)
	}

	return typeDef
}

func goTypeToElmType(dataType reflect.Type) (string, string, bool) {
	switch dataType.Kind().String() {
	case "bool":
		return "Bool", "Decode.bool", false
	case "string":
		return "String", "Decode.string", false
	case "int", "int64", "int32", "uint", "uint32", "uint64":
		return "Int", "Decode.int", false
	case "float", "float32", "float64":
		return "Float", "Decode.float", false
	case "struct":
		return dataType.Name(), "custom", false
	case "slice":
		elmT, decodeT, _ := goTypeToElmType(dataType.Elem())

		return elmT, decodeT, true
	default:
		return "", "", false
	}
}
