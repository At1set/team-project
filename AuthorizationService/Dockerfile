# Используем более новую версию Go
FROM golang:1.23-alpine

# Остальные инструкции остаются без изменений
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]