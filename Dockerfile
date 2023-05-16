FROM golang:1.19

WORKDIR /service

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/crnt-chat-service ./cmd/*.go

CMD ["./bin/crnt-chat-service"]
