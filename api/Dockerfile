FROM golang:1.27.1

WORKDIR /app

# ローカルからイメージ内にコピー
COPY . .

RUN go build -o main ./cmd/server

EXPOSE 8080

CMD ["./main"]