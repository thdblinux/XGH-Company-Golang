package main

import "fmt"

func pointerFunction(a *int) {
	*a = 480
}
func main() {
	var x = 100
	fmt.Printf("o valor da variavel x antes da funcao é: %v\n", x)
	var pa *int = &x

	pointerFunction(pa)
	fmt.Printf("o valor de x depois da chamada da funcao é: %v\n", x)

}
