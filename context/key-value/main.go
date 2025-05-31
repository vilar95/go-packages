package main

import "context"

func main() {
	ctx := context.Background()

	// Cria um novo contexto com um valor associado à chave
	ctx = context.WithValue(ctx, "token", "6XB3-4F2D-9C1A-8E7F-2B7H")

	// Recupera o valor associado à chave
	bookHotel(ctx, "Hotel White Lotus")
}

func bookHotel(ctx context.Context, _ string) {
	// Simula o processo de reserva de hotel
	token := ctx.Value("token")
	if token == nil {
		println("Reserva de hotel falhou: Token não encontrado")
		return
	}
	println("Reserva de hotel concluída com sucesso usando o token:", token.(string))
}
