# BTC Billionaire [AnyMind](https://anymindgroup.com/) Project

#### 👨‍💻 List of what has been used:

* [echo](https://github.com/labstack/echo) - Web framework
* [sqlx](https://github.com/jmoiron/sqlx) - Extensions to database/sql.
* [pgx](https://github.com/jackc/pgx) - PostgreSQL driver and toolkit for Go
* [viper](https://github.com/spf13/viper) - Go configuration with fangs
* [zap](https://github.com/uber-go/zap) - Logger
* [validator](https://github.com/go-playground/validator) - Go Struct and Field validation
* [uuid](https://github.com/google/uuid) - UUID
* [migrate](https://github.com/golang-migrate/migrate) - Database migrations. CLI and Golang library.
* [swag](https://github.com/swaggo/swag) - Swagger
* [testify](https://github.com/stretchr/testify) - Testing toolkit
* [gomock](https://github.com/golang/mock) - Mocking framework
* [CompileDaemon](https://github.com/githubnemo/CompileDaemon) - Compile daemon for Go
* [Docker](https://www.docker.com/) - Docker

---

### 🚀 Docker-compose files:

```bash
# run postgresql
docker-compose.local.yml 

# run docker development environment
docker-compose.dev.yml 
```

### 🗄️ Migrations:

Install
the [golang-migrate CLI](https://github.com/golang-migrate/migrate/blob/5bf05dc3236ef077e5927c9ca9ca02857a87c582/cmd/migrate/README.md#installation)

```bash
# create the database
make migrate_up

# drop the database
make migrate_down
```

---

### 📦 Docker development usage:

```bash
make docker
```

### 📍 Recommendation for local development:

```bash
# run all containers
make local 

# it's easier way to attach debugger or rebuild/rerun project
make run 
```

---

## API

#### SWAGGER UI:

http://localhost:8080/swagger/index.html

### Send bitcoin

```bash
curl --location --request POST 'http://localhost:8080/api/v1/transactions' \
--header 'Content-Type: application/json' \
--data-raw '{
    "datetime": "2022-12-31T15:00:00+03:00",
    "amount": 10
}'
```

### Get transactions per hour

```bash
curl --location --request GET \
'http://localhost:8080/api/v1/transactions?startDatetime=2022-01-01T18:00:00-05:00&endDatetime=2023-12-31T21:00:00+08:00'
```

### Get total balance

```bash
curl --location --request GET 'http://localhost:8080/api/v1/transactions/balance' 
```

---

## Client

There's a mock client that sends a random amount of bitcoin from a random place (timezone) each second

```bash
make client
```