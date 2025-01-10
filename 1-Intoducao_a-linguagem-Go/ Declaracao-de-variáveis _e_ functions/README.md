# Entendendo o funcionamento do Import

Este repositório contém um exemplo simples de programa escrito em Go, desenvolvido como parte de um treinamento para DevOps e SRE. O objetivo deste programa é demonstrar conceitos básicos de manipulação de argumentos de linha de comando, lógica de controle e interação com o sistema operacional.

## Descrição do Código

Este programa Go tem a função principal de receber um argumento de linha de comando e imprimir uma mensagem personalizada no terminal. Ele também implementa validação para garantir que o número correto de argumentos seja fornecido.

### Código-Fonte

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args
	// ARGS[0,1]
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("Hello", os.Args[1])
}
```

### Explicação do Código

1. **Pacote Principal (`package main`)**
   - Todo programa Go inicia com o pacote `main`, que indica que este é o ponto de entrada do programa.

2. **Importação de Pacotes (`import`)**
   - `fmt`: Fornece funções para formatação e impressão de dados no terminal.
   - `os`: Permite interagir com o sistema operacional, como acessar argumentos de linha de comando e definir códigos de saída.

3. **Função `main`**
   - A função `main` é o ponto de entrada do programa.

4. **Uso de `os.Args`**
   - `os.Args` é uma fatia (`slice`) que contém os argumentos de linha de comando passados para o programa.
   - `os.Args[0]`: Contém o nome do programa.
   - `os.Args[1]`: Contém o primeiro argumento fornecido pelo usuário.

5. **Validação do Número de Argumentos**
   - O programa verifica se o número de argumentos fornecido é exatamente 2 (o nome do programa + 1 argumento).
   - Caso contrário, ele encerra o programa com um código de saída `1`.

6. **Impressão da Mensagem**
   - Se a validação passar, o programa imprime a mensagem "Hello" seguida do argumento fornecido pelo usuário.

## Como Executar

1. Clone este repositório:
   ```bash
   git clone <URL_DO_REPOSITORIO>
   cd <NOME_DO_DIRETORIO>
   ```

2. Compile o programa:
   ```bash
   go build -o hello
   ```

3. Execute o programa passando um argumento de linha de comando:
   ```bash
   ./hello [SEU_NOME]
   ```
   Exemplo:
   ```bash
   ./hello Thadeu
   ```
   Saída esperada:
   ```
   Hello Thadeu
   ```

4. Teste o comportamento sem argumentos ou com mais de um argumento:
   ```bash
   ./hello
   ```
   Código de saída esperado:
   ```bash
   echo $?
   1
   ```

## Objetivo do Treinamento

- Demonstrar conceitos básicos de programação em Go.
- Aprender a utilizar argumentos de linha de comando.
- Implementar validação de entrada e lógica de controle.
- Integrar aplicações Go com práticas de DevOps e SRE.

## Próximos Passos

- Expandir o programa para lidar com múltiplos argumentos.
- Adicionar testes automatizados para validar o comportamento do programa.
- Integrar o programa em pipelines de CI/CD para automação de build e deploy.