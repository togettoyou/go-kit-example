package main

import (
	"context"
	"github.com/togettoyou/go-kit-example/hello/transport/grpc/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9999", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewHelloServiceClient(conn)
	log.Println(client.GetName(context.Background(), &pb.Request{}))
	log.Println(client.GetAge(context.Background(), &pb.Request{}))
}
