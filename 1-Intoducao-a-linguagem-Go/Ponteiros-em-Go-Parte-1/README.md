# Aula Golang: Ponteiros

Este projeto demonstra o uso de **ponteiros** em Golang, um conceito essencial para manipular referências de memória. O exemplo detalha como declarar variáveis, criar ponteiros e acessar valores e endereços usando ponteiros.

---

## Conceitos Abordados

1. **Variáveis e Endereços de Memória**
   - Como acessar o endereço de memória de uma variável usando o operador `&`.

2. **Ponteiros**
   - Como declarar e inicializar ponteiros.
   - Como acessar valores armazenados no endereço referenciado por um ponteiro usando o operador `*`.

3. **Desreferenciamento**
   - Como obter o valor armazenado no endereço apontado por um ponteiro.

---

## Código do Projeto

O código do arquivo `ponteiros.go` é o seguinte:

```go
package main

import "fmt"

func main() {
    // Declaração de uma variável inteira
    var inteiro = 45

    // Declaração de um ponteiro que armazena o endereço da variável "inteiro"
    var ponteiro *int = &inteiro

    // Exibição do valor da variável "inteiro"
    fmt.Println("Valor da variável inteiro:", inteiro)

    // Exibição do endereço de memória da variável "inteiro"
    fmt.Println("Endereço da variável inteiro:", &inteiro)

    // Exibição do valor armazenado no ponteiro (endereço da variável "inteiro")
    fmt.Println("Valor da variável ponteiro:", ponteiro)

    // Exibição do valor da variável "inteiro" usando o ponteiro (desreferenciamento)
    fmt.Println("Valor da variável ponteiro pelo ponteiro:", *ponteiro)
}
```

## Explicação Detalhada
**1. Declaração de Variáveis**

```
var inteiro = 45
```
- Declara uma variável `inteiro` e atribui o valor `45`.
- O valor é armazenado em um endereço de memória específico.
  
```
var ponteiro *int = &inteiro
```
- Declara uma variável `ponteiro` do tipo `*int` (ponteiro para um inteiro).
- Inicializa o ponteiro com o endereço de memória da variável `inteiro `usando o operador `&`.

**2. Exibição de Valores e Endereços**
Valor da variável:
```
fmt.Println("Valor da variável inteiro:", inteiro)
```
- Imprime o valor armazenado na variável `inteiro`, que é `45`.
  
**Endereço de memória da variável:**
```
fmt.Println("Endereço da variável inteiro:", &inteiro)
```
- Imprime o endereço de memória onde o valor `45` está armazenado.
  
**Valor armazenado no ponteiro:**
```
fmt.Println("Valor da variável ponteiro:", ponteiro)
```

- Imprime o valor armazenado no ponteiro `ponteiro`, que é o endereço de memória da variável `inteiro`.
**Valor do endereço apontado pelo ponteiro:**
```
fmt.Println("Valor da variável ponteiro pelo ponteiro:", *ponteiro)
```
- Imprime o valor armazenado no endereço referenciado pelo ponteiro.
- O operador `*`(desreferenciamento) permite acessar o valor armazenado no endereço apontado pelo ponteiro.
- Neste caso, será exibido o valor `45`.
  
**Saída do Programa**

Ao executar o código, a saída será semelhante a esta:

```
Valor da variável inteiro: 45
Endereço da variável inteiro: 0xc000014090
Valor da variável ponteiro: 0xc000014090
Valor da variável ponteiro pelo ponteiro: 45
```
- 0xc000014090 é um exemplo de endereço de memória (pode variar a cada execução).