FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod tidy

CMD ["go", "run", "cmd/ordersystem/main.go", "cmd/ordersystem/wire_gen.go"]
