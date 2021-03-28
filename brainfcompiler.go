package brainfcompiler

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func Compile(file string, compiler string, output string) {

	log.SetPrefix("BrainF Compiler:")

	if file == "" {
		log.Fatal("Error -  No input files specified!")
	}

	if output == "" {
		log.Fatal("Error - No output files specified!")
	}

	log.Println("Reading file")
	code := readFile(file)

	// lex and parse
	log.Println("Generating tokens")
	tokens := lex(code)

	log.Println("Transpiling code")
	transpiled := transpile(tokens)

	log.Println("Writing to " + file + ".c")
	writeFile(file+".c", transpiled)

	log.Println("Compiling C code to " + output)
	compileCPP(file, compiler, output)

	log.Println("Done!")

}

func readFile(filename string) string {

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func writeFile(filename string, code string) {

	f, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	writer := bufio.NewWriter(f)
	_, err = writer.WriteString(code)

	if err != nil {
		log.Fatal(err)
	}

	writer.Flush()
}

func compileCPP(filename string, compiler string, output string) {
	cmd := exec.Command(compiler, filename+".c", "-o", output)

	_, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	// Delete the C++ source file

	err = os.Remove(filename + ".c")

	if err != nil {
		log.Fatal(err)
	}
}
