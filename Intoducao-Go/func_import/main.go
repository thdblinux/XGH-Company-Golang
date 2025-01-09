package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args
	// ARGS[0,1]
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("Hello", os.Args[1])
}
