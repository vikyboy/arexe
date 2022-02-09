package main

import (
	"fmt"
	"bufio"
	"os"
	
	"example.org/lexer"
	"example.org/parser"
)

func main(){
	var input string
	consoleScanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	fmt.Println("Enter Expression:")
	
	for consoleScanner.Scan(){
		fmt.Println("You Entered:", consoleScanner.Text())
		input = consoleScanner.Text()
		tokens, _ := lexer.Tokenize(input)
		fmt.Println("Tokens:", tokens)
		fmt.Printf("Result:\t\t%v\n",parser.GetAst(parser.GetPostfixNotation(tokens)).Evaluate())
	} 
}//end main
