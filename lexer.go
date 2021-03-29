package brainfcompiler

import (
	//	"fmt"
	"strings"
)

type instruction struct {
	inst      string
	magnitude int
}

func makeInstruction(inst string, magn int) instruction {

	return instruction{inst: inst, magnitude: magn}
}

func lex(code string) []instruction {

	code = strings.Replace(code, "[-]", "c", -1)
	code = strings.Replace(code, "[+]", "c", -1) // Clear cell

	currentinstruct := string(code[0])
	magnitude := 1
	var instruct instruction
	var instructionlist []instruction
	var char = ""

	for _, chr := range code[1:] {

		char = string(chr)

		if char == currentinstruct && char != "[" && char != "]" && char != "c" {
			// Increase the current magnitude, if the instruction is repeated
			magnitude++

		} else {

			// Jump out of the current instruction, add it to the slice of instructions
			instruct = makeInstruction(currentinstruct, magnitude)
			instructionlist = append(instructionlist, instruct)
			currentinstruct = char
			magnitude = 1
		}
	}

	// Add the last instruction
	instruct = makeInstruction(currentinstruct, magnitude)
	instructionlist = append(instructionlist, instruct)

	return instructionlist
}
