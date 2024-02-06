package analysis

import (
	"bufio"
	"fmt"
	"github.com/singlaanish56/jsonparsergo/pkg/errors"
	"github.com/singlaanish56/jsonparsergo/pkg/stack"
	"os"
)

var store []byte

func parseTheTokens(input string) int {
	lexer := CreateLexer(input)
	bracketStack := stack.New()

	for lexer.currentChar != 0 {
		token := lexer.GetToken()
		fmt.Printf("%s\n", token.Char)

		if token.Char == "{" || token.Char == "[" {
			bracketStack.Push(token.Char)
		} else if (token.Char == "}" && (bracketStack.IsEmpty() || *bracketStack.Pop() != "{")) || (token.Char == "]" && (bracketStack.IsEmpty() || *bracketStack.Pop() != "[")) {
			return 0
		}
	}

	if !bracketStack.IsEmpty() {
		return 0
	}
	return 1
}

func ParseTheFile(fileName string) {

	f, err := os.Open(fileName)
	errors.HandleError(err)

	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		line, _, err := reader.ReadLine()
		if errors.HandleFileError(err) {
			break
		}

		store = append(store, line...)
	}

	fmt.Println(parseTheTokens(string(store)))
}

func ParseTheString(stringName string) {

	fmt.Println(parseTheTokens(stringName))

}
