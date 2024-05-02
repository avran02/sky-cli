package tpl

var Dockerfile = `FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main", "serve"]
`

var Dockerignore = `.git/
.gitignore
.env
`
