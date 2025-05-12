package doListeners

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/karetskiiVO/DoCompiler/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	"github.com/karetskiiVO/DoCompiler/compiler"
)

// DoSourceListener implements the parser.DoListener interface
type DoSourceListener struct {
	*parser.BaseDoListener

	currfunc *ir.Func
	program  *compiler.Program

	blocks []*ir.Block
	values [][]value.Value
}

// NewDoSourceListener creates a new listener for the Do language
func NewDoSourceListener(program *compiler.Program) antlr.ParseTreeListener {
	if program.Error() != nil {
		return new(parser.BaseDoListener)
	}

	return &DoSourceListener{
		program: program,
	}
}

// Block management helpers
func (l *DoSourceListener) pushBlock(block *ir.Block) { l.blocks = append(l.blocks, block) }
func (l *DoSourceListener) popBlock() *ir.Block {
	lastBlock := l.blocks[len(l.blocks)-1]
	l.blocks = l.blocks[:len(l.blocks)-1]
	return lastBlock
}
func (l *DoSourceListener) topBlock() *ir.Block { return l.blocks[len(l.blocks)-1] }
func (l *DoSourceListener) newBlock() *ir.Block { return l.currfunc.NewBlock("") }

// Value management helpers
func (l *DoSourceListener) addValue(val ...value.Value) {
	l.values[len(l.values)-1] = append(l.values[len(l.values)-1], val...)
}

func (l *DoSourceListener) currentValues() []value.Value {
	return l.values[len(l.values)-1]
}

func (l *DoSourceListener) popValues() []value.Value {
	vals := l.values[len(l.values)-1]
	l.values = l.values[:len(l.values)-1]
	return vals
}

// Error reporting
func (l *DoSourceListener) reportError(token antlr.Token, err error) {
	l.reportErrorAt(token.GetLine(), token.GetColumn(), token.GetInputStream().GetSourceName(), err)
}

func (l *DoSourceListener) reportErrorAt(line, column int, source string, err error) {
	l.program.AddErrorf("%v:%v:%v: %v", source, line, column, err)
}

// Utility methods for common operations
func (l *DoSourceListener) getVariableAndType(varname string) (value.Value, types.Type, error) {
	variable, err := l.program.GetVariable(varname)
	if err != nil {
		return nil, nil, err
	}
	
	varType := variable.Type().(*types.PointerType).ElemType
	return variable, varType, nil
}

func (l *DoSourceListener) getStructField(structVar value.Value, fieldname string) (value.Value, types.Type, error) {
	structType := structVar.Type().(*types.PointerType).ElemType
	fieldidx, err := l.program.GetFieldIdx(structType, fieldname)
	if err != nil {
		return nil, nil, err
	}
	
	fieldType := structType.(*types.StructType).Fields[fieldidx]
	fieldptr := l.topBlock().NewGetElementPtr(
		fieldType,
		structVar,
		constant.NewIndex(constant.NewInt(types.I64, int64(fieldidx))),
	)
	
	return fieldptr, fieldType, nil
}

// Function Definition
func (l *DoSourceListener) EnterFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	funcname := ctx.NAME().GetText()
	function, _ := l.program.GetFunction(funcname)
	
	l.program.NewScope()
	l.currfunc = function
	
	// Create entry block
	entry := l.newBlock()
	l.pushBlock(entry)
	
	// Process function arguments
	l.processFunctionArguments(ctx)
}

func (l *DoSourceListener) processFunctionArguments(ctx *parser.FunctionDefinitionContext) {
	paramIndex := 0
	
	for _, arglistToken := range ctx.Arglist().AllArgsublist() {
		argtype := arglistToken.Type_().GetText()
		
		for _, argnameToken := range arglistToken.AllArgname() {
			paramIndex++
			argname := argnameToken.GetText()

			variable, err := l.program.RegisterVariable(l.topBlock(), argname, argtype)
			if err != nil {
				l.reportError(argnameToken.GetStart(), err)
				continue
			}

			l.topBlock().NewStore(l.currfunc.Params[paramIndex-1], variable)
		}
	}
}

func (l *DoSourceListener) ExitFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	rettype := l.currfunc.Sig.RetType.(*types.StructType)
	
	// Add a return statement for functions without explicit returns
	if len(rettype.Fields) == 0 {
		retvalRef := l.topBlock().NewAlloca(l.currfunc.Sig.RetType)
		retval := l.topBlock().NewLoad(l.currfunc.Sig.RetType, retvalRef)
		retval.SetName("_ret")
		l.topBlock().NewRet(retval)
	} else {
		l.topBlock().NewUnreachable()
	}

	l.popBlock()
	l.program.TerminateScope()
}

