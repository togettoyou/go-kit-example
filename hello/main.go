package main

import (
	"github.com/togettoyou/go-kit-example/hello/dao"
	"github.com/togettoyou/go-kit-example/hello/endpoints"
	"github.com/togettoyou/go-kit-example/hello/service"
	httpTransport "github.com/togettoyou/go-kit-example/hello/transport/http"
	"log"
	"net/http"
)

func main() {
	helloDao := dao.NewHelloDAOImpl()
	helloService := service.NewHelloServiceImpl(helloDao)
	helloEndpoint := endpoints.HelloEndPoints{
		GetNameEndpoint: endpoints.MakeGetNameEndpoint(helloService),
		GetAgeEndpoint:  endpoints.MakeGetAgeEndpoint(helloService),
	}
	httpHandler := httpTransport.MakeHttpHandler(helloEndpoint)
	log.Fatalln(http.ListenAndServe(":8888", httpHandler))
}
