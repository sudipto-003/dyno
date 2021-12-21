package evaluator

import (
	"github.com/sudipto-003/dyno/ast"
	"github.com/sudipto-003/dyno/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node.Statements)

	case *ast.ReturnStatement:
		val := Eval(node.ReturnedValue)
		return &object.ReturnValue{Value: val}

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		return naitveBoolToBooleanObject(node.Value)

	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)

	case *ast.BlockStatements:
		return evalBlockStatements(node)

	case *ast.IfExpression:
		return evalIfExpression(node)
	}

	return nil
}

func evalProgram(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt)

		if returnedValue, ok := result.(*object.ReturnValue); ok {
			return returnedValue.Value
		}
	}

	return result
}

func evalBlockStatements(block *ast.BlockStatements) object.Object {
	var result object.Object

	for _, stmt := range block.Statements {
		result = Eval(stmt)

		if result != nil && result.Type() == object.RETURN_VALUE_OBJ {
			return result
		}
	}

	return result
}

func naitveBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}

	return FALSE
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)

	case "-":
		return evalMinusPrefixOperator(right)

	default:
		return NULL
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusPrefixOperator(right object.Object) object.Object {
	integer, ok := right.(*object.Integer)
	if !ok {
		return NULL
	}

	return &object.Integer{Value: -integer.Value}
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)

	case operator == "==":
		return naitveBoolToBooleanObject(left == right)
	case operator == "!=":
		return naitveBoolToBooleanObject(left != right)
	default:
		return NULL
	}
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return naitveBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return naitveBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return naitveBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return naitveBoolToBooleanObject(leftVal != rightVal)
	default:
		return NULL
	}
}

func evalIfExpression(ie *ast.IfExpression) object.Object {
	condition := Eval(ie.Condition)

	if isTruthy(condition) {
		return Eval(ie.Consequence)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative)
	} else {
		return NULL
	}
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}
