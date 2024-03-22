package pkg

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewServerHTTP(endpoints *Endpoints) *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("/createhash", httptransport.NewServer(
		endpoints.CreateHash,
		decodeRequest,
		encodeResponse,
	))

	router.Handle("/checkhash", httptransport.NewServer(
		endpoints.CheckHash,
		decodeRequest,
		encodeResponse,
	))

	router.Handle("/exists", httptransport.NewServer(
		endpoints.ExistsHash,
		decodeExistsRequest,
		encodeResponse,
	))

	return router
}
