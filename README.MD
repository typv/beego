# BEEGO

# Getting Started

## Server Requirements

- Golang v1.22.5
- postgreSQL 16

## Installing preparation

1. Default Application $BASEPATH : `/home/app.user/beego`

2. Copy the .env file from .env.example under $BASEPATH, fill your config in .env file instead of example config

# I. Build with Docker

## 1. Prerequisite

In the root folder, find the .env file and change the following values

```bash
  APP_PORT=8888
    
  DB_HOST=db
  DB_PORT=5432
  DB_USERNAME=postgres
  DB_PASSWORD=secret
  DB_DATABASE=beego_fw
  DB_SSLMODE=disable
  DB_TIMEZONE=UTC
```

## 2. Setup docker


```bash
  make up
  make installCLI
  make install
  make vendor
```

### 3. Run server

Development:
```bash
  make dev
```


Develop url: http://localhost:9889