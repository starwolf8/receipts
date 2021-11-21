# syntax=docker/dockerfile:1

#learning reference - https://docs.docker.com/language/golang/build-images/

FROM golang:1.16-alpine

# create a directory inside the image that we are building
WORKDIR /app


COPY go.mod ./
COPY go.sum ./


ADD . /app

COPY *.go ./
RUN go mod tidy
RUN go mod download

RUN go get github.com/go-sql-driver/mysql 

RUN go build -o /api-product

# EXPOSE 5001

CMD [ "/api-product" ]
