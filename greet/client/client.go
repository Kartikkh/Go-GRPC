package main

import (
	"context"
	"github.com/Go-GRPC/greet"
	"google.golang.org/grpc"
	"log"
)

func main() {

	clientCon, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to client %v", err)
	}
	defer clientCon.Close()

	c := greet.NewGreetServiceClient(clientCon)

	req := &greet.GreetRequest{
		Greeting: &greet.Greeting{
			FirstName: "Hello ",
			LastName:  "world",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling greet ", err)
	}
	log.Println("response from greet", res.Result)
}
