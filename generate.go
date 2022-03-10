package gelm

import (
	"fmt"
	"strings"
)

const (
	indent1 = 1
	indent2 = 2
)

func indent(n int, line string) string {
	output := ""

	for ; n > 0; n-- {
		output += "    "
	}

	return output + line
}

func generateTypeAlias(def typeAlias) string {
	lines := []string{fmt.Sprintf("type alias %s =", def.Name)}

	for idx, fieldDef := range def.Fields {
		prefix := ","

		if idx == 0 {
			prefix = "{"
		}

		lines = append(
			lines,
			indent(indent1, fmt.Sprintf("%s %s : %s", prefix, fieldDef.ElmName, fieldDef.Type())),
		)
	}

	lines = append(lines, indent(indent1, "}"))

	return strings.Join(lines, "\n")
}

func generateDecoder(def typeAlias) string {
	lines := []string{
		fmt.Sprintf("decode%s : Decode.Decoder %s", def.Name, def.Name),
		fmt.Sprintf("decode%s =", def.Name),
		indent(1, fmt.Sprintf("Decode.succeed %s", def.Name)),
	}

	for _, fieldDef := range def.Fields {
		lines = append(lines, indent(indent2, fieldDef.DecoderFn()))
	}

	return strings.Join(lines, "\n")
}
