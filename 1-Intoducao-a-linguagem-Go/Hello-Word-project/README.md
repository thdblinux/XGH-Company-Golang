# Montando Seu Primeiro Projeto

Este repositório apresenta o exemplo clássico "Hello World!" escrito em Go. Ele é o ponto de partida ideal para quem está aprendendo a programar com a linguagem Go, especialmente para iniciantes em DevOps e SRE.

## Descrição do Código

O programa tem como objetivo:

1. Demonstrar a estrutura básica de um programa Go.
2. Exibir a mensagem "Hello World!" no terminal.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

## Explicação do Código

**1.Pacote Principal (package main)**
- Todo programa Go começa com o pacote main, que identifica o ponto de entrada.
  
**2. Importação do Pacote fmt**
- O pacote fmt fornece funcionalidades para exibição de dados no terminal, como a função Println.
  
**3. Função main**
- É a função principal que será executada quando o programa for iniciado.
- Contém a chamada à função fmt.Println para imprimir a mensagem "Hello World!".

# Comandos Go

## 1. Inicializar um Módulo Go
```sh
go mod init gohelloword
```
Cria um arquivo `go.mod`no diretório atual, que define o nome do módulo (neste caso, `gohelloword`) e permite gerenciar dependências no seu projeto Go.

### Exemplo de saída:
```
module gohelloword
go 1.20
```
## 2. Executar um Arquivo Go
```sh
go run main.go
```
Compila e executa o código no arquivo `main.go` diretamente.
Ideal para testar rapidamente seu código sem precisar criar um binário.

## Exemplo de execução:
**Código:**
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
Saída:

```
Hello, World!
```
## 3. Compilar um Arquivo Go
```
go build main.go
```

Compila o arquivo `main.go` e cria um executável na mesma pasta (no caso de sistemas Unix, será `main`, e no Windows, `main.exe`).

Executando o binário:
```
./main
```

**Saída (usando o exemplo anterior):**
```
Hello, World!
```