FROM golang:latest

WORKDIR /hash-service

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

RUN go build -o grpcserver .

EXPOSE 8080:8080

CMD ["./grpcserver"]