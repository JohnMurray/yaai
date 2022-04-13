grammar Yaai;

// Tokens
NUMBER: [0-9];
SPACE: ' ';
LETTER : [a-zA-Z];
UNDERSCORE : '_';
EOL: [\n\r]+;
WHITESPACE: [ \r\n\t]+ -> skip;

KEYWORD_PACKAGE : 'package';

// Rules
package_decl : KEYWORD_PACKAGE SPACE package_name EOL;
package_name : LETTER package_name_end+;
package_name_end
   : LETTER
   | NUMBER
   | UNDERSCORE
   ;

// start : expression EOF;
// 
// expression
//    : expression op=('*'|'/') expression # MulDiv
//    | expression op=('+'|'-') expression # AddSub
//    | NUMBER                             # Number
//    ;