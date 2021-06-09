package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/togettoyou/go-kit-example/hello/service"
	"log"
)

type HelloEndPoints struct {
	GetNameEndpoint endpoint.Endpoint
	GetAgeEndpoint  endpoint.Endpoint
}

func MakeGetNameEndpoint(helloService service.HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		log.Println("endpoint层-MakeGetNameEndpoint")
		return helloService.GetName(ctx), nil
	}
}

func MakeGetAgeEndpoint(helloService service.HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		log.Println("endpoint层-MakeGetAgeEndpoint")
		return helloService.GetAge(ctx), nil
	}
}
