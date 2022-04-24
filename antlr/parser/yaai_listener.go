// Code generated from /home/john/Dropbox/projects/yaai//antlr/Yaai.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Yaai
import "github.com/antlr/antlr4/runtime/Go/antlr"

// YaaiListener is a complete listener for a parse tree produced by Yaai.
type YaaiListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterPackageClause is called when entering the packageClause production.
	EnterPackageClause(c *PackageClauseContext)

	// EnterEos is called when entering the eos production.
	EnterEos(c *EosContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitPackageClause is called when exiting the packageClause production.
	ExitPackageClause(c *PackageClauseContext)

	// ExitEos is called when exiting the eos production.
	ExitEos(c *EosContext)
}