// Function Calls
func (l *DoSourceListener) EnterFunctioncall(ctx *parser.FunctioncallContext) {
	l.values = append(l.values, []value.Value{})
}

func (l *DoSourceListener) ExitFunctioncall(ctx *parser.FunctioncallContext) {
	funcname := ctx.Dividedname().GetText()
	function, err := l.program.GetFunction(funcname)
	if err != nil {
		l.reportError(ctx.Dividedname().GetStart(), err)
		return
	}
	
	args := l.popValues()
	
	// Validate argument count
	if len(args) != len(function.Sig.Params) {
		l.reportError(ctx.Dividedname().GetStart(), 
			fmt.Errorf("mistmatch in argument values: expected: %v actual: %v", 
				len(function.Sig.Params), len(args)))
		return
	}
	
	// Generate function call
	call := l.topBlock().NewCall(function, args...)

	// Extract return values
	returnType := function.Sig.RetType.(*types.StructType)
	for i := range returnType.Fields {
		l.addValue(l.topBlock().NewExtractValue(call, uint64(i)))
	}
}

// Constants
func (l *DoSourceListener) ExitConstantuse(ctx *parser.ConstantuseContext) {
	if ctx.Bool_() != nil {
		l.processBoolConstant(ctx)
	} else if ctx.Number() != nil {
		l.processNumberConstant(ctx)
	} else {
		panic(fmt.Errorf("unsupported constant type: %v", ctx.GetText()))
	}
}

func (l *DoSourceListener) processBoolConstant(ctx *parser.ConstantuseContext) {
	typ, _ := l.program.GetType("bool")
	val := int64(0)
	if ctx.Bool_().GetText() == "true" {
		val = 1
	}

	l.addValue(constant.NewInt(typ.(*types.IntType), val))
}

func (l *DoSourceListener) processNumberConstant(ctx *parser.ConstantuseContext) {
	val, err := strconv.ParseInt(ctx.Number().GetText(), 0, 64)
	if err != nil {
		l.reportError(ctx.GetStart(), err)
		return
	}

	typ, _ := l.program.GetType("int")
	l.addValue(constant.NewInt(typ.(*types.IntType), int64(val)))
}

// Variable Use
func (l *DoSourceListener) ExitVariableuse(ctx *parser.VariableuseContext) {
	varname := ctx.Dividedname().GetText()

	// Try to process as a simple variable
	if l.processSimpleVariable(varname) {
		return
	}

	// Try to process as a struct field
	if l.processStructField(varname, ctx) {
		return
	}

	// If we get here, the variable couldn't be found
	l.reportError(ctx.GetStart(), fmt.Errorf("unknown variable '%s'", varname))
}

func (l *DoSourceListener) processSimpleVariable(varname string) bool {
	variable, err := l.program.GetVariable(varname)
	if err != nil {
		return false
	}

	varType := variable.Type().(*types.PointerType).ElemType
	l.addValue(l.topBlock().NewLoad(varType, variable))
	return true
}

func (l *DoSourceListener) processStructField(varname string, ctx *parser.VariableuseContext) bool {
	idx := strings.LastIndexByte(varname, '.')
	if idx == -1 {
		return false
	}

	// Split into variable name and field name
	fieldname := varname[idx+1:]
	structVarname := varname[:idx]

	// Get the variable representing the struct
	structVar, err := l.program.GetVariable(structVarname)
	if err != nil {
		l.reportError(ctx.GetStart(), err)
		return true // We handled this case with an error
	}

	// Get the field's index in the struct
	structType := structVar.Type().(*types.PointerType).ElemType
	fieldidx, err := l.program.GetFieldIdx(structType, fieldname)
	if err != nil {
		l.reportError(ctx.GetStart(), err)
		return true
	}

	// Create a pointer to the struct field and load its value
	fieldType := structType.(*types.StructType).Fields[fieldidx]
	fieldptr := l.topBlock().NewGetElementPtr(
		fieldType,
		structVar,
		constant.NewIndex(constant.NewInt(types.I64, int64(fieldidx))),
	)

	l.addValue(l.topBlock().NewLoad(fieldptr.ElemType, fieldptr))
	return true
}

