# Descrição do Código

Este código foi escrito em Go e demonstra conceitos básicos da linguagem, como:

- Declaração de variáveis.
- Funções com retorno simples e múltiplo.
- Passagem de parâmetros para funções.
- Uso de operadores aritméticos.

O programa simula uma operação simples em que:
- Um nome é exibido na tela.
- Um salário inicial recebe um bônus, e o novo salário e o valor do bônus são retornados.

---

## Estrutura do Código

### Importação de Pacotes
```go
import (
    "fmt"
)
```

O pacote `fmt` é usado para imprimir mensagens no console.

### Função `main`
```go
func main() {
    name, salario := "Thadeu", 100
    setName(name)
    newSalary, bonus := addSalary(salario, 10)
    fmt.Println("Novo salário:", newSalary)
    fmt.Println("Novo bonus:", bonus)
}
```
A função `main` é o ponto de entrada do programa. Nela:
1. O nome `Thadeu` e o salário inicial `100` são declarados.
2. A função `setName` exibe uma saudação no console usando o nome fornecido.
3. A função `addSalary` calcula o novo salário e o valor do bônus, que são exibidos no console.

### Função `setName`
```go
func setName(name string) {
    fmt.Println("Hello", name)
}
```
Recebe uma string como parâmetro (`name`) e imprime uma saudação.

### Função `addSalary`
```go
func addSalary(valorAtual int, bonus int) (int, int) {
    return valorAtual + bonus, bonus
}
```
Recebe dois parâmetros inteiros:
1. `valorAtual`: o valor do salário atual.
2. `bonus`: o valor do bônus a ser adicionado ao salário.

Retorna dois valores inteiros:
1. O novo salário, calculado como `valorAtual + bonus`.
2. O valor do bônus.

### Função `getName`
```go
func getName() string {
    return "Thadeu"
}
```
Retorna uma string com o nome "Thadeu". No entanto, esta função não é utilizada no programa atual.

---

## Saída do Programa
Ao executar o código, a saída será:
```
Hello Thadeu
Novo salário: 110
Novo bonus: 10
```

---

## Instruções para Execução
1. Certifique-se de ter o Go instalado na sua máquina.
2. Salve o código em um arquivo chamado `main.go`.
3. Execute os seguintes comandos no terminal:
   ```bash
   go run main.go
   ```

---

## Explicação Adicional
### Pontos Importantes
- **Passagem de parâmetros:** O nome é passado como argumento para a função `setName`. Já o salário e o bônus são usados na função `addSalary` para calcular novos valores.
- **Retornos múltiplos:** A função `addSalary` demonstra como retornar múltiplos valores em Go.
- **Declaração curta de variáveis:** A sintaxe `:=` é usada para declarar e inicializar variáveis simultaneamente.

### Sugestões de Melhoria
- Utilizar a função `getName` no código para tornar o programa mais coeso.
- Adicionar validações para evitar valores negativos para salário ou bônus.
