package fifolifo

import "example.org/lexer"
import "example.org/ast"


type fifo []lexer.Token
type lifo []lexer.Token
type astStack []*ast.Ast


//FIFO
func NewFifo()*fifo {
	return &fifo{}
}

func (f *fifo)AddItem(i lexer.Token){
	*f = append(*f, i)
}


//LIFO
func NewLifo()*lifo {
	return &lifo{}
}

func (l *lifo)Push(i lexer.Token) {
	*l = append(*l, i)
}

func (l *lifo)Pop()lexer.Token {
	lastIdx := len(*l)-1
	poppedItem := (*l)[lastIdx]
	(*l)[lastIdx] = lexer.Token{} // suggested tip to prevent memory leak
	*l = (*l)[:lastIdx] // pop the item from slice
	return poppedItem
}

func (l *lifo)Peek()lexer.Token {
	lastIdx := len(*l)-1
	return (*l)[lastIdx]
}

//Ast
func NewAstStack()*astStack  {
	return &astStack{}
}

func (a *astStack) Push(i *ast.Ast){
	*a = append(*a, i)
}

func (a *astStack)Pop()*ast.Ast {
	lastIdx := len(*a)-1
	poppedItem := (*a)[lastIdx]
	(*a)[lastIdx] = nil // suggested tip to prevent memory leak
	*a = (*a)[:lastIdx] // pop the item from slice
	return poppedItem
}

