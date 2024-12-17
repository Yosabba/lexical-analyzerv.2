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
	if (err != nil) {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := os.ReadFile(fileName)
	if (err != nil) {
		log.Fatal(err)
	}

	return string(content)
}

func writeTokensToFile(tokens []Token, fileName string) {
	file, err := os.Create(fileName)
	if (err != nil) {
		log.Fatal(err)
	}
	defer file.Close()

	for _, token := range tokens {
		_, err := file.WriteString(fmt.Sprintf("%-10s %-10s\n", token.Lexeme, token.Type))
		if (err != nil) {
			log.Fatal(err)
		}
	}
}

func main() {
	var inputFileName, outputFileName string

	fmt.Println("What is the name of your input file?\nPlease add the extension (i.e: example.txt)")
	fmt.Scanln(&inputFileName)

	contentOfFile := readUserFile(inputFileName)
	tokens := tokenizeFile(contentOfFile)

	fmt.Println("What do you want to name your output file?\nPlease add the extension of file you want also. (i.e: example.txt, example.docx)")
	fmt.Scanln(&outputFileName)

	writeTokensToFile(tokens, outputFileName)

	// Print tokens to console
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


