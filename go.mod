module example.org/calc

go 1.17

require example.org/lexer v0.0.0
require example.org/parser v0.0.0
require example.org/fifolifo v0.0.0
require example.org/ast v0.0.0

replace example.org/lexer => ./lexer
replace example.org/parser => ./parser
replace example.org/fifolifo => ./fifolifo
replace example.org/ast => ./ast
