# 📦 Projeto Go + MySQL: CRUD de Produtos

Este projeto demonstra como criar, ler, atualizar e deletar produtos em um banco de dados MySQL usando Go. O ambiente é facilmente configurável com Docker. 


> ✨ Dica: Sempre verifique se o Docker Desktop está rodando antes de executar os comandos Docker! 😉

---

## 🚀 Como rodar o projeto

### 1. 🐳 Subindo o MySQL com Docker

O projeto já possui um `docker-compose.yaml` pronto. Para subir o banco de dados:

```powershell
# No PowerShell, dentro da pasta database
docker compose up -d
```

Isso irá:
- ⬇️ Baixar a imagem oficial do MySQL
- 🏗️ Criar um container chamado `mysql_container`
- 🚪 Expor a porta 3306
- 🗄️ Criar o banco `goexpert` com usuário `user` e senha `root`

### 2. 🛠️ Acessando o MySQL

Você pode acessar o MySQL do container usando:

```powershell
# Acesse o terminal do container
docker exec -it mysql_container bash

# Dentro do container, acesse o MySQL:
mysql -uroot -proot goexpert
```

Ou, se tiver o cliente MySQL instalado localmente:

```powershell
mysql -h 127.0.0.1 -P 3306 -u root -p goexpert
# senha: root
```

### 3. 📝 Criando a tabela `products`

Antes de rodar o código Go, crie a tabela:

```sql
CREATE TABLE products (
  id VARCHAR(36) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  price DOUBLE NOT NULL
);
```

### 4. 🏃‍♂️ Rodando o código Go

Instale as dependências (se necessário):

```powershell
go mod tidy
```

Execute o programa:

```powershell
go run main.go
```

O programa irá:
- ➕ Inserir um produto
- 📋 Listar todos os produtos
- ✏️ Atualizar o produto
- 📋 Listar novamente
- 🗑️ Deletar todos os produtos
- 📋 Listar novamente

---

## 🛠️ Comandos úteis

### 🐳 Docker
- Subir o banco: `docker compose up -d`
- Parar: `docker compose down`
- Ver logs: `docker logs -f mysql_container`
- Acessar o container: `docker exec -it mysql_container bash`

### 🐬 MySQL
- Listar bancos: `SHOW DATABASES;`
- Usar banco: `USE goexpert;`
- Listar tabelas: `SHOW TABLES;`
- Ver estrutura: `DESCRIBE products;`
- Consultar: `SELECT * FROM products;`

### 🦫 Go
- Instalar dependências: `go mod tidy`
- Rodar: `go run main.go`

---

## 📚 Referências
- [Documentação oficial do MySQL Docker](https://hub.docker.com/_/mysql)
- [Driver MySQL para Go](https://github.com/go-sql-driver/mysql)
- [Documentação database/sql](https://pkg.go.dev/database/sql)

---


