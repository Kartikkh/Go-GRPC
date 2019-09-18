package main

import (
	"context"
	"github.com/Go-GRPC/greet"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error) {
	firstName := req.GetGreeting().FirstName
	res := &greet.GreetResponse{
		Result: firstName,
	}
	return res, nil
}

func main() {

	conn, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	s := grpc.NewServer()
	greet.RegisterGreetServiceServer(s, &server{})
	err = s.Serve(conn)
	if err != nil {
		log.Fatalf("failed to start GRPC server : %v", err)
	}
}
