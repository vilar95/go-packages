package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	monitoringCycles = 2
	delaySeconds     = 5
	sitesFile        = "sites.txt"
	logFile          = "log.txt"
	httpTimeout      = 10 * time.Second
)

func main() {
	displayIntroduction()

	for {
		displayMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo Logs...")
			displayLogs()
		case 0:
			fmt.Println("Saindo do programa!")
			os.Exit(0)
		default:
			fmt.Println("Comando desconhecido.")
			os.Exit(-1)
		}
	}
}

func displayIntroduction() {
	name := "Eduardo"
	version := 2.0
	fmt.Println("Olá,", name)
	fmt.Println("Este programa está na versão", version)
}

func displayMenu() {
	fmt.Println("\n=== Sistema de Monitoramento de Sites ===")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
	fmt.Print("Escolha uma opção: ")
}

func readCommand() int {
	var command int
	_, err := fmt.Scan(&command)
	if err != nil {
		fmt.Println("Entrada inválida. Por favor, digite um número.")
		return -1
	}
	fmt.Println("Comando escolhido:", command)
	return command
}

func startMonitoring() {
	fmt.Println("\nIniciando monitoramento...")
	sites := readSitesFromFile()

	if len(sites) == 0 {
		fmt.Println("Nenhum site para monitorar. Por favor, adicione sites ao arquivo", sitesFile)
		return
	}

	for cycle := 1; cycle <= monitoringCycles; cycle++ {
		fmt.Printf("\n--- Ciclo de Monitoramento %d/%d ---\n", cycle, monitoringCycles)
		for index, site := range sites {
			fmt.Printf("Testando site %d: %s\n", index+1, site)
			testSite(site)
		}
		if cycle < monitoringCycles {
			fmt.Printf("\nAguardando %d segundos antes do próximo ciclo...\n", delaySeconds)
			time.Sleep(delaySeconds * time.Second)
		}
	}

	fmt.Println("\nMonitoramento concluído!")
}

func testSite(site string) {
	client := &http.Client{
		Timeout: httpTimeout,
	}

	resp, err := client.Get(site)
	if err != nil {
		fmt.Printf("Erro ao acessar o site %s: %v\n", site, err)
		logStatus(site, false, 0, err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("✓ Site %s carregado com sucesso! (Status: %d)\n", site, resp.StatusCode)
		logStatus(site, true, resp.StatusCode, "")
	} else {
		fmt.Printf("✗ Site %s está com problemas. Status Code: %d\n", site, resp.StatusCode)
		logStatus(site, false, resp.StatusCode, "")
	}
}

func readSitesFromFile() []string {
	var sites []string

	file, err := os.Open(sitesFile)
	if err != nil {
		log.Printf("Erro ao abrir o arquivo %s: %v\n", sitesFile, err)
		return sites
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line != "" && !strings.HasPrefix(line, "#") {
			sites = append(sites, line)
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Erro ao ler o arquivo: %v\n", err)
			break
		}
	}

	return sites
}

func logStatus(site string, online bool, statusCode int, errorMsg string) {
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Erro ao abrir o arquivo de log: %v\n", err)
		return
	}
	defer file.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("%s | %s | online: %s | status: %d",
		timestamp, site, strconv.FormatBool(online), statusCode)

	if errorMsg != "" {
		logEntry += fmt.Sprintf(" | error: %s", errorMsg)
	}
	logEntry += "\n"

	if _, err := file.WriteString(logEntry); err != nil {
		log.Printf("Erro ao escrever no arquivo de log: %v\n", err)
	}
}

func displayLogs() {
	content, err := os.ReadFile(logFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Nenhum log encontrado. Inicie o monitoramento para gerar logs.")
			return
		}
		log.Printf("Erro ao ler o arquivo de log: %v\n", err)
		return
	}

	if len(content) == 0 {
		fmt.Println("Arquivo de log está vazio.")
		return
	}

	fmt.Println("\n=== Logs de Monitoramento ===")
	fmt.Println(string(content))
}
