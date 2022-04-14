// Code generated from yaai/Yaai.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Yaai

import "github.com/antlr/antlr4/runtime/Go/antlr"

// YaaiListener is a complete listener for a parse tree produced by YaaiParser.
type YaaiListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterPackage_decl is called when entering the package_decl production.
	EnterPackage_decl(c *Package_declContext)

	// EnterActor is called when entering the actor production.
	EnterActor(c *ActorContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitPackage_decl is called when exiting the package_decl production.
	ExitPackage_decl(c *Package_declContext)

	// ExitActor is called when exiting the actor production.
	ExitActor(c *ActorContext)
}
