# Tratamento de Erros e defer

Este projeto demonstra como utilizar tratamento de `erros` e a instrução `defer` em Go para gerenciar recursos de forma eficiente e evitar `vazamentos`, com exemplos simples e funcionais.

## Conceitos Abordados

1. **Tratamento de Erros**
   - Utilização do pacote `os` para abrir `arquivos`.
   - Tratamento de `erros` retornados por funções que seguem o padrão `Go` de retornar `(resultado, erro)`.
   - Como validar `erros` e realizar uma saída controlada em caso de falhas.

2. **Instrução `defer`**
   - Uso de `defer` para garantir o `fechamento de recursos`, como `arquivos`, mesmo em casos de `erro`.

---

## Código do Projeto

### Estrutura do Código
O arquivo principal, `main.go`, contém o seguinte fluxo:

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Tentativa de abrir um arquivo
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Outros exemplos de funcionalidade (comentados no código original)
}
```
 ## Passo a Passo do Código

Abrindo um Arquivo com `os.Open`:

```
file, err := os.Open("file.txt")
if err != nil {
    fmt.Println(err)
    return
}
```
- A função os.Open tenta abrir o arquivo `file.txt`.
- Caso o arquivo não exista ou ocorra algum erro, o programa imprime a mensagem de erro e encerra.


**Utilizando defer para Fechar o Arquivo:**
```
defer file.Close()
```
A instrução defer garante que o arquivo será fechado ao final da execução da função main.
Isso evita vazamentos de recursos, mesmo que o programa termine prematuramente por erros.

**Exemplos Adicionais (Comentados):**

**O código inclui exemplos comentados que podem ser explorados, como:**
- Validação de argumentos da linha de comando.
- Conversão de strings para inteiros usando strconv.Atoi.
- Criação e manipulação de mapas e structs.
- Arquivo Necessário
- Certifique-se de criar um arquivo chamado file.txt na mesma pasta que o arquivo main.go. O conteúdo do arquivo pode ser qualquer texto simples.