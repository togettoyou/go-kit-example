package main

import (
	"github.com/togettoyou/go-kit-example/hello/dao"
	"github.com/togettoyou/go-kit-example/hello/endpoints"
	"github.com/togettoyou/go-kit-example/hello/service"
	httptransport "github.com/togettoyou/go-kit-example/hello/transport/http"
	"net/http"
)

func main() {
	helloService := service.MakeHelloServiceImpl(&dao.HelloDAOImpl{})
	httpHandler := httptransport.MakeHttpHandler(endpoints.HelloEndPoints{
		SayHelloEndpoint: endpoints.MakeSayHelloEndpoint(helloService),
	})
	http.ListenAndServe(":8888", httpHandler)
}
