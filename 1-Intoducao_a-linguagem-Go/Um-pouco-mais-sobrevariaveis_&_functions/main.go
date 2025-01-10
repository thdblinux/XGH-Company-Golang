package main

import (
	"fmt"
)

func main() {
	name, salario := "Thadeu", 100
	setName(name)
	newSalary, bonus := addSalary(salario, 10)
	fmt.Println("Novo salário:", newSalary)
	fmt.Println("Novo bonus:", bonus)
}

func setName(name string) {
	fmt.Println("Hello", name)
}

func addSalary(valorAtual int, bonus int) (int, int) {
	return valorAtual + bonus, bonus
}

func getName() string {
	return "Thadeu"
}
