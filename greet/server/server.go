package main

import (
	"context"
	"github.com/Go-GRPC/greet"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error) {
	firstName := req.GetGreeting().FirstName
	res := &greet.GreetResponse{
		Result: firstName,
	}
	return res, nil
}

func (s *server) GreetManyTimes(req *greet.GreetManyTimesRequest, stream greet.GreetService_GreetManyTimesServer) error {
	firstName := req.Greeting.GetFirstName()
	for i := 0; i < 10; i++ {
		result := "hello: " + firstName + "number " + strconv.Itoa(i)
		res := &greet.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(100 * time.Millisecond)
	}
	return nil
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
