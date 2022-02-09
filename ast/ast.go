package ast

import (
	"example.org/lexer"
	"strconv"
	"log"
)

type Ast struct {
	Operator lexer.Token
	Left *Ast
	Right *Ast
}

func (ast *Ast)Evaluate()int {
	if ast.Operator.Type == "INT" {
		val, err := strconv.Atoi(ast.Operator.Literal)
		if err != nil {
			log.Fatal(err)
		}
		return val
	}else if ast.Operator.Type == "OP" {
	
		switch ast.Operator.Literal {
			case "-" :
				return ast.Left.Evaluate() - ast.Right.Evaluate()
			case "+" :
				return ast.Left.Evaluate() + ast.Right.Evaluate()
			case "*" :
				return ast.Left.Evaluate() * ast.Right.Evaluate()
			case "/" :
				return ast.Left.Evaluate() / ast.Right.Evaluate()
		}//end switch

	}
	return -1
}





