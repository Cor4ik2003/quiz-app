FROM golang:1.21

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o auth-service ./cmd/main.go

EXPOSE 8080
CMD ["./auth-service"]
