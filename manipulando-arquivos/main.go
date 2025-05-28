package main

// Manipulando arquivos em Go
// Exemplo simples de leitura e escrita de arquivos

import (
	"bufio" // O package "bufio" é utilizado para leitura eficiente de arquivos
	"fmt"   // O package "fmt" é utilizado para formatação de strings e impressão no console
	"os"    // O package "os" é utilizado para manipulação de arquivos e diretórios
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// Escrevendo no arquivo
	_, err = f.WriteString("Hello, World!\n")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso!\n")

	// Criando um loop de escrita para adicionar várias linhas
	for i := 0; i < 5; i++ {
		// Escrevendo informações adicionais no arquivo
		_, err = f.Write([]byte("Manipulando arquivos em Go.\n"))
		if err != nil {
			panic(err)
		}
	}

	// Exibindo o tamanho do arquivo
	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Tamanho do arquivo: %d bytes\n", fileInfo.Size())

	f.Close()

	// Lendo o arquivo
	file, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	// Exibindo o conteúdo do arquivo
	fmt.Println("\nConteúdo do arquivo:")
	fmt.Println(string(file))

	// Leitura de arquivo linha por linha
	file2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	// NewReader para ler o arquivo
	// Utilizando bufio para ler o arquivo de forma eficiente
	reader := bufio.NewReader(file2)
	// Lendo o arquivo em blocos de 14 bytes
	buffer := make([]byte, 14)
	for {
		// Read lê até o tamanho do buffer ou até que não haja mais dados
		n, err := reader.Read(buffer)
		if err != nil {
			break // Sai do loop se não houver mais dados
		}
		fmt.Println(string(buffer[:n])) // Imprime o conteúdo lido
	}
	// Fechando o arquivo após a leitura
	file2.Close()
	fmt.Println("Leitura do arquivo concluída.")

	// Excluindo o arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Arquivo excluído com sucesso!")
	fmt.Println("Programa finalizado com sucesso!")

}
