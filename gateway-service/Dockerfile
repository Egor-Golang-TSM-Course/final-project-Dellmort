FROM golang:latest

COPY ./gateway-service/go.mod .
COPY ./gateway-service/go.sum .

RUN go mod download

COPY ./gateway-service .

RUN go build -o server .

EXPOSE 3000:3000

CMD ["./server"]