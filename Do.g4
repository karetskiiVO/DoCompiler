grammar Do;

program: definition* EOF;

definition: functionDefinition | typeDefinition | globalVariableDefinition; // TODO: метод и глобальные переменные

globalVariableDefinition: 'var' NAME typename; /* expression */
functionDefinition: 'act' NAME genericparamslist? '(' arglist ')' typetuple statementblock;
typeDefinition: 'with' NAME genericparamslist? type;

type:  ('*' type) | ('pipe' type) | typename | structdefinition | behavourdefinition; 

structdefinition: 'struct' '{' ((basetypefild | varfield | globalvarfield) ';')* '}';
behavourdefinition: 'behavour' '{' '}';

statementblock: '{' statement* '}';
typetuple: type? | ('(' type (',' type)* ')');
arglist: argsublist? | (argsublist (',' argsublist));
argsublist: argname (',' argname)* type;

basetypefild: type genericarglist;
varfield: fieldname type;
globalvarfield: fieldname 'glob' type;
fieldname: NAME;
argname: NAME;
typename: dividedname genericparamslist?; // TODO: лямбды 
genericparamslist: '<' (NAME (',' NAME)*)? '>';
genericarglist: '<' (type (',' type)*)? '>'; // TODO: behavour

statement: assign | ifstatement | returnstatement;
assign: (expressiontuplelhv '=')? expressiontuplerhv;
ifstatement: 'if' expression statementblock elsestatement?;
elsestatement: 'else' (ifstatement | statementblock);
returnstatement: 'return' expressiontuple;

expressiontuple: expression (',' expression)*;
expression: emptyexpression | variableuse | ('(' expressiontuple ')') | constantuse | functioncall;
functioncall: dividedname '(' expressiontuple? ')';

expressiontuplelhv: expressiontuple;
expressiontuplerhv: expressiontuple;
variableuse: dividedname;
constantuse: bool | string | number;

emptyexpression: '_';
dividedname: NAME ('.' NAME)*;

bool: BOOL;
string: STRING;
number: NUMBER;


BOOL: ('true' | 'false');
STRING: '"' .*? '"';
COMPARETOKEN: ('==' | '<=' | '>=' | '<' | '>' | '!=');
NUMBER: [-+]?[0-9]+;
NAME:   [a-zA-Z][a-zA-Z0-9]*;
EMPTY:  [ \t\r\n]+ -> skip;

// decl typename<T1 behavour, T2, ...> struct {}
// decl typename<T1 behavour, T2, ...> anothertype<T1, T2, ...>
// decl typename<T1 behavour, T2, ...> behavour {}
// act f<T1 behavour, T2, ...> (argname typename, ...) returnType {}
// act f<T1 behavour, T2, ...> (self typename) (argname typename, ...) returnType {}
// act f<T1 behavour, T2, ...> [typename] (argname typename, ...) returnType {}
// with varname typename
// with varname = typename{}
// with varname = act () returnType {}
// event.Register(act () {})
// f()? -- проброс error            -- откалывает последный член tupple -- error всегда последнее значение
// f()! -- роняет программу с error -- откалывает последный член tupple -- error всегда последнее значение
// self type