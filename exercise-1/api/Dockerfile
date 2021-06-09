FROM golang:latest

RUN go get -u -v github.com/go-chi/chi/v5 && go get -u -v go.mongodb.org/mongo-driver

WORKDIR /go/src/app
COPY ./src /go/src/app/

RUN go mod tidy