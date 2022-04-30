parser grammar Yaai;


options {
	tokenVocab = YaaiLexer;
}


file
	: packageClause
	  EOF;

packageClause: PACKAGE packageName = IDENTIFIER eos;

varDecl       : VAR IDENTIFIER type_ (eos | varAssignment);
varAssignment : ASSIGNMENT expression eos;
varInit       : IDENTIFIER VAR_INITIALIZER expression eos;

type_
	: builtinType
	// TODO: add other types
	;

builtinType
	: T_INT
	| T_INT32
	| T_INT64
	| T_UINT
	| T_UINT32
	| T_UINT64
	| T_FLOAT32
	| T_FLOAT64
	| T_STRING
	;

expression
	: STRING_LITERAL
	| NUMERIC_LITERAL
	| FLOAT_LITERAL
	// TODO: this is an imcomplete list
	;

eos
	: EOS
	| EOF
	;