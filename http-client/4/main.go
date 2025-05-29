package main

// Pacote de context é utilizado para gerenciar o contexto de uma requisição HTTP, como tempo limite e cancelamento.

import (
	"context"
	"io"
	"net/http"
	"time"
)

func main() {
	// O context.Background() cria um contexto vazio que pode ser usado como base para outros contextos.
	ctx := context.Background()
	// context.WithTimeout(ctx, time.Second) cria um novo contexto com um tempo limite de 1 segundo.
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	// É importante chamar cancel() para liberar os recursos associados ao contexto quando não for mais necessário.
	defer cancel()

	// http.NewRequestWithContext(ctx, "GET", "https://google.com", nil) cria uma nova requisição HTTP com o contexto especificado.
	// O primeiro parâmetro é o contexto, que pode ser usado para definir um tempo limite ou cancelar a requisição.
	req, err := http.NewRequestWithContext(ctx, "GET", "https://google.com", nil)
	if err != nil {
		panic(err)
	}
	// O método Do é utilizado para enviar a requisição HTTP que foi criada com http.NewRequestWithContext.
	// Ele executa a requisição e retorna a resposta do servidor.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}