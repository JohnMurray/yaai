package parser

import (
	"strings"
)

// Represents the _entire_ repo/project that is being parsed/built
type Project struct {
	Packages []*Package
}

func (p *Project) Print(depth int) string {
	prefix := strings.Repeat("\t", depth)
	divider := strings.Repeat("-", 80)

	var sb strings.Builder

	for _, pkg := range p.Packages {
		sb.WriteString(prefix)
		sb.WriteString(divider)
		sb.WriteString("\n")
		sb.WriteString(pkg.Print(depth))
	}
	return sb.String()
}
