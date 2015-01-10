package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func emitProgTop() {
	fmt.Printf(`#include <stdio.h>
void main() {
	unsigned char mem[6144];
	int i;
	for (i = 0; i < 256; i++) {
		mem[i] = 0;
	}
	int pointer = 0;
`)
}

func emitProgEnd() {
	fmt.Printf(`
	return;
}
`)
}

func emitPointerIncrement() {
	fmt.Print("\tpointer += 1;\n")
}
func emitPointerDecrement() {
	fmt.Print("\tpointer -= 1;\n")
}
func emitValueIncrement() {
	fmt.Print("\tmem[pointer] += 1;\n")
}
func emitValueDecrement() {
	fmt.Print("\tmem[pointer] -= 1;\n")
}
func emitPrintValue() {
	fmt.Print("\tprintf(\"%c\", mem[pointer]);\n")
}
func emitReadValue() {
	fmt.Print("\tscanf(\"%c\", &mem[pointer]);\n")
}
func emitLoopBegin() {
	fmt.Print("\twhile (mem[pointer] != 0) {\n")
}
func emitLoopEnd() {
	fmt.Print("\t}\n")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	emitProgTop()
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err.Error())
		}
		switch r {
		case '<':
			emitPointerDecrement()
		case '>':
			emitPointerIncrement()
		case '+':
			emitValueIncrement()
		case '-':
			emitValueDecrement()
		case '.':
			emitPrintValue()
		case ',':
			emitReadValue()
		case '[':
			emitLoopBegin()
		case ']':
			emitLoopEnd()
		}
	}

	emitProgEnd()
}
