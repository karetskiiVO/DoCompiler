grammar Do;

program: definition* EOF;

definition: functionDefinition | typeDefinition;

functionDefinition: 'do' NAME genericparamslist? '(' /* Здесь надо сделать лист аргументов */ ')' typetuple;
typeDefinition: 'with' NAME genericparamslist? type;

type:  ('*' type) | ('pipe' type) | typename | structdefinition | behavourdefinition; 

structdefinition: 'struct' '{' ((basetypefild | varfield | globalvarfield) ';')* '}';
behavourdefinition: 'behavour' '{' '}';

typetuple: typename? | (typename (',' typename)*);

basetypefild: typename genericarglist;
varfield: fieldname typename;
globalvarfield: fieldname 'glob' typename;
fieldname: NAME;
argumentname: NAME;
typename: DIVIDEDNAME genericparamslist?; // TODO: лямбды 
genericparamslist: '<' (NAME (',' NAME)*)? '>';
genericarglist: '<' (type (',' type)*)? '>';

BOOL: ('true' | 'false');
STRING: '"' .*? '"';
COMPARETOKEN: ('==' | '<=' | '>=' | '<' | '>' | '!=');
NUMBER: [-+]?[0-9]+;
NAME:   [a-zA-Z][a-zA-Z0-9]*;
DIVIDEDNAME:   NAME ('.' NAME)*?;
EMPTY:  [ \t\r\n]+ -> skip;

// decl typename<T1, T2, ...> struct {}
// decl typename<T1, T2, ...> anothertype<T1, T2, ...>
// decl typename<T1, T2, ...> behavour {}
// act f<T1, T2, ...> (argname typename, ...) returnType {}
// act f<T1, T2, ...> (self typename) (argname typename, ...) returnType {}
// act f<T1, T2, ...> [typename] (argname typename, ...) returnType {}
// with varname typename
// with varname = typename{}
// with varname = act () returnType {}
// event.Register(act () {})
// f()? -- проброс error            -- откалывает последный член tupple -- error всегда последнее значение
// f()! -- роняет программу с error -- откалывает последный член tupple -- error всегда последнее значение