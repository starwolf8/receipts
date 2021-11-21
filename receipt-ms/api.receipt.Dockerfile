# syntax=docker/dockerfile:1

#learning reference - https://docs.docker.com/language/golang/build-images/

FROM golang:1.16-alpine

# create a directory inside the image that we are building
WORKDIR /app

COPY receipt-ms/go.mod ./
COPY receipt-ms/go.sum ./

ADD ./receipt-ms/ /app

COPY receipt-ms/*.go ./
RUN go mod tidy
RUN go mod download

RUN go get github.com/go-sql-driver/mysql 

RUN go build -o /receipt-ms

CMD [ "/receipt-ms" ]
