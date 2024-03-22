package main

import (
	"gatewayservice/pkg"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	conn, err := grpc.Dial(":8080", opts)
	if err != nil {
		log.Fatal(err)
	}

	srv := pkg.NewGatewayService(conn)

	endpoints := pkg.Endpoints{
		CreateHash: pkg.CreateHashEndpoint(srv),
		CheckHash:  pkg.CheckHashEndpoint(srv),
		ExistsHash: pkg.ExistsHashEndpoint(srv),
	}

	handler := pkg.NewServerHTTP(&endpoints)

	if err := http.ListenAndServe(":3000", handler); err != nil {
		log.Println(err)
	}
}
