package pkg

import (
	"context"
	"fmt"
	"gatewayservice/hash/api"

	"google.golang.org/grpc"
)

type Service interface {
	CreateHash(ctx context.Context, payload string) (string, error)
	CheckHash(ctx context.Context, Hash string) (bool, error)
	GetHash(ctx context.Context, payload string) (string, error)
}

type GatewayService struct {
	client api.HashServiceClient
}

func NewGatewayService(conn *grpc.ClientConn) Service {
	c := api.NewHashServiceClient(conn)

	return &GatewayService{
		client: c,
	}
}

func (s *GatewayService) CreateHash(ctx context.Context, payload string) (string, error) {
	resp, err := s.client.CreateHash(
		ctx,
		&api.HashRequest{
			Payload: payload,
		},
	)
	if err != nil {
		return resp.GetHash(), fmt.Errorf("CreateHash %w", err)
	}

	return resp.GetHash(), nil
}

func (s *GatewayService) CheckHash(ctx context.Context, payload string) (bool, error) {
	resp, err := s.client.CheckHash(
		ctx,
		&api.HashRequest{
			Payload: payload,
		},
	)
	if err != nil {
		return false, err
	}

	return resp.GetExists(), nil
}

func (s *GatewayService) GetHash(ctx context.Context, payload string) (string, error) {
	resp, err := s.client.GetHash(
		ctx,
		&api.HashRequest{
			Payload: payload,
		},
	)
	if err != nil {
		return "", err
	}

	return resp.GetHash(), nil
}
