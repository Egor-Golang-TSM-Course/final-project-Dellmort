package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"hashservice/hashservice"
	"hashservice/storage"
	"net"

	"google.golang.org/grpc"
)

type HashService struct {
	storage storage.Storage
	hashservice.UnimplementedHashServiceServer
}

func NewHashService(s storage.Storage) *HashService {
	return &HashService{
		storage: s,
	}
}

func (hs *HashService) CreateHash(ctx context.Context, req *hashservice.HashRequest) (*hashservice.HashResponse, error) {
	h := sha256.New()
	h.Write([]byte(req.GetPayload()))

	hash := fmt.Sprintf("%x", h.Sum(nil))
	hs.storage.SaveHash(req.GetPayload(), hash)

	return &hashservice.HashResponse{
		Hash: hash,
	}, nil
}

func (hs *HashService) GetHash(ctx context.Context, req *hashservice.HashRequest) (*hashservice.HashResponse, error) {
	hash, err := hs.storage.GetHash(req.GetPayload())
	if err != nil {
		return nil, err
	}

	return &hashservice.HashResponse{
		Hash: hash,
	}, nil
}
func (hs *HashService) CheckHash(ctx context.Context, req *hashservice.HashRequest) (*hashservice.HashResponseOther, error) {
	exists, err := hs.storage.CheckHash(req.GetPayload())
	if err != nil {
		return nil, err
	}

	return &hashservice.HashResponseOther{
		Exists: exists,
	}, nil
}

func (hs *HashService) Run() error {
	// os.Getenv("PORT")
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	hashservice.RegisterHashServiceServer(server, hs)

	err = server.Serve(l)
	if err != nil {
		return err
	}

	return nil
}
