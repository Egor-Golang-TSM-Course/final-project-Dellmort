package pkg

import (
	"context"
	"hashservice/hashservice"

	"google.golang.org/grpc"
)

type Service interface {
	CreateHash(ctx context.Context, payload string) (string, error)
	CheckHash(ctx context.Context, Hash string) (bool, error)
	GetHash(ctx context.Context, payload string) (string, error)
}

type GatewayService struct {
	client hashservice.HashServiceClient
}

func NewGatewayService(conn *grpc.ClientConn) Service {
	c := hashservice.NewHashServiceClient(conn)

	return &GatewayService{
		client: c,
	}
}

func (s *GatewayService) CreateHash(ctx context.Context, payload string) (string, error) {
	resp, err := s.client.CreateHash(
		ctx,
		&hashservice.HashRequest{
			Payload: payload,
		},
	)
	if err != nil {
		return resp.GetHash(), err
	}

	return resp.GetHash(), nil
}

func (s *GatewayService) CheckHash(ctx context.Context, Hash string) (bool, error) {
	resp, err := s.client.CheckHash(
		ctx,
		&hashservice.HashRequest{
			Payload: Hash,
		},
	)
	if err != nil {
		return resp.GetExists(), err
	}

	return resp.GetExists(), nil
}

func (s *GatewayService) GetHash(ctx context.Context, payload string) (string, error) {
	resp, err := s.client.GetHash(
		ctx,
		&hashservice.HashRequest{
			Payload: payload,
		},
	)
	if err != nil {
		return resp.GetHash(), err
	}

	return resp.GetHash(), nil
}
