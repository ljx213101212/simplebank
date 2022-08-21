# simplebank

This project is for learning and bulding backend design & develop & deploy.

> **Specially Thanks to Tech School** > https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/

# Note

## cheat sheet (frequently used in this project)

```
docker exec -it 8f987a9d0acb53cfbc308326bfd559efe7058bbd97317aec3de15e09d068b4bc /bin/sh

psql -h 127.0.0.1 -d simple_bank -U root -W
password: secret
\l
\dt
truncate table accounts CASCADE;

truncate table accounts RESTART IDENTITY CASCADE;

ALTER DATABASE <db name> SET DEFAULT_TRANSACTION_ISOLATION TO 'read committed';

```

## Go

> If you are new => go through this first
> https://go.dev/tour/welcome/1

- `blank modifier _`: Prevent go formatter from removing this from import

- `composition`:

- `go routine`:

- `channel types`: a communication mechanism that allows Goroutines to exchange data.

- `function receiver`: https://go.dev/tour/methods/4

- `struct tag`: https://github.com/amarjeetanandsingh/tgcon

- `GO111MODULE`: is an environment variable that can be set when using go for changing how Go imports packages. One of the first pain-points is that depending on the Go version, its semantics change. (https://maelvls.dev/go111module-everywhere/)

- `gin`: [HTTP Web framework](https://github.com/gin-gonic/gin)

```

gin: https://pkg.go.dev/github.com/gin-gonic/gin#section-readme
validator: https://pkg.go.dev/github.com/go-playground/validator#section-readme

```

## Unit Test

- recommended to use random data to make the code to be more concise and easier to understand.
- Test\* prefixed go file will be take as Unit Test file

## Database

> CRUD

```
select * from entries;
truncate table entries;

select * from transfers;
truncate table transfers;

select * from accounts;
truncate table accounts CASCADE;

```

### migrate

`migrateup`, `migratedown`

```
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

migrate create -ext sql -dir db/migration -seq add_users

```

### key concept

`foreign key`, `constraint`

```
ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency");
```

`ACID`

```
`Atomicity`, `Consistency`, `Isolation`, `Durablity`
```

`Deadlock prevention`

```
`database transaction isolation levels`, `mysql vs postgresSql`, `open terminal side by side to test deadlock`
```

### tools tradeoffs

database/sql

```
pros: standard lib, straightforward
cons manual mapping, easy to make mistakes
```

gorms

```
pros: short prod code
cons: 3x - 5x slower standard lib on high load

```

sqlx

```
pros: fast, easy to use, fields mapping, struct tags
cons: easy to make mistakes
```

sqlc

```
pros: very fast, auto code gen, catch errors during dev time
cons: mysql is not fully been supported yet (2022.08.06 experimental).
```

### sqlc

> https://github.com/kyleconroy/sqlc

## CI/CD

```
history | grep "some command we typed in recently"
```

[Github Actions - postgreSQL](https://docs.github.com/en/actions/using-containerized-services/creating-postgresql-service-containers)

[Docker Hub - postgreSQL](https://hub.docker.com/_/postgres/)

## Network

[SSL](https://www.cloudflare.com/learning/ssl/what-is-ssl/)
