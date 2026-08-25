FROM golang:1.25.6

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/server

EXPOSE 8080

CMD ["./main"]