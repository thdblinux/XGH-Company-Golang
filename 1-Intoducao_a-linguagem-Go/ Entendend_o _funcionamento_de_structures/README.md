# README

Este documento explica o funcionamento do código em Go fornecido, que demonstra o uso de structs, ponteiros e funções.

---

## Visão Geral

O código cria uma struct `Pessoa` para representar uma pessoa e inclui funções para manipular seus atributos. O exemplo principal foca em adicionar um bônus ao salário de uma pessoa usando uma função que opera com ponteiros.

---

## Estrutura do Código

### Importação de Pacotes
```go
import (
    "fmt"
)
```
O pacote `fmt` é utilizado para imprimir saídas no console.

### Declaração da Struct `Pessoa`
```go
type Pessoa struct {
    nome    string
    idade   int
    salario int
}
```
A struct `Pessoa` contém:
- `nome`: Nome da pessoa (string).
- `idade`: Idade da pessoa (int).
- `salario`: Salário da pessoa (int).

### Função Principal `main`
```go
func main() {
    pessoa := &Pessoa{
        nome:    "Thadeu",
        idade:   40,
        salario: 100,
    }
    addSalaryPessoa(pessoa, 10)
    fmt.Println(pessoa.salario)
}
```
**Passos executados:**
1. Uma instância da struct `Pessoa` é criada com valores iniciais.
2. A função `addSalaryPessoa` é chamada para adicionar um bônus ao salário da pessoa.
3. O salário atualizado é impresso no console.

### Função `addSalaryPessoa`
```go
func addSalaryPessoa(p *Pessoa, bonus int) {
    p.salario += bonus
}
```
Recebe:
- `p`: Um ponteiro para a struct `Pessoa`.
- `bonus`: Valor do bônus a ser adicionado.

Essa função atualiza diretamente o campo `salario` da struct referenciada pelo ponteiro `p`.

### Funções Auxiliares

#### `setName`
```go
func setName(name string) {
    fmt.Println("Hello", name)
}
```
Imprime uma saudação no console com o nome fornecido. **Não é utilizada no código atual.**

#### `addSalary`
```go
func addSalary(valorAtual int, bonus int) (int, int) {
    return valorAtual + bonus, bonus
}
```
Retorna:
1. O novo salário.
2. O bônus.

Essa função é uma alternativa à `addSalaryPessoa`, mas não é usada no programa principal.

#### `getName`
```go
func getName() string {
    return "Thadeu"
}
```
Retorna uma string com o nome "Thadeu". **Não é utilizada no código atual.**

---

## Saída do Programa
Ao executar o programa, a saída será:
```
110
```
Isso ocorre porque o salário inicial é 100 e um bônus de 10 é adicionado.

---

## Como Executar
1. Certifique-se de ter o Go instalado.
2. Salve o código em um arquivo chamado `main.go`.
3. No terminal, execute:
   ```bash
   go run main.go
   ```

---

## Melhorias Sugeridas
- **Uso de Funções Auxiliares:** Integre ou remova funções não utilizadas.
- **Validação de Dados:** Adicione verificações para evitar valores negativos no bônus ou salário.
- **Teste Unitário:** Crie testes para validar o comportamento das funções.
