package parser

import "strings"

type Package struct {
	Name  string
	Files []*File
}

func (p *Package) Print(depth int) string {
	prefix := strings.Repeat("\t", depth)
	var sb strings.Builder

	sb.WriteString(prefix)
	sb.WriteString("package ")
	sb.WriteString(p.Name)
	sb.WriteString("\n")
	sb.WriteString(prefix)
	sb.WriteString("Files: \n")
	for _, f := range p.Files {
		sb.WriteString(prefix)
		sb.WriteString("  - ")
		sb.WriteString(f.Path)
		sb.WriteString("\n")
	}

	return sb.String()
}
