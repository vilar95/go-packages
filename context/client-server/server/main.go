package main

import (
	"log"
	"net/http"
	"time"
)

// O contexto no http é utilizado para gerenciar o contexto de uma requisição HTTP, como tempo limite e cancelamento.
// Isso evita que uma requisição HTTP fique pendurada indefinidamente, permitindo que o servidor responda rapidamente a cancelamentos ou prazos expirados.
// O contexto é especialmente útil em operações assíncronas, como requisições HTTP, onde você pode querer cancelar uma operação se ela demorar muito ou se não for mais necessária.

func main() {
	http.HandleFunc("/", hanler)
	http.ListenAndServe(":8080", nil)
}

func hanler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")

	defer log.Println("Request finalizada")
	select {
		case <-time.After(5 * time.Second):
			log.Println("Request processada com sucesso")
			w.Write([]byte("Request processada com sucesso"))
		case <-ctx.Done():
			log.Println("Request cancelada pelo cliente")
		}
	
}