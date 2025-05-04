// Code generated from Do.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Do
import "github.com/antlr4-go/antlr/v4"

// DoListener is a complete listener for a parse tree produced by DoParser.
type DoListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterGlobalVariableDefinition is called when entering the globalVariableDefinition production.
	EnterGlobalVariableDefinition(c *GlobalVariableDefinitionContext)

	// EnterFunctionDefinition is called when entering the functionDefinition production.
	EnterFunctionDefinition(c *FunctionDefinitionContext)

	// EnterTypeDefinition is called when entering the typeDefinition production.
	EnterTypeDefinition(c *TypeDefinitionContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterStructdefinition is called when entering the structdefinition production.
	EnterStructdefinition(c *StructdefinitionContext)

	// EnterBehavourdefinition is called when entering the behavourdefinition production.
	EnterBehavourdefinition(c *BehavourdefinitionContext)

	// EnterTypetuple is called when entering the typetuple production.
	EnterTypetuple(c *TypetupleContext)

	// EnterArglist is called when entering the arglist production.
	EnterArglist(c *ArglistContext)

	// EnterArgsublist is called when entering the argsublist production.
	EnterArgsublist(c *ArgsublistContext)

	// EnterBasetypefild is called when entering the basetypefild production.
	EnterBasetypefild(c *BasetypefildContext)

	// EnterVarfield is called when entering the varfield production.
	EnterVarfield(c *VarfieldContext)

	// EnterGlobalvarfield is called when entering the globalvarfield production.
	EnterGlobalvarfield(c *GlobalvarfieldContext)

	// EnterFieldname is called when entering the fieldname production.
	EnterFieldname(c *FieldnameContext)

	// EnterArgname is called when entering the argname production.
	EnterArgname(c *ArgnameContext)

	// EnterTypename is called when entering the typename production.
	EnterTypename(c *TypenameContext)

	// EnterGenericparamslist is called when entering the genericparamslist production.
	EnterGenericparamslist(c *GenericparamslistContext)

	// EnterGenericarglist is called when entering the genericarglist production.
	EnterGenericarglist(c *GenericarglistContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterExpressiontuple is called when entering the expressiontuple production.
	EnterExpressiontuple(c *ExpressiontupleContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFunctioncall is called when entering the functioncall production.
	EnterFunctioncall(c *FunctioncallContext)

	// EnterExpressiontuplelhv is called when entering the expressiontuplelhv production.
	EnterExpressiontuplelhv(c *ExpressiontuplelhvContext)

	// EnterExpressiontuplerhv is called when entering the expressiontuplerhv production.
	EnterExpressiontuplerhv(c *ExpressiontuplerhvContext)

	// EnterVariableuse is called when entering the variableuse production.
	EnterVariableuse(c *VariableuseContext)

	// EnterConstantuse is called when entering the constantuse production.
	EnterConstantuse(c *ConstantuseContext)

	// EnterEmptyexpression is called when entering the emptyexpression production.
	EnterEmptyexpression(c *EmptyexpressionContext)

	// EnterDividedname is called when entering the dividedname production.
	EnterDividedname(c *DividednameContext)

	// EnterBool is called when entering the bool production.
	EnterBool(c *BoolContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitGlobalVariableDefinition is called when exiting the globalVariableDefinition production.
	ExitGlobalVariableDefinition(c *GlobalVariableDefinitionContext)

	// ExitFunctionDefinition is called when exiting the functionDefinition production.
	ExitFunctionDefinition(c *FunctionDefinitionContext)

	// ExitTypeDefinition is called when exiting the typeDefinition production.
	ExitTypeDefinition(c *TypeDefinitionContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitStructdefinition is called when exiting the structdefinition production.
	ExitStructdefinition(c *StructdefinitionContext)

	// ExitBehavourdefinition is called when exiting the behavourdefinition production.
	ExitBehavourdefinition(c *BehavourdefinitionContext)

	// ExitTypetuple is called when exiting the typetuple production.
	ExitTypetuple(c *TypetupleContext)

	// ExitArglist is called when exiting the arglist production.
	ExitArglist(c *ArglistContext)

	// ExitArgsublist is called when exiting the argsublist production.
	ExitArgsublist(c *ArgsublistContext)

	// ExitBasetypefild is called when exiting the basetypefild production.
	ExitBasetypefild(c *BasetypefildContext)

	// ExitVarfield is called when exiting the varfield production.
	ExitVarfield(c *VarfieldContext)

	// ExitGlobalvarfield is called when exiting the globalvarfield production.
	ExitGlobalvarfield(c *GlobalvarfieldContext)

	// ExitFieldname is called when exiting the fieldname production.
	ExitFieldname(c *FieldnameContext)

	// ExitArgname is called when exiting the argname production.
	ExitArgname(c *ArgnameContext)

	// ExitTypename is called when exiting the typename production.
	ExitTypename(c *TypenameContext)

	// ExitGenericparamslist is called when exiting the genericparamslist production.
	ExitGenericparamslist(c *GenericparamslistContext)

	// ExitGenericarglist is called when exiting the genericarglist production.
	ExitGenericarglist(c *GenericarglistContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitExpressiontuple is called when exiting the expressiontuple production.
	ExitExpressiontuple(c *ExpressiontupleContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFunctioncall is called when exiting the functioncall production.
	ExitFunctioncall(c *FunctioncallContext)

	// ExitExpressiontuplelhv is called when exiting the expressiontuplelhv production.
	ExitExpressiontuplelhv(c *ExpressiontuplelhvContext)

	// ExitExpressiontuplerhv is called when exiting the expressiontuplerhv production.
	ExitExpressiontuplerhv(c *ExpressiontuplerhvContext)

	// ExitVariableuse is called when exiting the variableuse production.
	ExitVariableuse(c *VariableuseContext)

	// ExitConstantuse is called when exiting the constantuse production.
	ExitConstantuse(c *ConstantuseContext)

	// ExitEmptyexpression is called when exiting the emptyexpression production.
	ExitEmptyexpression(c *EmptyexpressionContext)

	// ExitDividedname is called when exiting the dividedname production.
	ExitDividedname(c *DividednameContext)

	// ExitBool is called when exiting the bool production.
	ExitBool(c *BoolContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)
}
