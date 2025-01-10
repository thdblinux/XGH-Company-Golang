package main

import (
	"fmt"
)

func main() {
	name := getName()
	idade := 40.000
	fmt.Println("Hello", name, idade)
}

func getName() string {
	return "Thadeu"
}

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// os.Args
// 	// ARGS[0,1]
// 	if len(os.Args) != 2 {
// 		os.Exit(1)
// 	}
// 	fmt.Println("Hello", os.Args[1])
// }
