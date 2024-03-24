package main

import (
	"fmt"
	service "hashservice/pkg"
	storagemap "hashservice/storage/map"
)

func main() {
	srv := service.NewHashService(
		storagemap.NewStorageMap(),
	)

	fmt.Println("GRPC server started...")
	srv.Run()
}
