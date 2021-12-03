# --- Base ----
FROM golang:1.16.0-buster AS base
WORKDIR $GOPATH/src/github.com/jhonyzam/curso_golang/aula_8/products

# ---- Dependencies ----
FROM base AS dependencies
ENV GO111MODULE=on
COPY go.mod go.sum ./
RUN go mod download

# ---- Build ----
FROM dependencies AS build
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
