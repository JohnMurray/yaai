grammar Yaai;

// --------------------------------------------
// Tokens

// Keywords
KEYWORD_PACKAGE : 'package';

IDENTIFIER: [a-zA-Z][a-zA-Z0-9_]+;
L_BRACKET: '{';
R_BRACKET: '}';
SPACE: ' ';
EOL: [\n\r]+;

// NUMBER: [0-9];
// LETTER : [a-zA-Z];
// UNDERSCORE : '_';

WHITESPACE: [ \r\n\t]+ -> skip;

// --------------------------------------------
// Rules

// Start rule
unit :
   package_decl
   EOF;

package_decl : KEYWORD_PACKAGE SPACE IDENTIFIER EOL;

actor : keyword='actor' SPACE IDENTIFIER L_BRACKET R_BRACKET;