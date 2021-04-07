package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/togettoyou/go-kit-example/hello/endpoints"
	"net/http"
	"os"
)

func MakeHttpHandler(eps endpoints.HelloEndPoints) http.Handler {
	r := mux.NewRouter()

	kitLog := log.NewLogfmtLogger(os.Stderr)
	kitLog = log.With(kitLog, "ts", log.DefaultTimestampUTC)
	kitLog = log.With(kitLog, "caller", log.DefaultCaller)
	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(kitLog)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	r.Methods("GET").Path("/").Handler(kithttp.NewServer(
		eps.SayHelloEndpoint,
		decodeSayHelloRequest,
		encodeJSONResponse,
		options...,
	))

	return r
}

// Request拦截
// 参数校验
func decodeSayHelloRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("decodeSayHelloRequest")
	return nil, nil
}

// Response拦截
func encodeJSONResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Println("encodeJSONResponse")
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// 错误拦截
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
