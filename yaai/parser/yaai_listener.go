// Code generated from yaai/Yaai.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Yaai
import "github.com/antlr/antlr4/runtime/Go/antlr"

// YaaiListener is a complete listener for a parse tree produced by YaaiParser.
type YaaiListener interface {
	antlr.ParseTreeListener

	// EnterPackageClause is called when entering the packageClause production.
	EnterPackageClause(c *PackageClauseContext)

	// ExitPackageClause is called when exiting the packageClause production.
	ExitPackageClause(c *PackageClauseContext)
}
