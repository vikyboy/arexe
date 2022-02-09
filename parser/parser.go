package parser

import (
	//"fmt"
	"example.org/lexer"
	"example.org/fifolifo"
	"example.org/ast"
)

func operatorPrecedenceHigher (curr, topOfStack string) bool {
	opMap := map[string]int{"-":0,"+":1,"*":2,"/":3,"(":0}
	
	if opMap[curr] > opMap[topOfStack] {
		return true
	}
	return false
}

func GetPostfixNotation(tokens []lexer.Token)[]lexer.Token {

	tokQueuePtr := fifolifo.NewFifo()
	tokStackPtr := fifolifo.NewLifo()

	for _, tok := range tokens { //iterate over all tokens		
		if tok.Type == "OP" { // operator
			//TODO operator precedence
			for len(*tokStackPtr) >= 0 {
				if len(*tokStackPtr) == 0 { // chen stack is empty
					tokStackPtr.Push(tok) // push on stack
					break
				}
				toS := tokStackPtr.Peek()
				
				// if current operator has higher precedence than top of stack operator, add it to the stack
				if operatorPrecedenceHigher(tok.Literal, toS.Literal) { 
					tokStackPtr.Push(tok) // push on stack
					break
				}else{ //else pop operator from stack and add to queue. then repeat
					poppedItem:=tokStackPtr.Pop() // pop the stack
					tokQueuePtr.AddItem(poppedItem) //add popped item to queue					
				}
			}//end for
		}else if tok.Type == "INT" {
			tokQueuePtr.AddItem(tok)// add to queue
		}else if tok.Type == "LPAREN"{
			tokStackPtr.Push(tok)// push on stack
		}else if tok.Type == "RPAREN" {
			// keep popping stack and adding elements to the queue till a LPAREN is found
			// LPAREN, RPAREN are then discarded
			for len(*tokStackPtr) > 0 {
				poppedItem := tokStackPtr.Pop() // pop the stack
				if poppedItem.Type == "LPAREN" {
					break	
				}
				tokQueuePtr.AddItem(poppedItem) //add to queue
			}//end for
		}			

	}//end for

	//pop everything from stack and add it to queue
	for len(*tokStackPtr) > 0 {
		poppedItem := tokStackPtr.Pop() // pop the stack
		tokQueuePtr.AddItem(poppedItem) // add to queue
	}		
			
	return *tokQueuePtr
}//end func GetPostFixNotation

func GetAst(pfTokens []lexer.Token)*ast.Ast {

	astStackPtr := fifolifo.NewAstStack()
	//read the postfix token list
	for _, tok := range pfTokens {
		if tok.Type == "INT" {
			astStackPtr.Push(&ast.Ast{Operator:tok}) // Left and Right are left as nil
		}else if tok.Type == "OP" {
			right := astStackPtr.Pop()
			left := astStackPtr.Pop()
			astStackPtr.Push(&ast.Ast{Operator:tok,Left:left,Right:right})
		}
	}//end for
	//after for loop is done, only one ast node should remain in the stack (the root node)
	rootAst := astStackPtr.Pop()	
	return rootAst
}//end func GetAst


