package main

import (
	"context"
	"github.com/Go-GRPC/greet"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {

	clientCon, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to client %v", err)
	}
	defer clientCon.Close()

	c := greet.NewGreetServiceClient(clientCon)
	//unary(c)
	streamingServer(c)

}

func streamingServer(client greet.GreetServiceClient) {
	req := &greet.GreetManyTimesRequest{
		Greeting: &greet.Greeting{
			FirstName: " kartik ",
		},
	}

	stream, err := client.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling greet ", err)
	}

	for {
		msg , err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal()
		}
		log.Println("response from Greet many times",msg.GetResult())
	}
}

func unary(client greet.GreetServiceClient){
	req := &greet.GreetRequest{
		Greeting: &greet.Greeting{
			FirstName: "Hello ",
			LastName:  "world",
		},
	}
	res, err := client.Greet(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling greet ", err)
	}
	log.Println("response from greet", res.Result)
}

