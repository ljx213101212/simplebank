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

## Network and Security

[SSL](https://www.cloudflare.com/learning/ssl/what-is-ssl/)

[Bcrypt](https://blog.boot.dev/cryptography/bcrypt-step-by-step/)

> ALG + COST + SALT + HASH

[Hash vs Encryption](https://www.ssl2buy.com/wiki/difference-between-hashing-and-encryption)

> In short, encryption is a two-way function that includes encryption and decryption whilst hashing is a one-way function that changes a plain text to a unique digest that is irreversible.

- Symmetric digital Signature Algorithm

  - The **Same secret key** is used to **sign & verify** token
  - For **local** use: internal services, where the secret key can be shared.
  - HS256, HS384, HS512
    - HS256 = HMAC + SHA256
    - HMAC: Hash-based Message Authentication Code
    - SHA: Secure Hash Algorithm
    - 256/384/512: number of output bits

- Asymmetric digital signature algorithm
  - The private key is used to **sign** token
  - The public key is used to **verify** token
  - For **public** use: internal service signs token, but external service needs to verify it
  - RS256, RS384, RS512 || PS256, PS384, PS512 || ES256, ES384, ES512
  - RS256 = RSA PKCSv1.5 + SHA256 [PKCS: Public-Key Cryptography Standards]
  - PS256 = RSA PSS + SHA256 [PSS: Probabilistic Signature Scheme]
  - ES256 = ECDSA + SHA256 [ECDSA: Elliptic Curve Digital Signature Algorithm]

[JWT](https://jwt.io/)

- Weak algorithms

  - Give developers too many algorithms to choose
  - Some algorithms are known to be vulnerable:
    - RSA PKCSv1.5: padding oracle attack
    - ECDSA: invalid-curve attack

- Trivial Forgery
  - Set "alg" header to "none"
  - Set "alg" header to "HS256" while the server normally verifies token with a RSA public key

[PASETO](https://dev.to/techschoolguru/why-paseto-is-better-than-jwt-for-token-based-authentication-1b0c)
