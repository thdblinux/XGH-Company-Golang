# Projeto em Golang: Gerenciamento de Funcionários

Este é um projeto simples em Golang que demonstra conceitos como organização de código em pacotes, manipulação de mapas, e a criação e utilização de structs e métodos. O objetivo é gerenciar informações de funcionários como nome, idade e salário.

## Estrutura do Projeto

O projeto é composto pelos seguintes arquivos e pastas:

```
.
├── main.go                // Arquivo principal do programa
└── funcionarios/
    └── pessoa.go          // Definição da struct Pessoa e seus métodos
```

## Arquivo `main.go`

O arquivo `main.go` é o ponto de entrada do programa e possui a seguinte funcionalidade:

1. **Importação do pacote `funcionarios`:**
   O pacote `funcionarios` é utilizado para manipular dados relacionados à struct `Pessoa`.

2. **Criação de uma instância de `Pessoa`:**
   ```go
   pessoa := &funcionarios.Pessoa{
       nome:    "Thadeu",
       idade:   40,
       salario: 100,
   }
   ```
   Aqui, é criada uma referência para um objeto do tipo `Pessoa` com valores iniciais para os campos `nome`, `idade` e `salario`.

3. **Trabalhando com mapas:**
   Um mapa chamado `salFunc` é criado para armazenar informações sobre o salário de diferentes funcionários.
   ```go
   salFunc := make(map[string]int)
   salFunc["Marcio"] = 10
   salFunc["Thadeu"] = 20
   ```
   - O mapa associa o nome do funcionário (chave) ao seu salário (valor).
   - O programa verifica se um funcionário existe no mapa e imprime o resultado:
     ```go
     sal, exist := salFunc["Marcio"]
     fmt.Println(sal, exist)
     ```

4. **Funções auxiliares:**
   O arquivo também define algumas funções adicionais, que não são utilizadas diretamente no programa principal:
   - `setName(name string)`: Imprime uma mensagem de saudação com o nome fornecido.
   - `addSalary(valorAtual int, bonus int)`: Calcula um novo salário adicionando um bônus ao valor atual e retorna os dois valores.
   - `getName()`: Retorna uma string fixa com o nome "Thadeu".

## Arquivo `funcionarios/pessoa.go`

Este arquivo contém a definição da struct `Pessoa` e um método associado.

1. **Definição da struct `Pessoa`:**
   ```go
   type Pessoa struct {
       nome    string
       idade   int
       salario int
   }
   ```
   - `nome`: Nome do funcionário.
   - `idade`: Idade do funcionário.
   - `salario`: Salário atual do funcionário.

2. **Método `addSalaryPessoa`:**
   Este método permite adicionar um bônus ao salário do funcionário.
   ```go
   func (p *Pessoa) addSalaryPessoa(bonus int) {
       p.salario += bonus
   }
   ```
   - Ele recebe um ponteiro para a instância de `Pessoa` e o valor do bônus a ser adicionado.

## Como Executar o Programa
   ```
1. Execute o programa:
   ```bash
   go run main.go
   ```

## Saída Esperada

Ao executar o programa, a saída será semelhante a:
```
10 true
```

## Conclusão

Este projeto demonstra como organizar um código em pacotes, trabalhar com structs, mapas e métodos em Go. Ele é uma introdução simples mas efetiva para iniciantes que desejam aprender os fundamentos da linguagem.