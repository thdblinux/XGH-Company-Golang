package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// if len(os.Args) != 2 {
	// 	os.Exit(1)

	// }
	// n, err := strconv.Atoi(os.Args[1])
	// if err != nil {
	// 	fmt.Println("erro, nao é um numero válido")
	// 	os.Exit(1)
	// }
	// fmt.Println("numeo convertido:", n)

	// pessoa := &funcionarios.Pessoa{
	// 	nome:    "Thadeu",
	// 	idade:   40,
	// 	salario: 100,
	// }

	// salFunc := make(map[string]int)
	// salFunc["Marcio"] = 10
	// salFunc["Thadeu"] = 20

	// sal, exist := salFunc["Marcio"]
	// fmt.Println(sal, exist)
	// totalSal := len(salFunc)
	// fmt.Println("Total de funcionários:", totalSal)
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

func setName(name string) {
	fmt.Println("Hello", name)
}

func addSalary(valorAtual int, bonus int) (int, int) {
	return valorAtual + bonus, bonus
}

func getName() string {
	return "Thadeu"
}
