# Arexe ("ar-ex-ee")
Arexe is a an ***a**ritmetic **ex**pression **e**valuator* written in Go (Golang).  
It uses a technique called "[Shunting Yard Algorithm](https://en.wikipedia.org/wiki/Shunting-yard_algorithm)" to parse and evaluate an arithmetic expression and produce a correct result. The general idea is presented below. Given an arithmetic expression in string form, we parse/interpret the string and transform it into a binary tree like data structure called an "**Abstract Syntax Tree**" (AST). Once we have an AST, we simply traverse the AST, evaluating each node and propagating the result upwards.
![string to ast](https://github.com/vikyboy/arexe/blob/main/images/arexe-1.png)

![general flow](https://github.com/vikyboy/arexe/blob/main/images/arexe-2.png)
