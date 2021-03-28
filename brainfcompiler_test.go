package brainfcompiler

import (
	"fmt"
	"testing"
)

func TestCompileHelloWorld(t *testing.T) {
	fmt.Println("\n----- Test a simple cat program -----")

	compile("cat.bf", "gcc")
}

func TestCompileMandelbrot(t *testing.T) {
	fmt.Println("\n----- Test a Mandelbrot Drawing Program -----")

	compile("mandelbrot.bf", "gcc")
}
