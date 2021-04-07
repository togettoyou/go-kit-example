package endpoints

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/togettoyou/go-kit-example/hello/service"
)

type HelloEndPoints struct {
	SayHelloEndpoint endpoint.Endpoint
}

func MakeSayHelloEndpoint(helloService service.HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		fmt.Println("endpoint层")
		return helloService.SayHello(ctx), nil
	}
}
