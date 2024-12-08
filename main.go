package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

const OUT = 0
const IN = 1

func main() {
	var nb, nl, nw, nc, state int
	args := os.Args
	wordInc := slices.Contains(args, "w")
	byteInc := slices.Contains(args, "c")
	lineInc := slices.Contains(args, "l")
	charInc := slices.Contains(args, "m")
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if info.Mode()&os.ModeCharDevice == 0 {
		reader := bufio.NewReader(os.Stdin)
		byt, err := reader.ReadBytes('0')
		if err != nil && err != io.EOF {
			panic(err)
		}
		for _, b := range byt {
			nb++
			nc++
			if b == 10 {
				nl++
			}
			if b == 10 || b == 32 || b == 9 {
				state = OUT
			} else if state == OUT {
				nw++
				state = IN
			}
		}

	} else {

		byt, err := os.ReadFile(args[len(args)-1])
		if err != nil {
			panic(err)
		}
		for _, b := range byt {
			nb++
			nc++
			if b == 10 {
				nl++
			}
			if b == 10 || b == 32 || b == 9 {
				state = OUT
			} else if state == OUT {
				nw++
				state = IN
			}
		}
	}
	if wordInc && byteInc && lineInc && charInc || !(wordInc && byteInc && lineInc && charInc) {
		fmt.Printf(" %d  %d %d %d %v\n", nl, nw, nb, nc, args[len(args)-1])
		return
	}
	if wordInc && byteInc && !lineInc {
		fmt.Printf(" %d %d %v\n", nw, nb, args[len(args)-1])
		return
	}
	if wordInc && !byteInc && lineInc {
		fmt.Printf(" %d %d %v\n", nl, nw, args[len(args)-1])
		return
	}
	if !wordInc && byteInc && lineInc {
		fmt.Printf(" %d %d %v\n", nl, nb, args[len(args)-1])
		return
	}

}
