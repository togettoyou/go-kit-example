package grpc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	kitGRPC "github.com/go-kit/kit/transport/grpc"
	"github.com/togettoyou/go-kit-example/hello/endpoints"
	"github.com/togettoyou/go-kit-example/hello/transport/grpc/pb"
	"log"
)

type grpcServer struct {
	getName kitGRPC.Handler
	getAge  kitGRPC.Handler
	pb.UnimplementedHelloServiceServer
}

func (s *grpcServer) GetName(ctx context.Context, r *pb.Request) (*pb.NameResponse, error) {
	_, resp, err := s.getName.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return &pb.NameResponse{
		Name: resp.(string),
	}, nil
}

func (s *grpcServer) GetAge(ctx context.Context, r *pb.Request) (*pb.AgeResponse, error) {
	_, resp, err := s.getAge.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return &pb.AgeResponse{
		Age: uint64(resp.(uint)),
	}, nil
}

func NewGRPCHandler(eps endpoints.HelloEndPoints) pb.HelloServiceServer {
	options := getServerOptions()
	return &grpcServer{
		getName: newServer(eps.GetNameEndpoint, options),
		getAge:  newServer(eps.GetAgeEndpoint, options),
	}
}

func newServer(e endpoint.Endpoint, options []kitGRPC.ServerOption) *kitGRPC.Server {
	return kitGRPC.NewServer(
		e,
		decodeRequest,
		encodeResponse,
		options...,
	)
}

func getServerOptions() []kitGRPC.ServerOption {
	var options []kitGRPC.ServerOption
	return options
}

func decodeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	log.Println("Request拦截-decodeRequest")
	return req, nil
}

func encodeResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	log.Println("Response拦截-encodeResponse")
	return resp, nil
}
