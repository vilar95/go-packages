package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	// O método Post é utilizado para enviar uma requisição HTTP POST.
	// Ele é usado para enviar dados para o servidor, geralmente em formato JSON ou formulário.
	// O primeiro parâmetro é a URL para onde a requisição será enviada.
	// O segundo parâmetro é o tipo de conteúdo que está sendo enviado (por exemplo, "application/json").
	// O terceiro parâmetro é o corpo da requisição, que pode ser um io.Reader contendo os dados a serem enviados.
	jsonData := bytes.NewBuffer([]byte(`{"name":"eduardo"}`))
	resp, err := c.Post("https://google.com", "application/json", jsonData)
	if err != nil {
		panic(err)
	}
	
	defer resp.Body.Close()

	// io.CopyBuffer é utilizado para copiar dados de um io.Reader para um io.Writer.
	// Ele lê os dados do corpo da resposta (resp.Body) e os escreve na saída padrão (os.Stdout).
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
