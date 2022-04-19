lexer grammar YaaiLexer;

ACTOR   : 'actor';
PACKAGE : 'package';
STRUCT  : 'struct';
RECEIVE : 'receive';
TYPE    : 'type';
VAR     : 'var';

T_INT    : 'int';
T_INT32  : 'int32';
T_INT64  : 'int64';
T_UINT   : 'uint';
T_UINT32 : 'uint32';
T_UINT64 : 'uint64';
T_STRING : 'string';

IDENTIFIER: [a-zA-Z][a-zA-Z0-9_]*;

VAR_INITIALIZER : ':=';

L_BRACKET    : '{';
R_BRACKET    : '}';
L_PAREN      : '(';
R_PAREN      : ')';
// SEMI: ';';

NUMERIC_LITERAL : [0-9]+;

// non-breaking whitespace
NB_WS: [ \t] -> skip;

LINE_COMMENT           : '//' ~[\r\n]*      -> channel(HIDDEN);

EOS:  ([\r\n]+ | ';' | EOF) -> mode(DEFAULT_MODE);