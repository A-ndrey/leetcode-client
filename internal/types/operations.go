package types

import "strings"

type Operation struct {
	Name   string
	Fields []Field
}

func (o Operation) String(sb *strings.Builder) {
	sb.WriteString("query " + o.Name + "{")
	for _, f := range o.Fields {
		f.String(sb)
	}
	sb.WriteString("}")
}
