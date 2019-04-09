FROM golang:latest

RUN apt-get update
RUN go get -u "github.com/gin-gonic/gin"
RUN go get -u "github.com/jinzhu/gorm"
RUN go get -u "github.com/jinzhu/gorm/dialects/mysql"
RUN go get -u "github.com/go-sql-driver/mysql"
