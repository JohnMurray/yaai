package parser

import (
	"fmt"
	"os"

	ap "github.com/JohnMurray/yaai/antlr/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func ParseFile(path string) Tree {
	contents, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Error opening file to parse: %v", err))
	}

	ypl := &yaaiParseListener{
		project: &Project{
			Packages: []*Package{},
		},
	}

	// setup the input
	is := antlr.NewInputStream(string(contents))

	// create the lexer
	lexer := ap.NewYaaiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// create the parser
	p := ap.NewYaai(stream)

	antlr.ParseTreeWalkerDefault.Walk(ypl, p.File())
	return ypl.project
}

type yaaiParseListener struct {
	*ap.BaseYaaiListener
	// *antlr.DefaultErrorListener

	project *Project

	activeFile    *File
	activePackage *Package
}

func (p *yaaiParseListener) EnterFile(ctx ap.FileContext) {
	// The active file should be set before parsing begins, so
	// let's just make sure it's there since we'll be referencing
	// this quite a bit
	if p.activeFile == nil {
		panic("current active file being parsed is unknown")
	}
}

func (p *yaaiParseListener) ExitFile(ctx ap.FileContext) {
	fmt.Println("exit file")
	// We're done parsing the file, reset the state on the listener.
	// These should be reset appropriately on the next pass
	p.activeFile = nil
	p.activePackage = nil

	// TOOD: make sure that there is nothing "open" still
}

func (p *yaaiParseListener) EnterPackageClause(ctx ap.PackageClauseContext) {
	fmt.Println("enter package")

	if p.activePackage != nil {
		panic(fmt.Sprintf("package name cannot be set, found '%s'", p.activePackage.Name))
	}

	name := ctx.GetPackageName().GetText()
	if len(name) == 0 {
		panic("package name cannot be an empty string, at least one character")
	}

	for _, pkg := range p.project.Packages {
		if pkg.Name == name {
			p.activePackage = pkg
			break
		}
	}
	if p.activePackage == nil {
		pkg := &Package{
			Name:  name,
			Files: []*File{},
		}
		p.activePackage = pkg
		p.project.Packages = append(p.project.Packages, pkg)
	}
	p.activePackage.Files = append(p.activePackage.Files, p.activeFile)
}
