package pkg

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewServerHTTP(endpoints *Endpoints) *http.ServeMux {
	router := http.NewServeMux()

	// ставим json заголовок в каждом ответе
	handleMiddleware := func(pattern string, handler http.Handler) {
		router.Handle(pattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			handler.ServeHTTP(w, r)
		}))
	}

	handleMiddleware(
		"/createhash",
		httptransport.NewServer(
			endpoints.CreateHash,
			decodeRequest,
			encodeResponse,
		),
	)

	handleMiddleware(
		"/gethash",
		httptransport.NewServer(
			endpoints.GetHash,
			decodeRequest,
			encodeResponse,
		),
	)

	handleMiddleware(
		"/checkhash",
		httptransport.NewServer(
			endpoints.CheckHash,
			decodeExistsRequest,
			encodeResponse,
		),
	)

	return router
}
