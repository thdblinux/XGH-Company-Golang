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
	salarios := []int{}
	// salarios := make([]int, 5) // Use essa linha se quiser criar um slice vazio com capacidade inicial.

	// Atualizando os salários, somando 101 ao valor original
	for i := 0; i < 10; i++ {
		salarios = append(salarios, 100+i)
	}
	// Imprimindo os salários atualizados
	for _, salario := range salarios {
		fmt.Println(salario)
	}
	// pessoa2 := new(Pessoa)
	// pessoa2.nome = "Marcio"
	// pessoa2.idade = 40
	// pessoa2.salario = 100
	// pessoa := &Pessoa{
	// 	nome:    "Thadeu",
	// 	idade:   40,
	// 	salario: 100,
	// }
	// pessoa.addSalaryPessoa(10)
	// fmt.Println(pessoa.salario)

	// name, salario := "Thadeu", 100
	// setName(name)
	// newSalary, bonus := addSalary(salario, 10)
	// fmt.Println("Novo salário:", newSalary)
	// fmt.Println("Novo bonus:", bonus)
}

func (p *Pessoa) addSalaryPessoa(bonus int) {
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
