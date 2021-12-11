# Ebanx Account

Api Versão 1 - Account


### 📋 Pré-requisitos

Ferramentas:

- [Golang](https://golang.org/doc/install)

## 📦 Desenvolvimento

Existe dois comandos básicos para execução do projeto:

- `go run main.go`: ira executar a aplicação a partir dos pacotes do Go em sua maquina.
- `go build`: ira realizar o build da aplicação gerando um arquivo executavel da mesma.

## 🗂 Arquitetura

- `Hexagonal`: dentro das possibilidades com o que foi solicitado foi aplicado a arquitetura.

### Descrição dos diretórios e arquivos mais importantes:

- `./app`: Camada contendo as rotas utilizadas pela api
- `./handle`: Camada de captura das chamadas pelas rotas
- `./domain`: Camada contendo as Structs (models).
- `./services`: Camada contendo as regras de negócio.

## 🛠️ Construído com

* [go mod](https://blog.golang.org/using-go-modules) - Dependência


