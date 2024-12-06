package main

import (
	"fmt"
	"os"
	"slices"
)

const OUT = 0
const IN = 1

func main() {
	args := os.Args
	var nc, nl, nw, state int
	byt, err := os.ReadFile(args[len(args)-1])
	if err != nil {
		panic(err)
	}
	wordInc := slices.Contains(args, "w")
	charInc := slices.Contains(args, "c")
	lineInc := slices.Contains(args, "l")
	for _, b := range byt {
		if charInc {
			nc++
		}
		if lineInc {
			if b == 10 {
				nl++
			}
		}
		if wordInc {
			if b == 10 || b == 32 || b == 9 {
				state = OUT
			} else if state == OUT {
				nw++
				state = IN
			}
		}
	}
	if wordInc && charInc && lineInc || !(wordInc && charInc && lineInc) {
		fmt.Printf(" %d  %d %d %v\n", nl, nw, nc, args[len(args)-1])
		return
	}
	if wordInc && charInc && !lineInc {
		fmt.Printf(" %d %d %v\n", nw, nc, args[len(args)-1])
		return
	}
	if wordInc && !charInc && lineInc {
		fmt.Printf(" %d %d %v\n", nl, nw, args[len(args)-1])
		return
	}
	if !wordInc && charInc && lineInc {
		fmt.Printf(" %d %d %v\n", nl, nc, args[len(args)-1])
		return
	}

}
