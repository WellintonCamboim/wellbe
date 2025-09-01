# 🚀 WellBe - API em Go com PostgreSQL

[![Go Version](https://img.shields.io/badge/go-1.22+-00ADD8?logo=go)](https://golang.org/dl/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-336791?logo=postgresql)](https://www.postgresql.org/)
[![Docker](https://img.shields.io/badge/docker-23+-2496ED?logo=docker)](https://www.docker.com/)

API desenvolvida em Go com PostgreSQL e Docker, utilizando migrações SQL.

---

## 📋 Índice

- [Pré-requisitos](#-pré-requisitos)
- [Configuração Inicial](#-configuração-inicial)
- [Execução Recomendada: Docker Compose](#-execução-recomendada-docker-compose)
- [Execução Local (sem Docker)](#-execução-local-sem-docker)
- [Banco de Dados](#-banco-de-dados)
- [Migrações](#-migrações)
- [Estrutura do Projeto](#-estrutura-do-projeto)
- [Solução de Problemas](#-solução-de-problemas)
- [Licença](#-licença)
- [Como usar](#como-usar)
- [Personalização](#personalização)

---

## 🛠 Pré-requisitos

### 💻 Software Necessário

- [Go 1.22+](https://golang.org/dl/)
- [Docker 23+](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Git](https://git-scm.com/)

### 🔍 Verificação

```bash
go version
docker --version
docker-compose --version
```

---

## 🚀 Configuração Inicial

1. **Clone o repositório**
    ```bash
    git clone https://github.com/seu-usuario/wellbe.git
    cd wellbe
    ```
2. **Configure as variáveis de ambiente**
    ```bash
    cp configs/app.env.example .env
    nano .env  # Edite com suas credenciais
    ```
3. **Permissões (Linux/Mac)**
    ```bash
    chmod +x scripts/*.sh  # Se tiver scripts auxiliares
    ```

---

## 🐳 Execução Recomendada: Docker Compose

> **Recomendado:** Use Docker Compose para rodar a aplicação e o banco de dados facilmente, sem precisar instalar o PostgreSQL localmente.

1. **Configure o arquivo `.env`**  
   Use as variáveis já sugeridas:
   ```env
   DB_HOST=db
   DB_PORT=5432
   DB_USER=wellbe
   DB_PASSWORD=wellbe
   DB_NAME=wellbe
   ```

2. **Suba todos os serviços**
    ```bash
    docker-compose up --build -d
    ```

3. **Acompanhe os logs**
    ```bash
    docker-compose logs -f app
    ```

4. **Acesse a aplicação**
    - [http://localhost:8080](http://localhost:8080)

5. **Acesse o banco de dados (opcional)**
    ```bash
    docker exec -it wellbe_db psql -U wellbe -d wellbe
    ```

---

## 💻 Execução Local (sem Docker)

> ⚠️ **Atenção:** Só siga esta seção se realmente quiser rodar tudo fora de containers. Você precisará instalar e configurar o PostgreSQL manualmente.

1. **Instale e inicie o PostgreSQL localmente**  
   Certifique-se de que o PostgreSQL está instalado e rodando.  
   Crie o banco e usuário conforme seu `.env`:

   ```bash
   sudo apt update
   sudo apt install postgresql postgresql-client
   sudo -u postgres createdb wellbe
   sudo -u postgres createuser wellbe --pwprompt
   # Dê permissões ao usuário:
   sudo -u postgres psql
   ALTER USER wellbe WITH PASSWORD 'wellbe';
   GRANT ALL PRIVILEGES ON DATABASE wellbe TO wellbe;
   \q
   ```

2. **Configure o arquivo `.env`**  
   Certifique-se de que as variáveis de ambiente estão corretas para o ambiente local:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=wellbe
   DB_PASSWORD=wellbe
   DB_NAME=wellbe
   ```

3. **Instale dependências**
    ```bash
    go mod download
    ```

4. **Execute a aplicação**
    ```bash
    go run ./cmd/api
    ```
   Acesse: [http://localhost:8080](http://localhost:8080)

---

## 🗃 Banco de Dados

### 🔌 Conexão manual (Docker)

```bash
docker exec -it wellbe_db psql -U wellbe -d wellbe
```

### ⚙️ Variáveis de conexão (.env)

```env
DB_HOST=db
DB_PORT=5432
DB_USER=wellbe
DB_PASSWORD=wellbe
DB_NAME=wellbe
```

---

## 🏗 Migrações

### 📦 Gerenciamento com golang-migrate

Este projeto utiliza a ferramenta [golang-migrate/migrate](https://github.com/golang-migrate/migrate) para aplicar e reverter migrations de banco de dados de forma simples e segura.

#### 📥 Instalação do migrate

Execute no terminal:

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Certifique-se de que o binário `migrate` está no seu PATH.

#### 🚀 Aplicando migrations

Para aplicar todas as migrations pendentes:

```bash
migrate -path migrations -database "postgres://wellbe:wellbe@localhost:5432/wellbe?sslmode=disable" up
```

#### ⏪ Revertendo a última migration

```bash
migrate -path migrations -database "postgres://wellbe:wellbe@localhost:5432/wellbe?sslmode=disable" down 1
```

#### 🆕 Criando uma nova migration

Crie dois arquivos na pasta `migrations/` seguindo o padrão:

```
X_nome_da_migration.up.sql   # Para aplicar
X_nome_da_migration.down.sql # Para reverter
```
Exemplo:
```
4_add_updated_at_to_task.up.sql
4_add_updated_at_to_task.down.sql
```

#### 💡 Dicas
- O migrate controla o histórico das migrations aplicadas.
- Se ocorrer erro e o banco ficar "dirty", limpe com:
  ```bash
  migrate -path migrations -database "postgres://wellbe:wellbe@localhost:5432/wellbe?sslmode=disable" force <versão>
  ```
  (Substitua `<versão>` pela última migration válida.)
- Sempre mantenha as migrations versionadas e em ordem.

---

## 📂 Estrutura do Projeto

```
wellbe/
├── cmd/
│   └── api/
│       └── main.go          # Ponto de entrada
├── configs/
│   ├── app.env.example      # Template de configuração
├── migrations/              # Scripts SQL
├── pkg/
│   ├── database/            # Conexão com DB
│   └── models/              # Modelos
├── .env                     # Config local (não versionado)
├── docker-compose.yml       # Orquestração
└── Dockerfile               # Build da aplicação
```

---

## 🆘 Solução de Problemas

### 🔥 Erros comuns

1. **Banco não inicia**
    ```bash
    docker-compose logs db | grep -i error
    ```
2. **Migrações não aplicadas**  
   Verifique permissões na pasta `migrations/` e logs do PostgreSQL:
    ```bash
    docker-compose logs db | grep "executing"
    ```
3. **Portas em uso**
    ```bash
    sudo lsof -i :5432  # PostgreSQL
    sudo lsof -i :8080  # Aplicação
    ```

- **Se rodar localmente, o banco deve estar rodando no seu sistema.**
- **Se rodar via Docker, não precisa instalar PostgreSQL localmente.**
- **Se der erro de conexão, confira as variáveis do `.env` e se o banco está ativo.**
- **Para reiniciar só o banco:**
    ```bash
    docker-compose restart db
    ```
- **Para parar tudo:**
    ```bash
    docker-compose down
    ```

---

## 📄 Licença

MIT License - Veja LICENSE para detalhes.

---

### Como usar

1. **Copie todo o conteúdo acima**
2. **Cole em um novo arquivo** chamado `README.md` na raiz do seu projeto
3. **Salve** com codificação UTF-8

### Personalização

- Substitua `seu-usuario` pelo seu nome de usuário do GitHub
- Ajuste os nomes de arquivos/pastas conforme sua estrutura real
- Adicione seções específicas se necessário (ex: testes, deploy)

---