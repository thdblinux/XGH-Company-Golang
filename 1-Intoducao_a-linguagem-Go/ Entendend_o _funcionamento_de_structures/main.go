package main

import (
	"fmt"
)

type Pessoa struct {
	nome    string
	idade   int
	salario int
}

func main() {
	pessoa := &Pessoa{
		nome:    "Thadeu",
		idade:   40,
		salario: 100,
	}
	addSalaryPessoa(pessoa, 10)
	fmt.Println(pessoa.salario)

	// name, salario := "Thadeu", 100
	// setName(name)
	// newSalary, bonus := addSalary(salario, 10)
	// fmt.Println("Novo salário:", newSalary)
	// fmt.Println("Novo bonus:", bonus)
}

func addSalaryPessoa(p *Pessoa, bonus int) {
	p.salario += bonus
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
