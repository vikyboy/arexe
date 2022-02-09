package lexer

import (
	"fmt"
	"strings"
	"text/scanner"
	"strconv"
)

type Token struct {
	Literal string
	Type string
}

func Tokenize(s string) ([]Token, error) {
	var tokenScanner scanner.Scanner
	var tokens []Token
	tokenScanner.Init(strings.NewReader(s))
	unaryOperator := ""
	previousTok := Token{}
	for tok:= tokenScanner.Scan();tok != scanner.EOF; tok = tokenScanner.Scan() {
		//fmt.Println("TOKEN:",tokenScanner.TokenText(), tok)
		literal := tokenScanner.TokenText()
		if literal == "*" || literal == "/" || literal == "+" || literal == "-" {
			if (literal == "-" || literal == "+" )&& previousTok.Type != "INT"{
				unaryOperator = literal
				continue
			}
			tokens = append(tokens, Token{Literal:literal, Type:"OP"})
		}else if literal == "(" {
			tokens = append(tokens, Token{Literal:literal, Type:"LPAREN"})
		}else if literal == ")" {
			tokens = append(tokens, Token{Literal:literal, Type:"RPAREN"})
		}else if _, err := strconv.Atoi(literal); err == nil {
			if unaryOperator != "" {
				tokens = append(tokens, Token{Literal:unaryOperator+literal, Type:"INT"})
				unaryOperator = ""
			}else{
				tokens = append(tokens, Token{Literal:literal, Type:"INT"})				
			}
		}else{
			return tokens, fmt.Errorf("ERROR: Invalid Token %v", literal)
		}
		previousTok = tokens[len(tokens)-1]
	}//end for
	return tokens, nil
}
