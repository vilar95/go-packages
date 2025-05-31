package main

// O contexto no http é utilizado para gerenciar o contexto de uma requisição HTTP, como tempo limite e cancelamento.
// O pacote "context" é usado para criar e gerenciar contextos que podem ser passados entre funções e goroutines.
// Ele permite definir prazos, cancelar operações e compartilhar informações entre diferentes partes do código.
// O contexto é especialmente útil em operações assíncronas, como requisições HTTP, onde você pode querer cancelar uma operação se ela demorar muito ou se não for mais necessária.
import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	// Cria um contexto com um prazo de 3 segundos
	// a regra está sendo passada para a função bookHotel
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	// Simula o processo de reserva de hotel
	select {
		// Se o contexto for cancelado ou o prazo for atingido, a reserva é cancelada
	case <-ctx.Done():
		fmt.Println("Reserva de hotel cancelada: Timeout reached")
		return
	// Se o prazo não for atingido, a reserva é concluída com sucesso
	case <-time.After(1 * time.Second):
		fmt.Println("Reserva de hotel concluída com sucesso")
		return
	}
}
