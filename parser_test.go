package brainfcompiler

import (
	"fmt"
	"testing"
)

func TestSimpleCode(t *testing.T) {
	code := `  
	++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.++
	+.------.--------.>>+.>++.
	`
	fmt.Println(parse(lex(code)))
}
