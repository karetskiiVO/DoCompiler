package doListeners

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/karetskiiVO/DoCompiler/parser"
	"github.com/karetskiiVO/slices"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	"github.com/karetskiiVO/DoCompiler/compiler"
)

type DoSourceListener struct {
	*parser.BaseDoListener

	currfunc *ir.Func
	program  *compiler.Program

	blocks []*ir.Block
	values [][]value.Value

	oneExpressionContext []int
}

func NewDoSourceListener(program *compiler.Program) antlr.ParseTreeListener {
	if program.Error() != nil {
		return new(parser.BaseDoListener)
	}

	return &DoSourceListener{
		program: program,
	}
}

func (l *DoSourceListener) pushBlock(block *ir.Block) { l.blocks = append(l.blocks, block) }
func (l *DoSourceListener) popBlock() *ir.Block {
	lastBlock := l.blocks[len(l.blocks)-1]
	l.blocks = l.blocks[:len(l.blocks)-1]
	return lastBlock
}
func (l *DoSourceListener) topBlock() *ir.Block { return l.blocks[len(l.blocks)-1] }
func (l *DoSourceListener) newBlock() *ir.Block { return l.currfunc.NewBlock("") }

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

func (l *DoSourceListener) reportError(token antlr.Token, err error) {
	l.reportErrorAt(token.GetLine(), token.GetColumn(), token.GetInputStream().GetSourceName(), err)
}

func (l *DoSourceListener) reportErrorf(token antlr.Token, format string, args ...interface{}) {
	l.reportErrorAt(
		token.GetLine(),
		token.GetColumn(),
		token.GetInputStream().GetSourceName(),
		fmt.Errorf(format, args...),
	)
}

func (l *DoSourceListener) reportErrorAt(line, column int, source string, err error) {
	l.program.AddErrorf("%v:%v:%v: %v", source, line, column, err)
}

func (l *DoSourceListener) EnterFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	funcname := ctx.NAME().GetText()
	function, _ := l.program.GetFunction(funcname)

	l.program.NewScope()
	l.currfunc = function

	entry := l.newBlock()
	l.pushBlock(entry)

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

	if len(args) != len(function.Sig.Params) {
		l.reportErrorf(ctx.Dividedname().GetStart(),
			"mistmatch in argument values: expected: %v actual: %v",
			len(function.Sig.Params),
			len(args),
		)
		return
	}

	retAlloca := l.topBlock().NewAlloca(function.Sig.RetType)
	l.topBlock().NewStore(l.topBlock().NewCall(function, args...), retAlloca)

	returnType := function.Sig.RetType.(*types.StructType)
	for i, fieldType := range returnType.Fields {
		l.addValue(
			l.topBlock().NewLoad(
				fieldType,
				l.topBlock().NewGetElementPtr(
					returnType,
					retAlloca,
					constant.NewIndex(constant.NewInt(types.I32, int64(0))),
					constant.NewIndex(constant.NewInt(types.I32, int64(i))),
				),
			),
		)
	}
}

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

func (l *DoSourceListener) ExitVariableuse(ctx *parser.VariableuseContext) {
	varname := ctx.Dividedname().GetText()

	if l.processSimpleVariable(varname) {
		return
	}

	if l.processStructField(varname, ctx) {
		return
	}

	l.reportErrorf(ctx.GetStart(), "variable `%v` is not declared in this scope", varname)
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

	fieldname := varname[idx+1:]
	structVarname := varname[:idx]
	structVar, err := l.program.GetVariable(structVarname)
	if err != nil {
		l.reportError(ctx.GetStart(), err)
		return true
	}

	structType := structVar.Type().(*types.PointerType).ElemType
	fieldidx, err := l.program.GetFieldIdx(structType, fieldname)
	if err != nil {
		l.reportError(ctx.GetStart(), err)
		return true
	}

	fieldptr := l.topBlock().NewGetElementPtr(
		structType,
		structVar,
		constant.NewInt(types.I32, 0),
		constant.NewInt(types.I32, int64(fieldidx)),
	)

	l.addValue(l.topBlock().NewLoad(fieldptr.ElemType, fieldptr))
	return true
}

func (l *DoSourceListener) ExitAssign(ctx *parser.AssignContext) {
	if ctx.Expressiontuplelhv() == nil {
		return
	}

	inlinedecl := (ctx.ASSIGNTOKEN().GetText() == ":=")

	lhvExpressions := ctx.Expressiontuplelhv().AllExpressionlhv()
	lhvLen := len(lhvExpressions)
	values := l.currentValues()

	if lhvLen != len(values) {
		stop := ctx.Expressiontuplelhv().GetStop()
		l.reportErrorf(stop,
			"incorrect number of expressions in the assignment expected: %v actual: %v",
			lhvLen,
			len(values),
		)
		return
	}

	l.processAssignments(lhvExpressions, values, inlinedecl)
}

