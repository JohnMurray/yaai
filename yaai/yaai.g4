grammar yaai;

// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
NUMBER: [0-9];
WHITESPACE: [ \r\n\t]+ -> skip;
SPACE: ' ';
LETTER : [a-z];
UNDERSCORE : '_';

// Rules
package_decl : 'package' SPACE package_name;
package_name : package_name_start package_name_end+ EOL;
package_name_start : LETTER;
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