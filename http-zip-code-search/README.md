# 🔎 Busca de CEP com Go e ViaCEP

Este projeto demonstra como criar um servidor HTTP simples em Go para buscar informações de CEPs utilizando a API ViaCEP.

## 🚀 Como funciona

- O servidor escuta na porta `8080`.
- Recebe requisições HTTP GET na rota `/` com o parâmetro `cep`.
- Faz uma requisição para a API [ViaCEP](https://viacep.com.br/) e retorna os dados do CEP em formato JSON.

## 📦 Como executar

1. Clone o repositório ou copie o arquivo `main.go`.
2. Execute o servidor:

```bash
# No terminal, dentro da pasta do projeto
 go run main.go
```

3. Faça uma requisição para buscar um CEP (exemplo usando `curl`):

```bash
curl "http://localhost:8080/?cep=01304000"
```

## 📝 Exemplo de resposta

```json
{
"cep": "01304-000",
"logradouro": "Rua Augusta",
"complemento": "até 698 - lado par",
"unidade": "",
"bairro": "Consolação",
"localidade": "São Paulo",
"uf": "SP",
"estado": "São Paulo",
"regiao": "Sudeste",
"ibge": "3550308",
"gia": "1004",
"ddd": "11",
"siafi": "7107"
}
```

## ⚠️ Observações

- O parâmetro `cep` é obrigatório.
- Se o CEP não for informado ou for inválido, o servidor retorna um erro apropriado.

## 📚 Referências

- [Documentação da API ViaCEP](https://viacep.com.br/)
- [Documentação net/http (Go)](https://pkg.go.dev/net/http)

---

> 💡 Dica: Descobrir o CEP certo pode ser o primeiro passo para encontrar o destino desejado!
