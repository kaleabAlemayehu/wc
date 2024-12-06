package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	args := os.Args
	nc := 0
	byt, err := os.ReadFile(args[len(args)-1])
	if err != nil {
		panic(err)
	}
	for range byt {
		if slices.Contains(args, "-c") {
			nc++
		}
	}
	fmt.Printf("%d %v\n", nc, args[len(args)-1])

}
