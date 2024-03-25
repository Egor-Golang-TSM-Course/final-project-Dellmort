package main

import (
	"fmt"
	"gatewayservice/config"
	"gatewayservice/pkg"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config.MustLoadEnv()

	addr := os.Getenv("SERVER_ADDRESS")
	if addr == "" {
		addr = ":8080"
	}

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(addr, opts)
	if err != nil {
		log.Fatal(err)
	}

	srv := pkg.NewGatewayService(conn)

	endpoints := pkg.Endpoints{
		CreateHash: pkg.CreateHashEndpoint(srv),
		GetHash:    pkg.GetHashEndpoint(srv),
		CheckHash:  pkg.ExistsHashEndpoint(srv),
	}

	handler := pkg.NewServerHTTP(&endpoints)

	fmt.Println("Server started...")

	if err := http.ListenAndServe(os.Getenv("SERVER_PORT"), handler); err != nil {
		log.Println(err)
	}
}
