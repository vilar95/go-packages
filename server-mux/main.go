package main

import "net/http"

func main() {
	// Cria um novo servidor HTTP e define o manipulador para a rota "/"
	// Neste exemplo, vamos criar um servidor HTTP simples que responde a requisições
	// mux é um multiplexer de roteamento que permite definir diferentes manipuladores para diferentes rotas
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "Blog Page - Bem-vindo ao meu blog!"})
	http.ListenAndServe(":8080", mux)
	
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Home Page!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/blog" {
		// Se não for, retorna um erro 404
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "404 - Not Found | O recurso solicitado não foi encontrado.", http.StatusNotFound)
		return
	}
	w.Write([]byte(b.title))
}