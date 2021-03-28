package brainfcompiler

import (
	"bufio"
	"os/exec"

	//"fmt"
	"io/ioutil"
	"os"
)

func compile(file string, compiler string) {

	// read from the file
	code := readFile(file)

	// lex and parse
	tokens := lex(code)
	transpiled := transpile(tokens)

	writeFile(file+".c", transpiled)

	compileCPP(file, compiler)

}

func readFile(filename string) string {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(data)
}

func writeFile(filename string, code string) {

	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	writer := bufio.NewWriter(f)
	_, err = writer.WriteString(code)

	if err != nil {
		panic(err)
	}

	writer.Flush()
}

func compileCPP(filename string, compiler string) {
	cmd := exec.Command(compiler, filename+".c", "-o", filename+".exe")

	_, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	// Delete the C++ source file

	err = os.Remove(filename + ".c")

	if err != nil {
		panic(err)
	}
}
