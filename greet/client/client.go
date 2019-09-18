package main

import (
	"context"
	"github.com/Go-GRPC/greet"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {

	clientCon, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to client %v", err)
	}
	defer clientCon.Close()

	c := greet.NewGreetServiceClient(clientCon)
	//unary(c)
	//streamingServer(c)
	streamingClient(c)

}

func streamingClient(client greet.GreetServiceClient) {

	var request []*greet.LongGreetRequest
	request = append(request,
		&greet.LongGreetRequest{
			Greeting: &greet.Greeting{
				FirstName: "X",
			},
		})

	request = append(request,
		&greet.LongGreetRequest{
			Greeting: &greet.Greeting{
				FirstName: "Y",
			},
		})


	request = append(request,
		&greet.LongGreetRequest{
			Greeting: &greet.Greeting{
				FirstName: "Z",
			},
		})

	stream, err := client.LongGreet(context.Background())
	if err != nil {
		log.Fatal("error while sending request", err)
	}

	for _, req := range request {
		stream.Send(req)
		time.Sleep(100*time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("error while receiving response",err)
	}
	log.Println(res.Result)
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
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal()
		}
		log.Println("response from Greet many times", msg.GetResult())
	}
}

func unary(client greet.GreetServiceClient) {
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
