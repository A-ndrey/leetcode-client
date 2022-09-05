package types

import (
	"fmt"
	"strings"
)

type Field struct {
	Name   string
	Args   map[string]any
	Fields []Field
}

func (f Field) String(sb *strings.Builder) {
	sb.WriteString(f.Name)
	if len(f.Args) > 0 {
		sb.WriteString("(")
		for k, v := range f.Args {
			sb.WriteString(fmt.Sprintf("%s: \"%v\", ", k, v))
		}
		sb.WriteString(")")
	}
	if len(f.Fields) == 0 {
		sb.WriteString(" ")
		return
	}
	sb.WriteString("{")
	for _, f := range f.Fields {
		f.String(sb)
	}
	sb.WriteString("}")
}