// Assignment
func (l *DoSourceListener) ExitAssign(ctx *parser.AssignContext) {
	if ctx.Expressiontuplelhv() == nil {
		return
	}

	lhvExpressions := ctx.Expressiontuplelhv().Expressiontuple().AllExpression()
	lhvLen := len(lhvExpressions)
	values := l.currentValues()
	emptyExpr := l.countEmptyExpressions(lhvExpressions)

	// Validate expressions are mutable
	l.validateMutableExpressions(lhvExpressions)

	// Adjust values based on empty expressions
	values = values[lhvLen-emptyExpr:]

	// Validate correct number of expressions
	if lhvLen != len(values) {
		stop := ctx.Expressiontuplelhv().GetStop()
		l.reportError(stop, fmt.Errorf(
			"incorrect number of expressions in the assignment expected: %v actual: %v",
			lhvLen, len(values)))
		return
	}

	// Process assignments
	l.processAssignments(lhvExpressions, values)
}

func (l *DoSourceListener) countEmptyExpressions(expressions []parser.IExpressionContext) int {
	count := 0
	for _, expr := range expressions {
		if expr.Emptyexpression() != nil {
			count++
		}
	}
	return count
}

func (l *DoSourceListener) validateMutableExpressions(expressions []parser.IExpressionContext) {
	for _, expr := range expressions {
		if expr.Variableuse() == nil && expr.Emptyexpression() == nil {
			loc := expr.GetStart()
			l.reportError(loc, fmt.Errorf("expression `%v` is not mutable", expr.GetText()))
		}
	}
}

func (l *DoSourceListener) processAssignments(expressions []parser.IExpressionContext, values []value.Value) {
	for i, expr := range expressions {
		if expr.Variableuse() != nil {
			l.processVariableAssignment(expr, values[i])
		}
	}
}

func (l *DoSourceListener) processVariableAssignment(expr parser.IExpressionContext, value value.Value) {
	varname := expr.GetText()

	// Try simple variable assignment first
	if err := l.assignToVariable(varname, value, expr); err == nil {
		return
	}

	// Try struct field assignment
	idx := strings.LastIndexByte(varname, '.')
	if idx != -1 {
		fieldname := varname[idx+1:]
		varname = varname[:idx]
		
		if err := l.assignToStructField(varname, fieldname, value, expr); err != nil {
			l.reportError(expr.GetStart(), err)
		}
	} else {
		// If we reach here, we couldn't assign to a variable or struct field
		l.reportError(expr.GetStart(), fmt.Errorf("unknown variable '%s'", varname))
	}
}

func (l *DoSourceListener) assignToVariable(varname string, value value.Value, expr parser.IExpressionContext) error {
	variable, err := l.program.GetVariable(varname)
	if err != nil {
		return err
	}

	varType := variable.Type().(*types.PointerType).ElemType
	valType := value.Type()

	if varType != valType {
		l.reportError(expr.GetStart(), fmt.Errorf(
			"type mismatch in assignment: variable '%v' has type %v, but assigned value has type %v",
			varname, l.program.Typename(varType), l.program.Typename(valType)))
		return nil
	}

	l.topBlock().NewStore(value, variable)
	return nil
}

func (l *DoSourceListener) assignToStructField(varname, fieldname string, value value.Value, expr parser.IExpressionContext) error {
	variable, err := l.program.GetVariable(varname)
	if err != nil {
		return err
	}

	fieldidx, err := l.program.GetFieldIdx(variable.Type().(*types.PointerType).ElemType, fieldname)
	if err != nil {
		return err
	}

	structType := variable.Type().(*types.PointerType).ElemType.(*types.StructType)
	fieldType := structType.Fields[fieldidx]
	valType := value.Type()

	if fieldType != valType {
		l.reportError(expr.GetStart(), fmt.Errorf(
			"type mismatch in assignment: variable '%v' has type %v, but assigned value has type %v",
			varname+"."+fieldname, l.program.Typename(fieldType), l.program.Typename(valType)))
		return nil
	}

	fieldptr := l.topBlock().NewGetElementPtr(
		fieldType,
		variable,
		constant.NewIndex(constant.NewInt(types.I64, int64(fieldidx))),
	)

	l.topBlock().NewStore(value, fieldptr)
	return nil
}

