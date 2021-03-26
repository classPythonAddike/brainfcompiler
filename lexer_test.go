package brainfcompiler

import (
	"fmt"
	"testing"
)

// Print code neatly
func formatcode(code []instruction) string {

	result := "|"

	for _, char := range code {
		// Format as |cmd=magnitude|cmd=magnitude| . . . . |
		result += fmt.Sprintf("%v=%v|", char.inst, char.magnitude)
	}

	return result
}

func TestLexerBasic(t *testing.T) {
	fmt.Println("\n----- Test whether the lexer runs without errors -----")

	code := "[+-<.>]"
	result := formatcode(lex(code))
	want := "|[=1|+=1|-=1|<=1|.=1|>=1|]=1|"

	if result != want {
		t.Errorf("Got |%v|, want |%v|!", result, want)
	}
}

func TestLexerShorten(t *testing.T) {
	fmt.Println("\n----- Test whether the lexer can shorten code -----")

	code := "[++++>>>---->>>..<<<++,,,]"
	result := formatcode(lex(code))
	want := "|[=1|+=4|>=3|-=4|>=3|.=2|<=3|+=2|,=3|]=1|"

	if result != want {
		t.Errorf("Got %v, want %v !", result, want)
	}
}

func TestNestedLoops(t *testing.T) {
	fmt.Println("\n----- Test whether the lexer handles nested loops properly -----")

	code := "++[[--++.,]>><.-[..+++,,,]---]..>"
	result := formatcode(lex(code))
	want := "|+=2|[=1|[=1|-=2|+=2|.=1|,=1|]=1|>=2|<=1|.=1|-=1|" +
		"[=1|.=2|+=3|,=3|]=1|-=3|]=1|.=2|>=1|"

	if result != want {
		t.Errorf("Got %v, want %v !", result, want)
	}
}

func TestIgnoreCharacters(t *testing.T) {
	fmt.Println("\n----- Test whether the lexer ignores characters that are not brainf commands -----")

	code := "..[--akdjfdk..///>><=--uieui],"
	result := formatcode(lex(code))
	want := "|.=2|[=1|-=2|.=2|>=2|<=1|-=2|]=1|,=1|"

	if result != want {
		t.Errorf("Got %v, want %v !", result, want)
	}
}
