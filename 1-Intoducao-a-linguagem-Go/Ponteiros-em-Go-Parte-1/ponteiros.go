package main

import "fmt"

func main() {
	var inteiro = 45
	var ponteiro *int = &inteiro
	fmt.Println("Valor da variavel inteiro:", inteiro)
	fmt.Println("Endereço da variavel inteiro:", &inteiro)
	fmt.Println("Valo da variavel ponteiro:", ponteiro)
	fmt.Println("Valor da variavel ponteiro pelo ponteio :", *ponteiro)
}
