package main

// Manipulando HTTP em Go
// Exemplo de requisições HTTP para buscar CEPs utilizando a API ViaCEP
// Neste exemplo, vamos criar um servidor HTTP simples que responde a requisições
// com informações de CEPs utilizando a API ViaCEP.

import (
	"encoding/json"
	"io"

	"net/http"
)

// Exemplo de requisições HTTP para buscar CEPs utilizando a API ViaCEP https://viacep.com.br/
// Convert dados JSON em uma struct https://mholt.github.io/json-to-go/
type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/", ZipCodeSearchHandler)
	http.ListenAndServe(":8080", nil)
}

func ZipCodeSearchHandler(w http.ResponseWriter, r *http.Request) {
	// Verifica se o caminho da URL é "/"
	if r.URL.Path != "/" {
		// Se não for, retorna um erro 404
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "404 - Not Found | O recurso solicitado não foi encontrado.", http.StatusNotFound)
		return
	}
	// Obtém o parâmetro "cep" da URL
	zipCodeParam := r.URL.Query().Get("cep")
	if zipCodeParam == "" {
		// Se o parâmetro "cep" não for fornecido, retorna um erro 400
		http.Error(w, "400 - Bad Request  | O parâmetro CEP é obrigatório.", http.StatusBadRequest)
		return
	}
	// Chama a função para obter os dados do CEP com o parâmetro fornecido
	zipCode, err := getZipCodeData(zipCodeParam)
	if err != nil {
		// Se ocorrer um erro ao obter os dados do CEP, retorna um erro 500
		http.Error(w, "500 - Internal Server Error | Falha ao obter dados do ViaCEP.", http.StatusInternalServerError)
		return
	}
	// Se tudo estiver correto, define o cabeçalho de resposta como JSON e retorna os dados do CEP
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Utiliza json.NewEncoder para codificar a estrutura ViaCEP em JSON e escrever na resposta
	json.NewEncoder(w).Encode(zipCode)
}

func getZipCodeData(zipCode string) (*ViaCEP, error) {
	// Implementar a lógica para buscar os dados do CEP na API ViaCEP
	// A função deve fazer uma requisição HTTP para a API ViaCEP com o CEP fornecido
	res, err := http.Get("https://viacep.com.br/ws/" + zipCode + "/json/")
	if err != nil {
		return nil, err // Retorna o erro se a requisição falhar
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err // Retorna o erro se a leitura do corpo falhar
	}
	var viaCEP ViaCEP
	// Deserializa o JSON recebido na estrutura ViaCEP
	err = json.Unmarshal(body, &viaCEP)
	if err != nil {
		return nil, err // Retorna o erro se a deserialização falhar
	}
	// e retornar uma instância de ViaCEP ou um erro, se ocorrer.
	return &viaCEP, nil
}
