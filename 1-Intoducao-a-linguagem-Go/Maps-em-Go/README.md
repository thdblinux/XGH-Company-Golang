# Explicação do Código em Go

Este documento descreve e explica o código Go fornecido, destacando os conceitos utilizados e as funcionalidades implementadas.

## Estrutura do Código

### Definição da Estrutura `Pessoa`
A estrutura `Pessoa` representa uma pessoa com três atributos:
- `nome`: o nome da pessoa (tipo `string`).
- `idade`: a idade da pessoa (tipo `int`).
- `salario`: o salário da pessoa (tipo `int`).

```go
type Pessoa struct {
    nome    string
    idade   int
    salario int
}
```

### Função Principal: `main`
A função `main` é o ponto de entrada do programa. Nela, ocorre a manipulação de um mapa para armazenar e acessar salários de funcionários.

#### Uso de Maps (`map`)
Um mapa `salFunc` é criado para associar os nomes dos funcionários aos seus respectivos salários:

```go
salFunc := make(map[string]int)
salFunc["Thadeu"] = 10
salFunc["Marcio"] = 20
```

#### Verificação de Existência no Mapa
A função verifica se uma chave específica existe no mapa:

```go
sal, exist := salFunc["Marcioabc"]
fmt.Println(sal, exist)
```
Neste caso, como a chave `"Marcioabc"` não está presente, o valor de `sal` será `0` e `exist` será `false`.

#### Contagem Total de Funcionários
O número total de entradas no mapa é calculado usando `len`:

```go
totalSal := len(salFunc)
fmt.Println("Total de funcionários:", totalSal)
```

### Métodos e Funções Auxiliares

#### Método `addSalaryPessoa`
Este método é associado à estrutura `Pessoa` e adiciona um bônus ao salário de uma pessoa:

```go
func (p *Pessoa) addSalaryPessoa(bonus int) {
    p.salario += bonus
}
```

#### Função `setName`
Imprime uma saudação com o nome fornecido:

```go
func setName(name string) {
    fmt.Println("Hello", name)
}
```

#### Função `addSalary`
Calcula o novo salário somando um valor de bônus e retorna o salário atualizado junto com o bônus:

```go
func addSalary(valorAtual int, bonus int) (int, int) {
    return valorAtual + bonus, bonus
}
```

#### Função `getName`
Retorna o nome "Thadeu" como uma string:

```go
func getName() string {
    return "Thadeu"
}
```

## Conceitos Demonstrados
- **Estruturas (`struct`)**: Representação de dados com múltiplos atributos.
- **Mapas (`map`)**: Estrutura de dados chave-valor para armazenar e acessar informações de forma eficiente.
- **Métodos**: Funções associadas a tipos definidos pelo usuário.
- **Funções**: Estruturas modulares e reutilizáveis de código.

## Execução do Código
Ao executar este programa, o seguinte será impresso no console:
1. O valor e a existência da chave `"Marcioabc"` no mapa.
2. O número total de funcionários registrados no mapa.

## Possíveis Extensões
- Adicionar funcionalidades para manipulação mais complexa do mapa.
- Implementar validações ou interações dinâmicas com o usuário.