func (l *DoSourceListener) processAssignments(expressions []parser.IExpressionlhvContext, values []value.Value, inlinedecl bool) {
	for i, expr := range expressions {
		if expr.Variableuselhv() != nil {
			l.processVariableAssignment(expr, values[i], inlinedecl)
		}
	}
}

func (l *DoSourceListener) processVariableAssignment(expr parser.IExpressionlhvContext, value value.Value, inlinedecl bool) {
	varname := expr.GetText()

	var err error
	if err = l.assignToVariable(varname, value, expr, inlinedecl); err == nil {
		return
	}

	// что если перед нами структура
	idx := strings.LastIndexByte(varname, '.')
	if !inlinedecl && idx != -1 {
		fieldname := varname[idx+1:]
		varname = varname[:idx]

		if err := l.assignToStructField(varname, fieldname, value, expr); err != nil {
			l.reportError(expr.GetStart(), err)
		}
	} else {
		// все-таки не структура
		l.reportError(expr.GetStart(), err)
	}
}

func (l *DoSourceListener) assignToVariable(varname string, value value.Value, expr parser.IExpressionlhvContext, inlinedecl bool) error {
	variable, err := l.program.GetVariable(varname)
	if err != nil {
		if inlinedecl {
			variable, err = l.program.RegisterVariable(
				l.topBlock(),
				varname,
				l.program.Typename(value.Type()),
			)

			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	varType := variable.Type().(*types.PointerType).ElemType
	valType := value.Type()

	if varType != valType {
		l.reportErrorf(expr.GetStart(),
			"type mismatch in assignment: variable '%v' has type %v, but assigned value has type %v",
			varname,
			l.program.Typename(varType),
			l.program.Typename(valType),
		)
		return nil
	}

	l.topBlock().NewStore(value, variable)
	return nil
}

func (l *DoSourceListener) assignToStructField(varname, fieldname string, value value.Value, expr parser.IExpressionlhvContext) error {
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
		l.reportErrorf(expr.GetStart(),
			"type mismatch in assignment: variable `%v.%v` has type %v, but assigned value has type %v",
			varname,
			fieldname,
			l.program.Typename(fieldType),
			l.program.Typename(valType),
		)
		return nil
	}

	fieldptr := l.topBlock().NewGetElementPtr(
		structType,
		variable,
		constant.NewInt(types.I32, 0),
		constant.NewInt(types.I32, int64(fieldidx)),
	)

	l.topBlock().NewStore(value, fieldptr)
	return nil
}

func (l *DoSourceListener) EnterStatement(ctx *parser.StatementContext) {
	l.values = append(l.values, []value.Value{})
}

func (l *DoSourceListener) ExitStatement(ctx *parser.StatementContext) {
	l.values = l.values[:len(l.values)-1]
}

func (l *DoSourceListener) EnterIfstatement(ctx *parser.IfstatementContext) {
	trueBlock := l.newBlock()
	falseBlock := l.newBlock()

	l.pushBlock(falseBlock)
	l.pushBlock(trueBlock)
}

func (l *DoSourceListener) EnterIfexpression(ctx *parser.IfexpressionContext) {
	l.pushBlock(l.blocks[len(l.blocks)-3])
}

func (l *DoSourceListener) ExitIfexpression(ctx *parser.IfexpressionContext) {
	l.popBlock()

	values := l.currentValues()

	if len(values) != 1 {
		l.reportErrorf(ctx.Expression().GetStart(),
			"if statement expected 1 bool, actual: %v",
			len(values),
		)
		return
	}

	booltype, _ := l.program.GetType("bool")
	if values[0].Type() != booltype {
		l.reportErrorf(ctx.Expression().GetStart(),
			"if statement expected bool, actual: %v",
			l.program.Typename(values[0].Type()),
		)
		return
	}

	baseBlock := l.blocks[len(l.blocks)-3]
	falseBlock := l.blocks[len(l.blocks)-2]
	trueBlock := l.blocks[len(l.blocks)-1]

	baseBlock.NewCondBr(values[0], trueBlock, falseBlock)
}

func (l *DoSourceListener) ExitIfstatement(ctx *parser.IfstatementContext) {
	falseBlock := l.blocks[len(l.blocks)-2]
	trueBlock := l.blocks[len(l.blocks)-1]

	l.popBlock() // trueBlock
	l.popBlock() // falseBlock
	l.popBlock() // baseBlock

	resBlock := l.newBlock()
	falseBlock.NewBr(resBlock)
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

func (l *DoSourceListener) ExitReturnstatement(ctx *parser.ReturnstatementContext) {
	values := l.currentValues()
	rettype := l.currfunc.Sig.RetType.(*types.StructType)

	if len(rettype.Fields) != len(values) {
		l.reportErrorf(ctx.GetStart(),
			"incorrect number of expressions in the return expected: %v actual: %v",
			len(rettype.Fields),
			len(values),
		)
		return
	}

	l.generateReturnValue(ctx, rettype, values)
}

func (l *DoSourceListener) generateReturnValue(ctx *parser.ReturnstatementContext, rettype *types.StructType, values []value.Value) {
	retvalRef := l.topBlock().NewAlloca(l.currfunc.Sig.RetType)

	for i := range rettype.Fields {
		fieldPtr := l.topBlock().NewGetElementPtr(
			rettype,
			retvalRef,
			constant.NewInt(types.I32, 0),
			constant.NewInt(types.I32, int64(i)),
		)

		if values[i].Type() != rettype.Fields[i] {
			l.reportErrorf(ctx.GetStart(),
				"return value expected: %v actual: %v",
				l.program.Typename(rettype.Fields[i]),
				l.program.Typename(values[i].Type()),
			)
			return
		}

		l.topBlock().NewStore(values[i], fieldPtr)
	}

	retval := l.topBlock().NewLoad(l.currfunc.Sig.RetType, retvalRef)

	l.topBlock().NewRet(retval)
	l.popBlock()
	l.pushBlock(l.newBlock())
}

func (l *DoSourceListener) ExitVardeclarationstatement(ctx *parser.VardeclarationstatementContext) {
	vartype := ctx.Typename().GetText()

	for _, varnameToken := range ctx.AllNAME() {
		varname := varnameToken.GetText()

		variable, err := l.program.RegisterVariable(l.topBlock(), varname, vartype)
		if err != nil {
			l.reportError(varnameToken.GetSymbol(), err)
			continue
		}

		varType := variable.Type().(*types.PointerType).ElemType
		l.topBlock().NewStore(constant.NewZeroInitializer(varType), variable)
	}
}

func (l *DoSourceListener) EnterStatementblock(ctx *parser.StatementblockContext) {
	l.program.NewScope()
}

func (l *DoSourceListener) ExitStatementblock(ctx *parser.StatementblockContext) {
	l.program.TerminateScope()
}

func (l *DoSourceListener) EnterOneexpression(ctx *parser.OneexpressionContext) {
	l.oneExpressionContext = append(l.oneExpressionContext, len(l.values))
}

func (l *DoSourceListener) ExitOneexpression(ctx *parser.OneexpressionContext) {
	newExprCnt := len(l.values) - l.oneExpressionContext[len(l.oneExpressionContext)-1]
	l.oneExpressionContext = l.oneExpressionContext[:len(l.oneExpressionContext)-1]

	if newExprCnt != 0 {
		l.reportErrorf(ctx.GetStart(), "expectedd 1 value actual: %v", newExprCnt)
	}
}

func (l *DoSourceListener) ExitLogic(ctx *parser.LogicContext) {
	cnt := len(ctx.AllAndlogic())
	if cnt == 1 {
		return
	}

	values := slices.Clone(l.currentValues()[len(l.currentValues())-cnt:])
	l.values[len(l.values)-1] = l.values[len(l.values)-1][:len(l.currentValues())-cnt]

	booltype, _ := l.program.GetType("bool")

	for i, val := range values {
		if val.Type() != booltype {
			l.reportErrorf(ctx.Andlogic(i).GetStart(),
				"expected bool actual: %v",
				l.program.Typename(val.Type()),
			)

			return
		}
	}

	// TODO: ленивое вычисление
	for len(values) > 1 {
		values[1] = l.topBlock().NewOr(
			values[0],
			values[1],
		)
		values = values[1:]
	}

	l.addValue(values[0])
}

func (l *DoSourceListener) ExitAndlogic(ctx *parser.AndlogicContext) {
	cnt := len(ctx.AllCompare())
	if cnt == 1 {
		return
	}

	values := slices.Clone(l.currentValues()[len(l.currentValues())-cnt:])
	l.values[len(l.values)-1] = l.values[len(l.values)-1][:len(l.currentValues())-cnt]

	booltype, _ := l.program.GetType("bool")

	for i, val := range values {
		if val.Type() != booltype {
			l.reportErrorf(ctx.Compare(i).GetStart(),
				"expected bool actual: %v",
				l.program.Typename(val.Type()),
			)

			return
		}
	}

	// TODO: ленивое вычисление
	for len(values) > 1 {
		values[1] = l.topBlock().NewAnd(
			values[0],
			values[1],
		)
		values = values[1:]
	}

	l.addValue(values[0])
}

func (l *DoSourceListener) ExitCompare(ctx *parser.CompareContext) {
	if len(ctx.AllOneexpression()) == 1 {
		return
	}

	values := slices.Clone(l.currentValues()[len(l.currentValues())-2:])
	l.values[len(l.values)-1] = l.values[len(l.values)-1][:len(l.currentValues())-2]

	inttype, _ := l.program.GetType("int")
	booltype, _ := l.program.GetType("bool")

	compareToken := ctx.COMPARETOKEN().GetText()
	compare := map[string]struct {
		acceptableTypes map[types.Type]struct{}
		pred            enum.IPred
	}{
		"==": {
			map[types.Type]struct{}{inttype: {}, booltype: {}},
			enum.IPredEQ,
		},
		"!=": {
			map[types.Type]struct{}{inttype: {}, booltype: {}},
			enum.IPredNE,
		},
		"<": {
			map[types.Type]struct{}{inttype: {}},
			enum.IPredSLT,
		},
		"<=": {
			map[types.Type]struct{}{inttype: {}},
			enum.IPredSLE,
		},
		">": {
			map[types.Type]struct{}{inttype: {}},
			enum.IPredSGT,
		},
		">=": {
			map[types.Type]struct{}{inttype: {}},
			enum.IPredSGE,
		},
	}[compareToken]

	if values[0].Type() != values[1].Type() {
		l.reportErrorf(ctx.GetStart(),
			"expected %v actual: %v",
			l.program.Typename(values[0].Type()),
			l.program.Typename(values[1].Type()),
		)
		return
	}

	if _, ok := compare.acceptableTypes[values[0].Type()]; !ok {
		l.reportErrorf(ctx.GetStart(),
			"type `%v` is not accepted to comprarsion with `%v`",
			l.program.Typename(values[0].Type()),
			compareToken,
		)
		return
	}
	if _, ok := compare.acceptableTypes[values[1].Type()]; !ok {
		l.reportErrorf(ctx.GetStart(),
			"type `%v` is not accepted to comprarsion with `%v`",
			l.program.Typename(values[1].Type()),
			compareToken,
		)
		return
	}

	l.addValue(
		l.topBlock().NewICmp(
			compare.pred,
			values[0],
			values[1],
		),
	)
}

func (l *DoSourceListener) ExitArithmetic(ctx *parser.ArithmeticContext) {
	cnt := len(ctx.AllMultiply())
	if cnt == 1 {
		return
	}

	values := slices.Clone(l.currentValues()[len(l.currentValues())-cnt:])
	l.values[len(l.values)-1] = l.values[len(l.values)-1][:len(l.currentValues())-cnt]

	inttype, _ := l.program.GetType("int")

	for i, val := range values {
		if val.Type() != inttype {
			l.reportErrorf(ctx.Multiply(i).GetStart(),
				"expected int actual: %v",
				l.program.Typename(val.Type()),
			)

			return
		}
	}

	sumtokens := ctx.AllSUMTOKEN()

	for len(values) > 1 {
		if sumtokens[0].GetText() == "+" {
			values[1] = l.topBlock().NewAdd(
				values[0],
				values[1],
			)
		} else {
			values[1] = l.topBlock().NewSub(
				values[0],
				values[1],
			)
		}

		values = values[1:]
		sumtokens = sumtokens[1:]
	}

	l.addValue(values[0])
}

func (l *DoSourceListener) ExitMultiply(ctx *parser.MultiplyContext) {
	cnt := len(ctx.AllOneexpression())
	if cnt == 1 {
		return
	}

	values := slices.Clone(l.currentValues()[len(l.currentValues())-cnt:])
	l.values[len(l.values)-1] = l.values[len(l.values)-1][:len(l.currentValues())-cnt]

	inttype, _ := l.program.GetType("int")

	for i, val := range values {
		if val.Type() != inttype {
			l.reportErrorf(ctx.Oneexpression(i).GetStart(),
				"expected int actual: %v",
				l.program.Typename(val.Type()),
			)

			return
		}
	}

	multokens := ctx.AllMULTTOKEN()

	for len(values) > 1 {
		if multokens[0].GetText() == "^" {
			values[1] = l.topBlock().NewMul(
				values[0],
				values[1],
			)
		} else {
			values[1] = l.topBlock().NewSDiv(
				values[0],
				values[1],
			)
		}

		values = values[1:]
		multokens = multokens[1:]
	}

	l.addValue(values[0])
}
