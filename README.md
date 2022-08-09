# simplebank

This project is for learning and bulding backend design & develop & deploy.

> **Specially Thanks to Tech School** > https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/

# Note

## Go

- `blank modifier _`: Prevent go formatter from removing this from import

- `composition`:

- `go routine`:

- `channel types`: a communication mechanism that allows Goroutines to exchange data.

## Unit Test

- recommended to use random data to make the code to be more concise and easier to understand.
- Test\* prefixed go file will be take as Unit Test file

## Database

### cheat sheet

```
psql -h 127.0.0.1 -d simple_bank -U root -W
\l
\dt

ALTER DATABASE <db name> SET DEFAULT_TRANSACTION_ISOLATION TO 'read committed';

```

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

### transaction

> ACID  
> `Atomicity`, `Consistency`, `Isolation`, `Durablity`

> Deadlock prevention
> `database transaction isolation levels`, `mysql vs postgresSql`, `open terminal side by side to test deadlock`

## CI/CD

```
history | grep "some command we typed in recently"
```

## Network

[SSL](https://www.cloudflare.com/learning/ssl/what-is-ssl/)
