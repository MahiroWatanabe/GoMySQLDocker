FROM golang:latest

RUN mkdir /app
WORKDIR /app

# Gin
RUN go get github.com/gin-gonic/gin
# connect sql
RUN go get github.com/go-sql-driver/mysql
# orm
RUN go get github.com/jinzhu/gorm

RUN go mod download