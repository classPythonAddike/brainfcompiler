package brainfcompiler

import (
//	"fmt"
)

type instruction struct {
	inst      string
	magnitude int
}

func makeInstruction(inst string, magn int) instruction {

	return instruction{inst: inst, magnitude: magn}
}

func iscmd(cmd string) bool {
	var allowed = [8]string{"[", "]", ".", ",", ">", "<", "+", "-"}

	for _, char := range allowed {
		if char == cmd {
			return true
		}
	}

	return false
}

func lex(code string) []instruction {

	currentinstruct := string(code[0])
	magnitude := 1
	var instruct instruction
	var instructionlist []instruction
	var char = ""

	for _, chr := range code[1:] {

		char = string(chr)

		if iscmd(char) {

			if char == currentinstruct && char != "[" && char != "]" {
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
	}

	// Add the last instruction
	instruct = makeInstruction(currentinstruct, magnitude)
	instructionlist = append(instructionlist, instruct)

	return instructionlist
}
