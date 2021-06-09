package main

import (
	"github.com/togettoyou/go-kit-example/hello/dao"
	"github.com/togettoyou/go-kit-example/hello/endpoints"
	"github.com/togettoyou/go-kit-example/hello/service"
	grpcTransport "github.com/togettoyou/go-kit-example/hello/transport/grpc"
	"github.com/togettoyou/go-kit-example/hello/transport/grpc/pb"
	httpTransport "github.com/togettoyou/go-kit-example/hello/transport/http"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	helloDao := dao.NewHelloDAOImpl()
	helloService := service.NewHelloServiceImpl(helloDao)
	helloEndpoint := endpoints.HelloEndPoints{
		GetNameEndpoint: endpoints.MakeGetNameEndpoint(helloService),
		GetAgeEndpoint:  endpoints.MakeGetAgeEndpoint(helloService),
	}
	go newHttpServer(helloEndpoint)
	go newGRPCServer(helloEndpoint)
	select {}
}

func newHttpServer(endpoints endpoints.HelloEndPoints) {
	httpHandler := httpTransport.NewHttpHandler(endpoints)
	log.Fatalln(http.ListenAndServe(":8888", httpHandler))
}

func newGRPCServer(endpoints endpoints.HelloEndPoints) {
	server := grpc.NewServer()
	grpcHandler := grpcTransport.NewGRPCHandler(endpoints)
	pb.RegisterHelloServiceServer(server, grpcHandler)
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(server.Serve(lis))
}
