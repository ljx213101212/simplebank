# simplebank

This project is for learning and bulding backend design & develop & deploy.

> **Specially Thanks to Tech School** > https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/

# Note

## Go

- `blank modifier _`: Prevent go formatter from removing this from import

## Unit Test

- recommended to use random data to make the code to be more concise and easier to understand.
- Test\* prefixed go file will be take as Unit Test file

## Database

### migrate

`migrateup`, `migratedown`

```
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

```

#### tools tradeoffs

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

#### sqlc

> https://github.com/kyleconroy/sqlc

## CI/CD

```
history | grep "some command we typed in recently"
```

## Network

[SSL](https://www.cloudflare.com/learning/ssl/what-is-ssl/)
