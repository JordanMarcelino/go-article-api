# Go-Lang Article API

## Description

Practicing creating APIs using the Go programming language.

## Tech Stack

- Golang : https://github.com/golang/go
- PostgreSQL (Database) : https://github.com/postgres/postgres

## Framework & Library

- GoFiber (HTTP Framework) : https://github.com/gofiber/fiber
- GORM (ORM) : https://github.com/go-gorm/gorm
- Viper (Configuration) : https://github.com/spf13/viper
- Golang Migrate (Database Migration) : https://github.com/golang-migrate/migrate
- Go Playground Validator (Validation) : https://github.com/go-playground/validator
- Logrus (Logger) : https://github.com/sirupsen/logrus
- Wire (Dependency Injetion) : https://github.com/google/wire

## API Documentation

All API endpoint is in `api` folder

## How To Run / Install on Your Local Machine

--> Clone the repository using command bellow:

```bash
git clone https://github.com/JordanMarcelino/go-article-api.git
```

--> Move into the directory :

```bash
cd golang-article-api
```

--> Build & run docker container :

```bash
docker build -t golang-article-api -f ./build/package/Dockerfile .
docker compose -f ./build/package/docker-compose.yaml -up -d
```

--> Stop the container :

```bash
cd ./build/package/
docker-compose down
```