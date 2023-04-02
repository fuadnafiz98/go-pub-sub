FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go mod download github.com/bsm/ginkgo/v2
RUN go build -o bin/go-pubsub cmd/main.go

CMD ["./bin/go-pubsub"]
