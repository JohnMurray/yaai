parser grammar Yaai;

/// // --------------------------------------------
/// // Tokens
/// 
/// // Keywords
/// KEYWORD_PACKAGE : 'package';
/// 
/// IDENTIFIER: [a-zA-Z][a-zA-Z0-9_]+;
/// L_BRACKET: '{';
/// R_BRACKET: '}';
/// // EOL: [\n\r]+;
/// 
/// WS: [ \n\r\t]+ -> skip;
/// 
/// // --------------------------------------------
/// // Rules
/// 
/// // Start rule
/// file :
///    package_decl
///    top_level_expression*
///    EOF;
/// 
/// package_decl : 'package ' IDENTIFIER '\n';
/// 
/// top_level_expression
///    : actor
///    ;
/// 
/// actor : 'actor ' IDENTIFIER L_BRACKET R_BRACKET;

options {
	tokenVocab = YaaiLexer;
}

// OTHER: -> mode(DEFAULT_MODE), channel(HIDDEN);

packageClause: PACKAGE packageName = IDENTIFIER; // eos;

// eos:
// 	SEMI
// 	| EOF
// 	| EOS
// 	;