// Statement handling
func (l *DoSourceListener) EnterStatement(ctx *parser.StatementContext) {
	l.values = append(l.values, []value.Value{})
}

func (l *DoSourceListener) ExitStatement(ctx *parser.StatementContext) {
	l.values = l.values[:len(l.values)-1]
}

// If statements
func (l *DoSourceListener) EnterIfstatement(ctx *parser.IfstatementContext) {
	trueBlock := l.newBlock()
	falseBlock := l.newBlock()

	l.pushBlock(falseBlock)
	l.pushBlock(trueBlock)
}

func (l *DoSourceListener) ExitIfstatement(ctx *parser.IfstatementContext) {
	baseBlock := l.blocks[len(l.blocks)-3]
	falseBlock := l.blocks[len(l.blocks)-2]
	trueBlock := l.blocks[len(l.blocks)-1]

	values := l.currentValues()

	if len(values) != 1 {
		l.reportError(ctx.Expression().GetStart(), 
			fmt.Errorf("if statement expected 1 bool, actual: %v", len(values)))
		return
	}

	// Pop blocks from stack
	l.popBlock() // trueBlock
	l.popBlock() // falseBlock
	l.popBlock() // baseBlock

	// Create result block
	resBlock := falseBlock
	if ctx.Elsestatement() != nil {
		resBlock = l.newBlock()
		falseBlock.NewBr(resBlock)
	}

	// Create conditional branch
	baseBlock.NewCondBr(values[0], trueBlock, falseBlock)
	trueBlock.NewBr(resBlock)

	l.pushBlock(resBlock)
}

func (l *DoSourceListener) EnterElsestatement(ctx *parser.ElsestatementContext) {
	idx1, idx2 := len(l.blocks)-1, len(l.blocks)-2
	l.blocks[idx1], l.blocks[idx2] = l.blocks[idx2], l.blocks[idx1]
	l.values = append(l.values, []value.Value{})
}

func (l *DoSourceListener) ExitElsestatement(ctx *parser.ElsestatementContext) {
	idx1, idx2 := len(l.blocks)-1, len(l.blocks)-2
	l.blocks[idx1], l.blocks[idx2] = l.blocks[idx2], l.blocks[idx1]
	l.values = l.values[:len(l.values)-1]
}

// Return statements
func (l *DoSourceListener) ExitReturnstatement(ctx *parser.ReturnstatementContext) {
	values := l.currentValues()
	rettype := l.currfunc.Sig.RetType.(*types.StructType)

	if len(rettype.Fields) != len(values) {
		l.reportError(ctx.GetStart(), fmt.Errorf(
			"incorrect number of expressions in the return expected: %v actual: %v", 
			len(rettype.Fields), len(values)))
		return
	}

	l.generateReturnValue(rettype, values)
}

func (l *DoSourceListener) generateReturnValue(rettype *types.StructType, values []value.Value) {
	retvalRef := l.topBlock().NewAlloca(l.currfunc.Sig.RetType)
	
	for i, field := range rettype.Fields {
		fieldPtr := l.topBlock().NewGetElementPtr(
			field, 
			retvalRef, 
			constant.NewIndex(constant.NewInt(types.I64, int64(i))),
		)
		l.topBlock().NewStore(values[i], fieldPtr)
	}
	
	retval := l.topBlock().NewLoad(l.currfunc.Sig.RetType, retvalRef)
	retval.SetName("_ret")

	l.topBlock().NewRet(retval)
	l.popBlock()
	l.pushBlock(l.newBlock())
}

// Variable declaration
func (l *DoSourceListener) ExitVardeclarationstatement(ctx *parser.VardeclarationstatementContext) {
	vartype := ctx.Typename().GetText()

	for _, varnameToken := range ctx.AllNAME() {
		varname := varnameToken.GetText()

		variable, err := l.program.RegisterVariable(l.topBlock(), varname, vartype)
		if err != nil {
			l.reportError(varnameToken.GetSymbol(), err)
			continue
		}

		// Initialize with zero value
		varType := variable.Type().(*types.PointerType).ElemType
		l.topBlock().NewStore(constant.NewZeroInitializer(varType), variable)
	}
}

// Statement blocks (scopes)
func (l *DoSourceListener) EnterStatementblock(ctx *parser.StatementblockContext) {
	l.program.NewScope()
}

func (l *DoSourceListener) ExitStatementblock(ctx *parser.StatementblockContext) {
	l.program.TerminateScope()
}
