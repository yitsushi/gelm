package gelm

import "fmt"

type typeAliasField struct {
	GoName     string
	ElmName    string
	ElmType    string
	DecodeType string
	Optional   bool
	Default    string
	List       bool
}

func (fieldDef typeAliasField) DecoderFn() string {
	function := "Dp.required"
	params := fmt.Sprintf("\"%s\" %s", fieldDef.GoName, fieldDef.Decoder())

	if fieldDef.Optional {
		function = "Dp.optional"
		params = fmt.Sprintf("%s %s", params, fieldDef.Default)
	}

	return fmt.Sprintf("|> %s %s", function, params)
}

func (fieldDef typeAliasField) Type() string {
	if fieldDef.List {
		return fmt.Sprintf("List %s", fieldDef.ElmType)
	}

	return fieldDef.ElmType
}

func (fieldDef typeAliasField) Decoder() string {
	var basicDecoder string

	if fieldDef.DecodeType == "custom" {
		basicDecoder = fmt.Sprintf("(Decode.lazy (\\_ -> decode%s))", fieldDef.ElmType)
	} else {
		basicDecoder = fieldDef.DecodeType
	}

	if fieldDef.List {
		return fmt.Sprintf("(Decode.list %s)", basicDecoder)
	}

	return basicDecoder
}
