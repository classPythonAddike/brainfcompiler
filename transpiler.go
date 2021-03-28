package brainfcompiler

import "fmt"

var template = `
#include <stdio.h>

char memory_blocks[30000] = {0};
int pointer = 0;

int main() {
`

func transpile(code []instruction) string {

	// Convert brainf to C++ code

	output := ""

	for _, char := range code {

		chr := char.inst

		switch chr {

		case ">":
			output += fmt.Sprintf("pointer += %v;pointer = pointer", char.magnitude) + " % 30000;"
		case "<":
			output += fmt.Sprintf("pointer -= %v;pointer = pointer", char.magnitude) + " % 30000;"
		case ".":
			output += "putchar(memory_blocks[pointer]);"
		case ",":
			output += "memory_blocks[pointer]=getchar();"
		case "+":
			output += fmt.Sprintf("memory_blocks[pointer] += %v;", char.magnitude)
		case "-":
			output += fmt.Sprintf("memory_blocks[pointer] -= %v;", char.magnitude)
		case "[":
			output += "while (memory_blocks[pointer]){"
		case "]":
			output += "}"
		}
	}

	return template + output + "\n}"
}
