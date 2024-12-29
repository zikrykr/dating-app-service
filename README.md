# dating-app-service

Dating App Backend Service

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
