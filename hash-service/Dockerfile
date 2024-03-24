FROM golang:latest

COPY ./hash-service/go.mod .
COPY ./hash-service/go.sum .

RUN go mod download

COPY /hash-service .

RUN go build -o servergrpc .

EXPOSE 8080:8080

CMD ["./servergrpc"]