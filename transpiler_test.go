package brainfcompiler

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	code := `  
	++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.++
	+.------.--------.>>+.>++.
	`

	fmt.Println("\n----- Test A Simple HelloWorld Program -----")
	transpile(lex(code))
}

func TestCatProgram(t *testing.T) {
	code := `,[.,]`

	fmt.Println("\n----- Test a CAT Program -----")
	transpile(lex(code))

}
