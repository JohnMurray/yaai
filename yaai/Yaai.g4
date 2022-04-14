grammar Yaai;

// --------------------------------------------
// Tokens

// Keywords
KEYWORD_PACKAGE : 'package';

IDENTIFIER: [a-zA-Z][a-zA-Z0-9_]+;
L_BRACKET: '{';
R_BRACKET: '}';
EOL: [\n\r]+;

fragment Space: ' ';
fragment NewLine: '\n';
fragment CariageReturn: '\r';
fragment Tab: '\t';
WS: (Space | NewLine | CariageReturn | Tab)+? -> skip;

// --------------------------------------------
// Rules

// Start rule
unit :
   package_decl
   EOF;

package_decl : 'package ' IDENTIFIER EOL;

actor : 'actor ' IDENTIFIER L_BRACKET R_BRACKET;