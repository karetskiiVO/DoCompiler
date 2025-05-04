// Code generated from Do.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Do
import "github.com/antlr4-go/antlr/v4"

// BaseDoListener is a complete listener for a parse tree produced by DoParser.
type BaseDoListener struct{}

var _ DoListener = &BaseDoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseDoListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseDoListener) ExitProgram(ctx *ProgramContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseDoListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseDoListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterGlobalVariableDefinition is called when production globalVariableDefinition is entered.
func (s *BaseDoListener) EnterGlobalVariableDefinition(ctx *GlobalVariableDefinitionContext) {}

// ExitGlobalVariableDefinition is called when production globalVariableDefinition is exited.
func (s *BaseDoListener) ExitGlobalVariableDefinition(ctx *GlobalVariableDefinitionContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *BaseDoListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *BaseDoListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterTypeDefinition is called when production typeDefinition is entered.
func (s *BaseDoListener) EnterTypeDefinition(ctx *TypeDefinitionContext) {}

// ExitTypeDefinition is called when production typeDefinition is exited.
func (s *BaseDoListener) ExitTypeDefinition(ctx *TypeDefinitionContext) {}

// EnterType is called when production type is entered.
func (s *BaseDoListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseDoListener) ExitType(ctx *TypeContext) {}

// EnterStructdefinition is called when production structdefinition is entered.
func (s *BaseDoListener) EnterStructdefinition(ctx *StructdefinitionContext) {}

// ExitStructdefinition is called when production structdefinition is exited.
func (s *BaseDoListener) ExitStructdefinition(ctx *StructdefinitionContext) {}

// EnterBehavourdefinition is called when production behavourdefinition is entered.
func (s *BaseDoListener) EnterBehavourdefinition(ctx *BehavourdefinitionContext) {}

// ExitBehavourdefinition is called when production behavourdefinition is exited.
func (s *BaseDoListener) ExitBehavourdefinition(ctx *BehavourdefinitionContext) {}

// EnterTypetuple is called when production typetuple is entered.
func (s *BaseDoListener) EnterTypetuple(ctx *TypetupleContext) {}

// ExitTypetuple is called when production typetuple is exited.
func (s *BaseDoListener) ExitTypetuple(ctx *TypetupleContext) {}

// EnterArglist is called when production arglist is entered.
func (s *BaseDoListener) EnterArglist(ctx *ArglistContext) {}

// ExitArglist is called when production arglist is exited.
func (s *BaseDoListener) ExitArglist(ctx *ArglistContext) {}

// EnterArgsublist is called when production argsublist is entered.
func (s *BaseDoListener) EnterArgsublist(ctx *ArgsublistContext) {}

// ExitArgsublist is called when production argsublist is exited.
func (s *BaseDoListener) ExitArgsublist(ctx *ArgsublistContext) {}

// EnterBasetypefild is called when production basetypefild is entered.
func (s *BaseDoListener) EnterBasetypefild(ctx *BasetypefildContext) {}

// ExitBasetypefild is called when production basetypefild is exited.
func (s *BaseDoListener) ExitBasetypefild(ctx *BasetypefildContext) {}

// EnterVarfield is called when production varfield is entered.
func (s *BaseDoListener) EnterVarfield(ctx *VarfieldContext) {}

// ExitVarfield is called when production varfield is exited.
func (s *BaseDoListener) ExitVarfield(ctx *VarfieldContext) {}

// EnterGlobalvarfield is called when production globalvarfield is entered.
func (s *BaseDoListener) EnterGlobalvarfield(ctx *GlobalvarfieldContext) {}

// ExitGlobalvarfield is called when production globalvarfield is exited.
func (s *BaseDoListener) ExitGlobalvarfield(ctx *GlobalvarfieldContext) {}

// EnterFieldname is called when production fieldname is entered.
func (s *BaseDoListener) EnterFieldname(ctx *FieldnameContext) {}

// ExitFieldname is called when production fieldname is exited.
func (s *BaseDoListener) ExitFieldname(ctx *FieldnameContext) {}

// EnterArgname is called when production argname is entered.
func (s *BaseDoListener) EnterArgname(ctx *ArgnameContext) {}

// ExitArgname is called when production argname is exited.
func (s *BaseDoListener) ExitArgname(ctx *ArgnameContext) {}

// EnterTypename is called when production typename is entered.
func (s *BaseDoListener) EnterTypename(ctx *TypenameContext) {}

// ExitTypename is called when production typename is exited.
func (s *BaseDoListener) ExitTypename(ctx *TypenameContext) {}

// EnterGenericparamslist is called when production genericparamslist is entered.
func (s *BaseDoListener) EnterGenericparamslist(ctx *GenericparamslistContext) {}

// ExitGenericparamslist is called when production genericparamslist is exited.
func (s *BaseDoListener) ExitGenericparamslist(ctx *GenericparamslistContext) {}

// EnterGenericarglist is called when production genericarglist is entered.
func (s *BaseDoListener) EnterGenericarglist(ctx *GenericarglistContext) {}

// ExitGenericarglist is called when production genericarglist is exited.
func (s *BaseDoListener) ExitGenericarglist(ctx *GenericarglistContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseDoListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseDoListener) ExitStatement(ctx *StatementContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseDoListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseDoListener) ExitAssign(ctx *AssignContext) {}

// EnterExpressiontuple is called when production expressiontuple is entered.
func (s *BaseDoListener) EnterExpressiontuple(ctx *ExpressiontupleContext) {}

// ExitExpressiontuple is called when production expressiontuple is exited.
func (s *BaseDoListener) ExitExpressiontuple(ctx *ExpressiontupleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDoListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDoListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFunctioncall is called when production functioncall is entered.
func (s *BaseDoListener) EnterFunctioncall(ctx *FunctioncallContext) {}

// ExitFunctioncall is called when production functioncall is exited.
func (s *BaseDoListener) ExitFunctioncall(ctx *FunctioncallContext) {}

// EnterExpressiontuplelhv is called when production expressiontuplelhv is entered.
func (s *BaseDoListener) EnterExpressiontuplelhv(ctx *ExpressiontuplelhvContext) {}

// ExitExpressiontuplelhv is called when production expressiontuplelhv is exited.
func (s *BaseDoListener) ExitExpressiontuplelhv(ctx *ExpressiontuplelhvContext) {}

// EnterExpressiontuplerhv is called when production expressiontuplerhv is entered.
func (s *BaseDoListener) EnterExpressiontuplerhv(ctx *ExpressiontuplerhvContext) {}

// ExitExpressiontuplerhv is called when production expressiontuplerhv is exited.
func (s *BaseDoListener) ExitExpressiontuplerhv(ctx *ExpressiontuplerhvContext) {}

// EnterVariableuse is called when production variableuse is entered.
func (s *BaseDoListener) EnterVariableuse(ctx *VariableuseContext) {}

// ExitVariableuse is called when production variableuse is exited.
func (s *BaseDoListener) ExitVariableuse(ctx *VariableuseContext) {}

// EnterConstantuse is called when production constantuse is entered.
func (s *BaseDoListener) EnterConstantuse(ctx *ConstantuseContext) {}

// ExitConstantuse is called when production constantuse is exited.
func (s *BaseDoListener) ExitConstantuse(ctx *ConstantuseContext) {}

// EnterEmptyexpression is called when production emptyexpression is entered.
func (s *BaseDoListener) EnterEmptyexpression(ctx *EmptyexpressionContext) {}

// ExitEmptyexpression is called when production emptyexpression is exited.
func (s *BaseDoListener) ExitEmptyexpression(ctx *EmptyexpressionContext) {}

// EnterDividedname is called when production dividedname is entered.
func (s *BaseDoListener) EnterDividedname(ctx *DividednameContext) {}

// ExitDividedname is called when production dividedname is exited.
func (s *BaseDoListener) ExitDividedname(ctx *DividednameContext) {}

// EnterBool is called when production bool is entered.
func (s *BaseDoListener) EnterBool(ctx *BoolContext) {}

// ExitBool is called when production bool is exited.
func (s *BaseDoListener) ExitBool(ctx *BoolContext) {}

// EnterString is called when production string is entered.
func (s *BaseDoListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseDoListener) ExitString(ctx *StringContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseDoListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseDoListener) ExitNumber(ctx *NumberContext) {}
