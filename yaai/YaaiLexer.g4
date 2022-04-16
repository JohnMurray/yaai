lexer grammar YaaiLexer;

PACKAGE : 'package';
ACTOR   : 'actor';
TYPE    : 'type';
STRUCT  : 'struct';

T_INT    : 'int';
T_INT32  : 'int32';
T_INT64  : 'int64';
T_UINT   : 'uint';
T_UINT32 : 'uint32';
T_UINT64 : 'uint64';
T_STRING : 'string';

// SEMI: ';';
IDENTIFIER: [a-zA-Z][a-zA-Z0-9_]+;

L_BRACKET: '{';
R_BRACKET: '}';

// non-breaking whitespace
NB_WS: [ \t] -> skip;

EOS:  ([\r\n]+ | ';' | EOF) -> mode(DEFAULT_MODE);