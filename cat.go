// Unix cat clone

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	nargs := len(os.Args)
	for i := 1; i < nargs; i++ {
		cont, err := ioutil.ReadFile(os.Args[i])
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", cont)
	}
}
