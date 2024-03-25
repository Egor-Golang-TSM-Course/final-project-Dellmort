package pkg

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Request получаем  от клиента по запросу
type Request struct {
	Payload string `json:"payload,omitempty"`
}

// Response для отправки ответа
type Response struct {
	Hash string `json:"hash"`
	Err  error  `json:"err,omitempty"`
}

type existsRequest struct {
	Payload string `json:"payload"`
}

type existsResponse struct {
	Exists bool  `json:"exists"`
	Err    error `json:"err,omitempty"`
}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("json error: %w", err)
	}

	return req, nil
}
func decodeExistsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req existsRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}

	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
