package pkg

import (
	"context"
	"encoding/json"
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
	Hash string `json:"hash"`
}

type existsResponse struct {
	Exists bool  `json:"exists"`
	Err    error `json:"err,omitempty"`
}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func decodeExistsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req existsRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
