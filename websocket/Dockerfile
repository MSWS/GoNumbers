FROM golang:alpine

ENV NUMBER_PORT=1234

WORKDIR /app

COPY go.mod go.sum numbers.go ./

RUN go mod tidy
RUN go build

CMD ["./numbers"]
