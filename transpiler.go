package brainfcompiler

import "fmt"

var template = `#include <stdio.h>
char m[30000]={0};int p=0;int o=30000;int main(){`

func transpile(code []instruction) string {

	// Convert brainf to C code

	output := ""

	for _, char := range code {

		chr := char.inst

		switch chr {

		case ">":
			output += fmt.Sprintf("p+=%v;p=p", char.magnitude) + "%" + "o;"
		case "<":
			output += fmt.Sprintf("p-=%v;p=p", char.magnitude) + "%" + "o;"
		case ".":
			output += "putchar(m[p]);"
		case ",":
			output += "m[p]=getchar();"
		case "+":
			output += fmt.Sprintf("m[p]+=%v;", char.magnitude)
		case "-":
			output += fmt.Sprintf("m[p]-=%v;", char.magnitude)
		case "[":
			output += "while(m[p]){"
		case "]":
			output += "}"
		case "c":
			output += "m[p]=0;"
		}
	}

	return template + output + "}"
}
