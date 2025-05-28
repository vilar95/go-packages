package main

// Manipulando arquivos em Go

import (
	"encoding/json" // O package "encoding/json" é utilizado para codificar e decodificar dados JSON
	"fmt" // O package "fmt" é utilizado para formatação de strings e impressão no console
	"io" // O package "io" é utilizado para ler e escrever dados em streams
	"net/http" // O package "net/http" é utilizado para realizar requisições HTTP
	"os" // O package "os" é utilizado para manipulação de arquivos e variáveis de ambiente
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
	// Verifica se foram passados argumentos na linha de comando
	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}

		defer req.Body.Close()

		// Lê o corpo da resposta HTTP
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}

		var data ViaCEP
		// Deserializa o JSON recebido na estrutura ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao deserializar JSON: %v\n", err)
		}
		// Cria um arquivo para salvar os dados do CEP
		file, err := os.Create(fmt.Sprintf("address_%s.txt", data.Cep))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}

		defer file.Close()
		// Formata os dados do CEP e escreve no arquivo
		address := fmt.Sprintf(
			"CEP: %s\nLogradouro: %s\nComplemento: %s\nUnidade: %s\nBairro: %s\nLocalidade: %s\nUF: %s\nEstado: %s\nRegião: %s\nIBGE: %s\nGIA: %s\nDDD: %s\nSIAFI: %s\n",
			data.Cep, data.Logradouro, data.Complemento, data.Unidade, data.Bairro, data.Localidade, data.Uf, data.Estado, data.Regiao, data.Ibge, data.Gia, data.Ddd, data.Siafi,
		)
		// Escreve os dados formatados no arquivo
		if _, err := file.WriteString(address); err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v\n", err)
		}
		// Exibe uma mensagem de sucesso
		fmt.Printf("Dados do CEP %s salvos com sucesso em address_%s.txt\n", data.Cep, data.Cep)

		defer file.Close()
	}
}
