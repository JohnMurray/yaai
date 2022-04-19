parser grammar Yaai;


options {
	tokenVocab = YaaiLexer;
}


file
	: packageClause
	  EOF;

packageClause: PACKAGE packageName = IDENTIFIER eos;

eos
	: EOS
	| EOF
	;