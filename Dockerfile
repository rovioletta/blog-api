FROM golang:1.25-alpine as builder

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN go build -o bin/blog-api ./cmd/blog-api

FROM alpine

COPY --from=builder /app/bin/blog-api /blog-api

CMD [ "/blog-api" ]