# golang 環境

## 初期化

- environment (using direnv)
  - service account for gcp.
    - save it as serviceAccount.json

```bash
> cp .envrc.tepl .envrc
> direnv allow
```

## server starting

- docker

```bash
# build image
> docker-compose build

# container start
> docker-compose up -d default-db

# server start
> realize start
```

- example of default server

```bash
> curl -X GET http://localhost:8080/customers/DUMMY_USER_ID -v
```

### database

- generate server code by sql boiler

```bash
> make sqlboiler
```

### testing

```bash
> dokcer-compose up -d default-db

> make test
```

## production

### build

```bash
> docker build -f ./docker/production/default/Dockerfile .
```

## about layer

### infrastructure

- data layer
- It is responsibility to handle the data
- interested in database etc.

#### entity

- struct for setting the result of SQL etc....

#### infra/repository

- write the actual data manipulation process

### domain

#### model

- domain model

#### domain/repository

- write interface for infrastructure/repository and convert entity to domain

#### domain/service

- write application logic using repository

### usecase layer

- write usecase using repository and service

### handler

- write the process about request and response

### internal/pkg

- shared code
