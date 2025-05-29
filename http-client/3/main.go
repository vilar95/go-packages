package main

import (
	"io"
	"net/http"
)

func main() {
	c := http.Client{}
	// http.NewRequest é utilizado para criar uma nova requisição HTTP.
	req, err := http.NewRequest("POST", "https://google.com", nil)
	if err != nil {
		panic(err)
	}
	
	req.Header.Set("Accept", "application/json")
	// c.Do(req) É utilizado para enviar a requisição HTTP que foi criada com http.NewRequest.
	// Ele executa a requisição e retorna a resposta do servidor.
	// O método Do é usado para enviar a requisição e receber a resposta.
	resp, err := c.Do(req)
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