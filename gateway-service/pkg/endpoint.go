package pkg

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

var (
	ErrUnexpected = errors.New("unexpected error")
)

func CreateHashEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(Request)
		if !ok {
			return Response{Err: ErrUnexpected}, ErrUnexpected
		}

		hash, err := srv.CreateHash(ctx, req.Payload)
		if err != nil {
			return Response{Err: err}, err
		}

		return Response{Hash: hash}, nil
	}
}

func CheckHashEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(Request)
		if !ok {
			return Response{Err: ErrUnexpected}, ErrUnexpected
		}

		hash, err := srv.GetHash(ctx, req.Payload)
		if err != nil {
			return Response{Err: err}, err
		}

		return Response{Hash: hash}, nil
	}
}

func ExistsHashEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(existsRequest)
		if !ok {
			return existsResponse{Err: ErrUnexpected}, ErrUnexpected
		}

		state, err := srv.CheckHash(ctx, req.Hash)
		if err != nil {
			return existsResponse{Err: err}, err
		}

		return existsResponse{Exists: state}, nil
	}
}

type Endpoints struct {
	CreateHash endpoint.Endpoint
	CheckHash  endpoint.Endpoint
	ExistsHash endpoint.Endpoint
}

func (e Endpoints) Create(ctx context.Context) (string, error) {
	req := Request{}

	resp, err := e.CreateHash(ctx, req)
	if err != nil {
		return "", err
	}

	r, ok := resp.(Response)
	if !ok {
		return "", ErrUnexpected
	}

	if r.Err != nil {
		return "", r.Err
	}

	return r.Hash, nil
}

func (e Endpoints) Check(ctx context.Context) (string, error) {
	req := Request{}

	resp, err := e.CheckHash(ctx, req)
	if err != nil {
		return "", err
	}

	r, ok := resp.(Response)
	if !ok {
		return "", ErrUnexpected
	}

	if r.Err != nil {
		return "", r.Err
	}

	return r.Hash, nil
}

func (e Endpoints) Exists(ctx context.Context) (bool, error) {
	req := Request{}

	resp, err := e.ExistsHash(ctx, req)
	if err != nil {
		return false, err
	}

	r, ok := resp.(existsResponse)
	if !ok {
		return false, ErrUnexpected
	}

	if r.Err != nil {
		return false, r.Err
	}

	return r.Exists, nil
}
