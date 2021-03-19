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
