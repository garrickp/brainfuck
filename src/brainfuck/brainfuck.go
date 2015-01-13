package main

import (
	"bufio"
	"fbp"
	"fmt"
	"io"
	"os"
	"sync"
)

const (
	PDec = iota
	PInc
	VDec
	VInc
	Emit
	Read
	StartL
	EndL
)

var wg sync.WaitGroup

func ToStdout(in fbp.Connection) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for ip := range in {
			if ip.IsConnEnd() {
				break
			}
			s := ip.String()
			fmt.Printf(s)
		}
	}()
}

func WrapProg(in fbp.Connection) fbp.Connection {
	out := fbp.NewConnection()
	wg.Add(1)
	go func() {
		defer wg.Done()
		out <- fbp.NewIPString("#include <stdio.h>\nvoid main() {\nunsigned char mem[6144];\nint i;\nfor (i = 0; i < 6144; i++) {\nmem[i] = 0;\n}\nint pointer = 0;\n")
		for ip := range in {
			if ip.IsConnEnd() {
				break
			}
			out <- ip
		}

		out <- fbp.NewIPString("}\n")
		out <- fbp.NewIPConnEnd()
	}()
	return out
}

func TokenToCommand(in fbp.Connection) fbp.Connection {
	out := fbp.NewConnection()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for ip := range in {
			if ip.IsConnEnd() {
				out <- ip
				break
			}
			r := ip.Int()
			switch r {
			case PDec:
				out <- fbp.NewIPString("pointer -= 1;\n")
			case PInc:
				out <- fbp.NewIPString("pointer += 1;\n")
			case VInc:
				out <- fbp.NewIPString("mem[pointer] += 1;\n")
			case VDec:
				out <- fbp.NewIPString("mem[pointer] -= 1;\n")
			case Emit:
				out <- fbp.NewIPString("putchar(mem[pointer]);\n")
			case Read:
				out <- fbp.NewIPString("mem[pointer] = getchar();\n")
			case StartL:
				out <- fbp.NewIPString("while (mem[pointer] != 0) {\n")
			case EndL:
				out <- fbp.NewIPString("}\n")
			}
		}
	}()
	return out
}

func RuneToToken(in fbp.Connection) fbp.Connection {
	out := fbp.NewConnection()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for ip := range in {
			if ip.IsConnEnd() {
				out <- ip
				break
			}
			r := ip.Rune()
			switch r {
			case '<':
				out <- fbp.NewIPInt(PDec)
			case '>':
				out <- fbp.NewIPInt(PInc)
			case '+':
				out <- fbp.NewIPInt(VInc)
			case '-':
				out <- fbp.NewIPInt(VDec)
			case '.':
				out <- fbp.NewIPInt(Emit)
			case ',':
				out <- fbp.NewIPInt(Read)
			case '[':
				out <- fbp.NewIPInt(StartL)
			case ']':
				out <- fbp.NewIPInt(EndL)
			}
		}
	}()
	return out
}

func StreamReadRune(stream io.Reader) fbp.Connection {
	out := fbp.NewConnection()
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stream)
		for {
			r, _, err := reader.ReadRune()
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
			out <- fbp.NewIPRune(r)
		}
		out <- fbp.NewIPConnEnd()
	}()
	return out
}

func main() {
	a := StreamReadRune(os.Stdin)
	b := RuneToToken(a)
	c := TokenToCommand(b)
	d := WrapProg(c)
	ToStdout(d)

	wg.Wait()
}
