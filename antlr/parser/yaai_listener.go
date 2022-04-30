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

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterVarAssignment is called when entering the varAssignment production.
	EnterVarAssignment(c *VarAssignmentContext)

	// EnterVarInit is called when entering the varInit production.
	EnterVarInit(c *VarInitContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterBuiltinType is called when entering the builtinType production.
	EnterBuiltinType(c *BuiltinTypeContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEos is called when entering the eos production.
	EnterEos(c *EosContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitPackageClause is called when exiting the packageClause production.
	ExitPackageClause(c *PackageClauseContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitVarAssignment is called when exiting the varAssignment production.
	ExitVarAssignment(c *VarAssignmentContext)

	// ExitVarInit is called when exiting the varInit production.
	ExitVarInit(c *VarInitContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitBuiltinType is called when exiting the builtinType production.
	ExitBuiltinType(c *BuiltinTypeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEos is called when exiting the eos production.
	ExitEos(c *EosContext)
}
