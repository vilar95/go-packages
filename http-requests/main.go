package main

// Manipulando arquivos em Go
// Exemplo de requisições HTTP

import (
	"fmt"      // O package "fmt" é utilizado para formatação de strings e impressão no console
	"io"       // O package "io" é utilizado para ler e escrever dados em streams
	"net/http" // O package "net/http" é utilizado para realizar requisições HTTP
)

func main() {
	// Realizando uma requisição HTTP GET para o Google
	req, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}

	// defer é utilizado para garantir que a função Close seja chamada no final da função main
	defer req.Body.Close()

	// io.ReadAll é utilizado para ler o corpo da resposta HTTP
	// Ele lê todos os dados do corpo da requisição e retorna um slice de bytes
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Status Code:", req.StatusCode)
	fmt.Println("Response Body:" + string(res))

	// Fechando o corpo da requisição para liberar recursos
	// É importante fechar o corpo da requisição após a leitura para evitar vazamentos de memória
}
