package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	nargs := len(os.Args)
	for i := 1; i < nargs; i++ {
		fmt.Println(os.Args[i])
	}
}
