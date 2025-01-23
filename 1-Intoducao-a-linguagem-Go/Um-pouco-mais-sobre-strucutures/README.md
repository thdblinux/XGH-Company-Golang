# README

Este documento explica o funcionamento do código em Go fornecido, que demonstra o uso de structs, ponteiros, métodos associados e funções.

---

## Visão Geral

O código cria e manipula instâncias de uma struct `Pessoa`, representando uma pessoa, e inclui um método para ajustar o salário da pessoa. Ele também exemplifica outras funções auxiliares que não são utilizadas diretamente no exemplo principal.

---

## Estrutura do Código

### Importação de Pacotes
```go
import (
    "fmt"
)
```
O pacote `fmt` é usado para imprimir saídas no console.

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
    pessoa2 := new(Pessoa)
    pessoa2.nome = "Marcio"
    pessoa2.idade = 40
    pessoa2.salario = 100

    pessoa := &Pessoa{
        nome:    "Thadeu",
        idade:   40,
        salario: 100,
    }
    pessoa.addSalaryPessoa(10)
    fmt.Println(pessoa.salario)
}
```
**Passos executados:**
1. Uma instância `pessoa2` é criada usando a função `new` e seus campos são inicializados manualmente.
2. Outra instância `pessoa` é criada usando inicialização direta com valores predefinidos.
3. O método `addSalaryPessoa` é chamado na instância `pessoa` para adicionar um bônus de 10 ao salário.
4. O salário atualizado é impresso no console.

### Método `addSalaryPessoa`
```go
func (p *Pessoa) addSalaryPessoa(bonus int) {
    p.salario += bonus
}
```
**Recebe:**
- `p`: Um ponteiro para a struct `Pessoa`.
- `bonus`: Valor do bônus a ser adicionado.

Este método atualiza diretamente o campo `salario` da struct referenciada pelo ponteiro `p`.

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

Essa função é uma alternativa ao método `addSalaryPessoa`, mas não é usada no programa principal.

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
Isso ocorre porque o salário inicial de 100 recebe um bônus de 10.

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
- **Teste Unitário:** Crie testes para validar o comportamento das funções e métodos.
