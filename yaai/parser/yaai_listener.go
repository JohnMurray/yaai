// Code generated from yaai/yaai.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // yaai

import "github.com/antlr/antlr4/runtime/Go/antlr"

// yaaiListener is a complete listener for a parse tree produced by yaaiParser.
type yaaiListener interface {
	antlr.ParseTreeListener

	// EnterPackage_decl is called when entering the package_decl production.
	EnterPackage_decl(c *Package_declContext)

	// EnterPackage_name is called when entering the package_name production.
	EnterPackage_name(c *Package_nameContext)

	// EnterPackage_name_start is called when entering the package_name_start production.
	EnterPackage_name_start(c *Package_name_startContext)

	// EnterPackage_name_end is called when entering the package_name_end production.
	EnterPackage_name_end(c *Package_name_endContext)

	// ExitPackage_decl is called when exiting the package_decl production.
	ExitPackage_decl(c *Package_declContext)

	// ExitPackage_name is called when exiting the package_name production.
	ExitPackage_name(c *Package_nameContext)

	// ExitPackage_name_start is called when exiting the package_name_start production.
	ExitPackage_name_start(c *Package_name_startContext)

	// ExitPackage_name_end is called when exiting the package_name_end production.
	ExitPackage_name_end(c *Package_name_endContext)
}
