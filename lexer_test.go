package brainfcompiler

import (
	"fmt"
	"testing"
)

// Print code neatly
func printcode(code []instruction) {
	for _, char := range code {
		fmt.Printf("%v - %v\n", char.inst, char.magnitude)
	}
}

func TestLexerBasic(t *testing.T) {
	fmt.Println("\n----- Test whether the lexer runs without errors -----")

	code := "[+-<.>.,]"
	result := lex(code)

	fmt.Println(code)
	printcode(result)
}

func TestLexerShorten(t *testing.T) {
	fmt.Println("\n----- Test whether the lexer can shorten code -----")

	code := "[++++>>>---->>>..<<<++,,,]"
	result := lex(code)

	fmt.Println(code)
	printcode(result)
}

func TestNestedLoops(t *testing.T) {
	fmt.Println("\n----- Test whether the lexer handles nested loops properly -----")

	code := "++[..[--++.,]>><.-[..+++,,,]---]..>"
	result := lex(code)

	fmt.Println(code)
	printcode(result)
}
