# Pismo Challenge API

## Requisitos
- Go 1.16
- PostgreSQL 13.2
- Docker 19.03
- Docker Compose 1.21

## Pacotes utilizados
- [Echo](https://echo.labstack.com) para expor os recursos REST
- [GORM](https://gorm.io) para persistência de dados

## Configuração
A configuração é feita por meio de variáveis de ambiente.
- `PORT`: porta na qual será exposta a API
- `DATABASE_HOST`: endereço do banco de dados
- `DATABASE_USER`: usuário do banco de dados
- `DATABASE_PASSWORD`: senha do banco de dados
- `DATABASE_DBNAME`: nome do banco de dados
- `DATABASE_PORT`: porta do banco de dados

## Execução
Para facilitar a execução o projeto foi portado para o Docker.

Executar o comando no terminal dentro da pasta do projeto:
```shell
docker-compose build --no-cache && docker-compose up
```

## Endpoints

### Criação de uma conta
```shell
curl --request POST \
  --url http://localhost:8080/accounts \
  --header 'Content-Type: application/json' \
  --data '{
	"document_number": "12345678900"
}'
```

### Consulta de informações de uma conta
```shell
curl --request GET \
  --url http://localhost:8080/accounts/1
```

### Consulta de tipos de operação disponíveis
```shell
curl --request GET \
  --url http://localhost:8080/operation-types
```

### Criação de uma transação
```shell
curl --request POST \
  --url http://localhost:8080/transactions \
  --header 'Content-Type: application/json' \
  --data '{
	"account_id": 1,
	"operation_type_id": 1,
	"amount": -123.45
}'
```