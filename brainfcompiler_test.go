package brainfcompiler

import (
	"fmt"
	"testing"
)

func TestCompileHelloWorld(t *testing.T) {
	fmt.Println("\n----- Test a simple cat program -----")

	Compile("cat.bf", "gcc", "cat.exe")
}

func TestCompileMandelbrot(t *testing.T) {
	fmt.Println("\n----- Test a Mandelbrot Drawing Program -----")

	Compile("mandelbrot.bf", "gcc", "cat.exe")
}
