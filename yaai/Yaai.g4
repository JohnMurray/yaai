grammar Yaai;

// --------------------------------------------
// Tokens

// Keywords
KEYWORD_PACKAGE : 'package';

IDENTIFIER: [a-zA-Z][a-zA-Z0-9_]+;
L_BRACKET: '{';
R_BRACKET: '}';
EOL: [\n\r]+;

WS: [ \n\r\t]+ -> skip;

// --------------------------------------------
// Rules

// Start rule
file :
   package_decl
   EOF;

package_decl : 'package ' IDENTIFIER EOL;

actor : 'actor ' IDENTIFIER L_BRACKET R_BRACKET;