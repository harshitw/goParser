package main

import (
	"fmt"
	"os"

	"github.com/harshitw/goParser/src/lexer"
)

func main() {

	/*
		Parser works on mordern syntax language, OOPs concepts, static types,
		implicit and expllicity type inference, array and string literals,
		control flow patterns like if, for each etc and imports.

		Tokenisation - process of translating source code to meaningful tokens
		we can understand.
		AST - Abstract syntax tree data structure represents our program.
		We create an AST by building lexer, and then do parsing to generate parser.
		Using parser we can do code generation, interpretation etc.

		This project is a building block of functional language, building interpretor
		and building compiler.

		Using this project we can build parser for any language.

	*/

	bytes, _ := os.ReadFile("./examples/00.lang")
	source := string(bytes)

	fmt.Printf("Code : %+v\n", source)

	tokens := lexer.Tokenize(source)
	for _, token := range tokens {
		token.Debug()
	}

}
