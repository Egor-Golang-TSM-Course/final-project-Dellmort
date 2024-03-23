FROM golang:latest

WORKDIR /gateway-service

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

RUN go build -o server .

EXPOSE 3000:3000

CMD ["./server"]