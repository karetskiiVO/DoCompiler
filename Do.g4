grammar Do;

program: definition* EOF;

definition: functionDefinition | typeDefinition | globalVariableDefinition; // TODO: метод и глобальные переменные

globalVariableDefinition: 'var' NAME (',' NAME)* typename; /* expression */
functionDefinition: 'act' NAME '(' arglist ')' typetuple statementblock;
typeDefinition: 'with' NAME type;

type:  ('*' type) | ('pipe' type) | typename | structdefinition | behavourdefinition; 

structdefinition: 'struct' '{' ((basetypefild | varfield | globalvarfield) /* ';' */)* '}';
behavourdefinition: 'behavour' '{' '}';

statementblock: '{' statement* '}';
typetuple: type? | ('(' type (',' type)* ')');
arglist: argsublist? | (argsublist (',' argsublist));
argsublist: argname (',' argname)* type;

basetypefild: type;
varfield: fieldname type;
globalvarfield: fieldname 'glob' type;
fieldname: NAME;
argname: NAME;
typename: dividedname; // TODO: лямбды 
genericparamslist: '<' (NAME (',' NAME)*)? '>';
genericarglist: '<' (type (',' type)*)? '>'; // TODO: behavour

statement: assign | ifstatement | returnstatement | vardeclarationstatement;
vardeclarationstatement: 'var' NAME (',' NAME)* typename /* ';' */;
assign: (expressiontuplelhv ASSIGNTOKEN)? expressiontuplerhv ';';
ifstatement: 'if' ifexpression statementblock elsestatement?;
elsestatement: 'else'  statementblock;
returnstatement: 'return' expressiontuple ';';

ifexpression: expression;

expressiontuple: expression (',' expression)*;
expression:
    emptyexpression             |
    arithmetic                  |
    variableuse                 |
    constantuse                 |
    functioncall                |
    logic;
oneexpression:
    constantuse                 |
    variableuse                 |
    functioncall                |
    ('(' expression ')');
functioncall: dividedname '(' expressiontuple? ')';

arithmetic: multiply (SUMTOKEN multiply)*;
multiply: oneexpression (MULTTOKEN oneexpression)*;

logic: andlogic ('||' andlogic)*;
andlogic: compare ('&&' compare)*;
compare: (oneexpression | arithmetic) (COMPARETOKEN (oneexpression | arithmetic))?;

expressiontuplelhv: expressionlhv (',' expressionlhv)*;
expressionlhv: emptyexpression | variableuselhv;
expressiontuplerhv: expressiontuple;
variableuselhv: dividedname;
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
ASSIGNTOKEN: ('=' | ':=');
NUMBER: [-+]?[0-9]+;
NAME:   [a-zA-Z][a-zA-Z0-9]*;

COMMENT: '//' .*? '\n' -> skip;
MULTILINECOMMENT: '/*' .*? '*/' -> skip;
EMPTY:  [ \t\r\n]+ -> skip;
SUMTOKEN: ('+' | '-');
MULTTOKEN: ('^' | '/'); // Тут какая-то бага у antlr
