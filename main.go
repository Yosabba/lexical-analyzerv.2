package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	IDENTIFIER  = "identifier"
	INT_LITERAL = "int_literal"
	OPERATOR    = "operator"
	DELIMITER   = "delimiter"
)

type Token struct {
	Lexeme string
	Type   string
}

func readUserFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func main() {
	var fileName string

	fmt.Println("What is the name of your file:?\nPlease add the extension (i.e: example.txt)")
	fmt.Scanln(&fileName)

	contentOfFile := readUserFile(fileName)
	tokens := tokenizeFile(contentOfFile)

	//lexeme & token needs to be be in a table so format it to be in a table with a 10 char min
	fmt.Printf("%-10s %-10s\n", "Lexeme", "Token")
	fmt.Println(strings.Repeat("-", 20))

	for _, token := range tokens {
		fmt.Printf("%-10s %-10s\n", token.Lexeme, token.Type)
	}

}

func isCharOperator(character rune) bool {
	switch character {
	case '+', '-', '*', '/', '=':
		return true
	default:
		return false
	}
}

func isCharDelimiter(character rune) bool {
	switch character {
	case '(', ')', '{', '}', '[', ']', ';', '_':
		return true
	default:
		return false
	}
}

func isNumber(word string) bool {
	for _, char := range word {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func isIdentifier(word string) bool {
	if len(word) == 0 {
		return false
	}
	for i, char := range word {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char != '_') && (i == 0 || (char < '0' || char > '9')) {
			return false
		}
	}
	return true
}

func determineTokenType(lexeme string) string {
	if isNumber(lexeme) {
		return INT_LITERAL
	}
	return IDENTIFIER
}

func tokenizeFile(input string) []Token {
	var tokens []Token
	current := ""

	for _, char := range input {
		if isCharOperator(char) {
			if current != "" {
				tokens = append(tokens, Token{current, determineTokenType(current)})
				current = ""
			}

			tokens = append(tokens, Token{string(char), OPERATOR})
		} else if isCharDelimiter(char) {

			if current != "" {
				tokens = append(tokens, Token{current, determineTokenType(current)})
				current = ""
			}

			tokens = append(tokens, Token{string(char), DELIMITER})
		} else {
			current += string(char)
		}
	}

	if current != "" {
		tokens = append(tokens, Token{current, determineTokenType(current)})
	}

	return tokens
}


