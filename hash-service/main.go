package main

import (
	service "hashservice/pkg"
	storagemap "hashservice/storage/map"
)

func main() {
	srv := service.NewHashService(
		storagemap.NewStorageMap(),
	)

	srv.Run()
}
