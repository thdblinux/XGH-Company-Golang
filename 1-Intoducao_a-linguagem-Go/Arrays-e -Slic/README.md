# README

## Introdução

Este código em Go demonstra conceitos importantes como:
- Uso de slices para manipulação de listas dinâmicas.
- Estruturas (`struct`) para modelar dados.
- Métodos associados às structs.
- Uso de laços (`for`) para operações iterativas.

---

## Estrutura do Código

### Importação de Pacotes

```go
import (
    "fmt"
)
```
O pacote `fmt` é usado para entrada e saída, como imprimir mensagens no console.

### Declaração da Struct `Pessoa`

```go
type Pessoa struct {
    nome    string
    idade   int
    salario int
}
```
A struct `Pessoa` define um modelo com os seguintes campos:
- `nome`: Nome da pessoa (string).
- `idade`: Idade da pessoa (int).
- `salario`: Salário da pessoa (int).

### Função Principal (`main`)

#### Inicialização de Slices

```go
salarios := []int{}
```
Um slice vazio é criado para armazenar salários. Slices são listas dinâmicas que permitem adição de elementos de forma eficiente.

#### Preenchendo o Slice

```go
for i := 0; i < 10; i++ {
    salarios = append(salarios, 100+i)
}
```
- Um loop `for` é usado para adicionar 10 salários ao slice, começando de 100.
- A função `append` adiciona elementos ao slice dinamicamente.

#### Exibindo os Salários

```go
for _, salario := range salarios {
    fmt.Println(salario)
}
```
- Um loop `for` com `range` percorre o slice, imprimindo cada salário no console.

### Método Associado à Struct `Pessoa`

#### `addSalaryPessoa`

```go
func (p *Pessoa) addSalaryPessoa(bonus int) {
    p.salario += bonus
}
```
Este método:
1. Recebe um ponteiro para uma instância de `Pessoa`.
2. Adiciona o valor do bônus ao campo `salario` da pessoa.

### Funções Auxiliares (Comentadas no Código)

#### `setName`

```go
func setName(name string) {
    fmt.Println("Hello", name)
}
```
Imprime uma saudação no console com o nome fornecido.

#### `addSalary`

```go
func addSalary(valorAtual int, bonus int) (int, int) {
    return valorAtual + bonus, bonus
}
```
Calcula e retorna:
- O novo salário.
- O valor do bônus.

#### `getName`

```go
func getName() string {
    return "Thadeu"
}
```
Retorna uma string fixa: "Thadeu".

---

## Saída do Programa

Ao executar o código, a saída será:

```
100
101
102
103
104
105
106
107
108
109
```
Isso ocorre porque o loop adiciona 10 valores incrementais ao slice `salarios`, começando de 100, e os imprime no console.

---

## Como Executar o Código

1. Certifique-se de ter o Go instalado em sua máquina.
2. Salve o código em um arquivo chamado `main.go`.
3. No terminal, execute:

```bash
go run main.go
```

---

## Possíveis Melhorias

1. **Validação de Dados:**
   - Garantir que os valores inseridos no slice sejam positivos.
2. **Testes Automatizados:**
   - Criar testes para verificar a lógica de incremento de salários.
3. **Refatoração:**
   - Remover funções não utilizadas, como `setName` e `getName`, se não forem necessárias.
4. **Comentários:**
   - Expandir comentários no código para melhorar a legibilidade.

---

## Conclusão

Este código é uma introdução prática ao uso de slices, structs e métodos no Go, ilustrando como manipular e organizar dados de forma eficiente.
