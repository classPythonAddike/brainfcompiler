# BrainF Compiler

This is a BrainF Compiler written in Golang. It converts BrainF code to C code, and then finally machine language, thus generating an executable. My aim is for it to be the fastest BrainF executor I've written yet.

As of now, this compiler is able to draw a mandelbrot set in 11 seconds. take a look at [mandelbrot.bf](/mandelbrot.bf) for the program

# Usage

## Setup

1. If you are on Windows or Debian, you can go to the Releases tab to download the latest release.

2. If you are on a different operating system, you will need to compile it on your machine (hopefully I will be able to do it for you in a few days). Create a file named `main.go` and paste the following code into it:
```go
package main

import (
  "flag"
  bfcompiler "github.com/classPythonAddike/brainfcompiler"
)

func main() {

  compiler := flag.String("compiler", "gcc", "Compiler to use at runtime")
  file := flag.String("filename", "", "BrainF file to compile")
  output := flag.String("out", "", "Output file")

  flag.Parse()

  bfcompiler.Compile(*file, *compiler, *output)
}
```

3. If you don't want to use the CLI, you can always invoke it from a script! You just need to download this package with `go get github.com/classPythonAddike/brainfcompiler` and then use the `Compile` function to compile a BrainF program.

## Usage

You can invoke the script like this:
```sh
> brainfcompiler.exe -filename=filename.bf -out=filename.exe [-compiler="g++"]
```
The compiler used by default is `gcc`. Supported compilers are `gcc` and `g++`
