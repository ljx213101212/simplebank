# simplebank

This project is for learning and bulding backend design & develop & deploy.

> **Specially Thanks to Tech School** > https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/

# Config

```
change prod.template.env to prod.env and fill in DB connection string from your local.
```

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

openssl rand -hex 64 | head -c 32

```

## Go

> If you are new => go through this first
> https://go.dev/tour/welcome/1

- `blank modifier _`: Prevent go formatter from removing this from import

- `composition`:

- `go routine`:

- `Type assertions`: https://go.dev/tour/methods/15

- `channel types`: a communication mechanism that allows Goroutines to exchange data.

- `function receiver`: https://go.dev/tour/methods/4

- `struct tag`: https://github.com/amarjeetanandsingh/tgcon

- `GO111MODULE`: is an environment variable that can be set when using go for changing how Go imports packages. One of the first pain-points is that depending on the Go version, its semantics change. (https://maelvls.dev/go111module-everywhere/)

- `gin`: [HTTP Web framework](https://github.com/gin-gonic/gin)

```

gin: https://pkg.go.dev/github.com/gin-gonic/gin#section-readme
validator: https://pkg.go.dev/github.com/go-playground/validator#section-readme

```

- `jwt_go`: [go packge to generate jwt token](https://github.com/dgrijalva/jwt-go)

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

```
alpine version:  lite version
```

[Dockerfile](https://docs.docker.com/engine/reference/builder/#from)

```
FROM: base image (check from Docker Hub)
WORKDIR /app
COPY [src] [target]
RUN go build -o main main.go

EXPOSE [network port]

```

[Docker CLI](https://docs.docker.com/engine/reference/commandline)

[Docker Compose](https://docs.docker.com/compose/startup-order/)

```
docker compose down
docker compose up
```

- [wait-for](https://github.com/Eficode/wait-for)

- [encrypt secret project info in openssl](https://www.madboa.com/geek/openssl/#encrypt-simple)

```make
encryptenv
decryptenv
```

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

- [Claim Standard](https://tools.ietf.org/html/rfc7519#section-4.1)

[PASETO](https://dev.to/techschoolguru/why-paseto-is-better-than-jwt-for-token-based-authentication-1b0c)

## Cloud

- Create a free tier amazon account
- Create a ECR repository
- Use Github Action to deploy
- Create IAM for aws
- Create AWS RDS as a production DB
- Create Secrets Manager for Application Secret storage
- Add Secrets Manager access permission from IAM
- AWS CLI
- install jq to process secret string to prod.env formart

- Create a EKS repository
- Create IAM role for accessing cluster service role
- wait for 15 mins
- Add a node group in EKS
- Create IAM role for node group
- wait for 5 mins

### Amazon ECR

### Github Action

```
Amazon ECR "Login" Action for GitHub Actions

```

### IAM

```
add
AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY
into Githu Action
```

### AWS RDS

```
migrateuprds:
	migrate -path db/migration -database "${DB_SOURCE_RDS}" -verbose up

```

### AWS CLI

- ARN -Amazon Source Name
- https://docs.aws.amazon.com/cli/v1/userguide/install-virtualenv.html

```
pip install --user virtualenv
python -m virtualenv ~/cli-ve
source ~/cli-ve/bin/activate

aws configure
 - create new access key from IAM to fill in

cat ~/.aws/credentials
cat ~/.aws/config

(remember to add secret manager permission to IAM group)
aws secretsmanager get-secret-value --secret-id simple_bank
aws secretsmanager get-secret-value --secret-id simple_bank --query SecretString --output text

aws secretsmanager get-secret-value --secret-id simple_bank --query SecretString --output text | jq -r 'to_entries|map("\(.key)=\(.value)")|.[]'


docker pull [image url from ecr]

https://docs.aws.amazon.com/cli/latest/reference/ecr/get-login-password.html

aws ecr get-login-password


aws ecr get-login-password \
| docker login \
    --username AWS \
    --password-stdin <aws_account_id>.dkr.ecr.<region>.amazonaws.com


docker run -p 8080:8080 [image url from ecr]
```

-

### AWS EKS

> a managed service that makes it easy for you to use Kubernetes on AWS without needing to install and operate your own Kubernetes control plane.

`Kubernates`:

- An open-source container orchestration engine
- For automating deployment, scaling, and management of containerized applications

```
Kubelet agaent: make sure containers run inside pods

Container runtimes: Docker, containerd, CRI-O

Kube-proxy: maintain network rules, allow communication with pods

vpc: Virtual Private Cloud

node group -> capacity type -> spot: cheaper but not stable
node group -> capacity type -> on-demand: stable, safe for api

```
