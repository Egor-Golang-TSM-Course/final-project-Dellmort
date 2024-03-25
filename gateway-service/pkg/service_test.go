package pkg

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func serverAddr() string {
	addr := os.Getenv("SERVER_ADDRESS")
	if addr == "" {
		addr = ":8080"
	}

	return addr
}

func TestService_CreateHashOK(t *testing.T) {
	ctx := context.Background()
	payload := ` {
		"payload": "test"
	} `

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(serverAddr(), opts)
	assert.NoError(t, err)

	srv := NewGatewayService(conn)

	answ, err := srv.CreateHash(ctx, payload)
	assert.NoError(t, err)

	assert.NotEmpty(t, answ)
}

func TestService_GetHashOK(t *testing.T) {
	ctx := context.Background()
	payload := ` {
		"payload": "test"
	} `

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(serverAddr(), opts)
	assert.NoError(t, err)

	srv := NewGatewayService(conn)

	answ, err := srv.GetHash(ctx, payload)
	assert.NoError(t, err)

	assert.NotEmpty(t, answ)
}

func TestService_CheckHashOK(t *testing.T) {
	ctx := context.Background()
	payload := ` {
		"payload": "test"
	} `

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(serverAddr(), opts)
	assert.NoError(t, err)

	srv := NewGatewayService(conn)

	ok, err := srv.CheckHash(ctx, payload)
	assert.NoError(t, err)

	assert.Equal(t, ok, true)
}

func TestService_GetHashOKError(t *testing.T) {
	ctx := context.Background()
	payload := ` {
		payload": "test
	 `

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(serverAddr(), opts)
	assert.NoError(t, err)

	srv := NewGatewayService(conn)

	answ, err := srv.GetHash(ctx, payload)
	assert.Error(t, err)

	assert.Equal(t, answ, "")
}
