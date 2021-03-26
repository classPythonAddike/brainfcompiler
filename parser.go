package brainfcompiler

import "fmt"

var template = `
#include <stdio.h>
#include <iostream>
#include <vector>

using namespace std;

int main(){
vector<int> memory_blocks;
int pointer = 0;
memory_blocks.push_back(0);
`

func parse(code []instruction) string {

	// Convert brainf to C++ code

	output := ""

	for _, char := range code {

		chr := char.inst

		if chr == "[" { // [ -> while loop
			output += "while (memory_blocks[pointer] != 0){"
		} else if chr == "]" { // ] -> end loop
			output += "}"
		} else if chr == "," { // , -> input once, as input overwrites
			output += "memory_blocks[pointer] = (int)getchar();"
		} else if chr == "." {
			output += fmt.Sprintf("for (int i = 0; i < %v; i++){cout << (char)memory_blocks[pointer];}", char.magnitude)
		} else if chr == "<" || chr == ">" {
			m := make(map[string]string)
			m[">"] = "+"
			m["<"] = "-"
			output += fmt.Sprintf("pointer %v= %v;", m[chr], char.magnitude)
		} else if chr == "+" || chr == "-" {
			output += fmt.Sprintf("memory_blocks[pointer] %v= %v;", chr, char.magnitude)
		}
	}

	return template + output + "}"
}
