# dating-app-service

Dating App Backend Service

# Project Structure

All handler, repository, and service are inside `internal` folder

```bash
├───cmd
│   ├───rest
│   └───setup
├───config
│   └───db
│       └───migration
├───constants
├───internal
│   ├───auth
│   │   ├───handler
│   │   ├───model
│   │   ├───payload
│   │   ├───port
│   │   ├───repository
│   │   ├───routes
│   │   └───service
│   ├───premium
│   │   ├───handler
│   │   ├───model
│   │   ├───payload
│   │   ├───port
│   │   ├───repository
│   │   ├───routes
│   │   └───service
│   ├───recommendations
│   │   ├───handler
│   │   ├───model
│   │   ├───payload
│   │   ├───port
│   │   ├───repository
│   │   ├───routes
│   │   └───service
│   └───swipe
│       ├───handler
│       ├───model
│       ├───payload
│       ├───port
│       ├───repository
│       ├───routes
│       └───service
├───middleware
└───pkg
```

# Installation

Clone this repository

```
https://github.com/zikrykr/dating-app-service.git
```

## Prerequisite

- [Golang](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/engine/install/)
- [PostgreSQL](https://www.postgresql.org/download/)

## How to run

You can run this backend service via `docker-compose` file or you can run locally. If you run on your host machine, make sure you've install `go` and `postgresql`

- Copy `.env.example` to `.env` file (create new `.env` file on first set up), and set value as you want
- If you want to run via `docker-compose`, then simply just run this command line

```
$ make run
```

# Database Migrations

Sql migrations standard using open source tools

## Prerequisite

- [Go Migrate](https://github.com/golang-migrate/migrate)
- [Go Migrate - PGSQL Tutorial](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md)
- [Best practices: How to write migrations.](https://github.com/golang-migrate/migrate/blob/master/MIGRATIONS.md)

## Generate File migration

Generate file up and down version of migration sql. Warning: migrations file can only use `.sql` extension for scalability and concistency. Always wrap your sql queries with transaction.

`Assume under root folder`

```sh
$ migrate create -ext sql -dir config/db/migrations -seq create_users_table
```

## Migrate

### Patch (update version)

```sh
#need to export DB_DSN from .env. Only need once.

$ export $(cat .env | xargs)
$ migrate -database postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:5432/${DB_NAME}?sslmode=disable -path config/db/migrations up
```

### Rollback (downgrade version)

```sh
$ migrate -database postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:5432/${DB_NAME}?sslmode=disable -path config/db/migrations down
```
