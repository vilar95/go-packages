package main

// Manipulando arquivos em Go
// Exemplo de usos com JSON
// Neste exemplo, vamos trabalhar com JSON em Go, incluindo leitura.

import (
	"encoding/json" // O package "encoding/json" é utilizado para codificar e decodificar dados JSON
	"fmt"
	"os"
)

type Account struct {
	// tags são utilizados para especificar como os campos da struct serão codificados em JSON
	Number  string		`json:"n"` // O campo "n" será codificado como "number" no JSON 
	Balance float64		`json:"b"` // O campo "b" será codificado como "balance" no JSON
}

func main() {
	// Criando uma instância de account
	account := Account{
		Number:  "123456",
		Balance: 1000.50,
	}
	// Convertendo a struct account para JSON
	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON:", string(res))

	// O json.Encoder é utilizado para codificar dados JSON e escrever em um Write
	// o os.Stdout é um Writer que representa a saída padrão do programa
	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		panic(err)
	}

	// Escrevendo um slice de bytes JSON
	jsonPure := []byte(`{"Number":"123456","Balance":1810.50}`)
	
	// Criando uma variável do tipo Account para armazenar os dados decodificados
	var account2 Account

	// Convertendo JSON para struct account2
	// O json.Unmarshal é utilizado para decodificar dados JSON em uma struct
	err = json.Unmarshal(jsonPure, &account2)
	if err != nil {
		panic(err)
	}
	fmt.Println("Account2:", account2)

	jsonPure = []byte(`{"n":"123456","b":1810.50}`)
	err = json.Unmarshal(jsonPure, &account2)
	if err != nil {
		panic(err)
	}
	fmt.Println("Account2 with tags:", account2)

}
