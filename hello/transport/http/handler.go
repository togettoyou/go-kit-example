package http

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	kitLog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/togettoyou/go-kit-example/hello/endpoints"
	"log"
	"net/http"
	"os"
)

func NewHttpHandler(eps endpoints.HelloEndPoints) http.Handler {
	r := mux.NewRouter()
	options := getServerOptions()
	r.Methods("GET").Path("/name").Handler(newServer(eps.GetNameEndpoint, options))
	r.Methods("GET").Path("/age").Handler(newServer(eps.GetAgeEndpoint, options))
	return r
}

func newServer(e endpoint.Endpoint, options []kitHttp.ServerOption) http.Handler {
	return kitHttp.NewServer(
		e,
		decodeRequest,
		encodeJSONResponse,
		options...,
	)
}

func getServerOptions() []kitHttp.ServerOption {
	logger := kitLog.NewLogfmtLogger(os.Stderr)
	logger = kitLog.With(logger, "ts", kitLog.DefaultTimestampUTC)
	logger = kitLog.With(logger, "caller", kitLog.DefaultCaller)
	options := []kitHttp.ServerOption{
		kitHttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kitHttp.ServerErrorEncoder(encodeError),
	}
	return options
}

func decodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("Request拦截-decodeRequest")
	return r, nil
}

func encodeJSONResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("Response拦截-encodeJSONResponse")
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

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
