package main

import "fmt"

type Empregado struct {
	Nome string
	id   int
}

func main() {
	emp := Empregado{"Thadeu", 123}
	pts := &emp

	fmt.Println(pts)
	pts.Nome = "Marcio"
	fmt.Println("Valor pts", pts)
}
