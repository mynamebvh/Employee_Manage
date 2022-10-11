# Project Employee Manager

Project study golang - Zinza Company

## Manual Installation

If you need, configure the environment variables in file config.json, if you use docker-compose leave the variables set in the file config.json.example

```bash 
cp config.json.example config.json
docker-compose up  --build  -d
```

## Table of Contents
- [Features](#features)

## Features

- **Golang v1.19**: Stable version of go
- **Framework**: A stable version of [gin-go](https://github.com/gin-gonic/gin)
- **SQL databaseSQL**: [Mysql](https://www.mysql.com/) 
- **ORM**: [GORM](https://gorm.io/)
- **Testing**: 
- **API documentation**: with [swaggo](https://github.com/swaggo/swag) a go implementation
  of [swagger](https://swagger.io/)
- **Dependency management**: with [go modules](https://golang.org/ref/mod)
- **Environment variables**: using [viper](https://github.com/spf13/viper)
- **Docker support**