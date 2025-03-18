package grpc

import (
	"context"
	"fmt"
	"log"
	"time"
	"io"

	"github.com/tanalam2411/grpc-demo/protogen/go/hello"
	"google.golang.org/grpc"
)


func (a *GrpcAdapter) SayHello(ctx context.Context, req * hello.HelloRequest) (*hello.HelloResponse, error) {
	greet := a.helloService.GenerateHello(req.Name)

	return &hello.HelloResponse{
		Greet: greet,
	}, nil
}

func (a *GrpcAdapter) SayManyHellos(req *hello.HelloRequest, stream grpc.ServerStreamingServer[hello.HelloResponse]) error {

	for i :=0; i<10; i++ {
		greet := a.helloService.GenerateHello(req.Name)

		res := fmt.Sprintf("[%d] %s", i, greet)

		stream.Send(
			&hello.HelloResponse{
				Greet: res,
			},
		)

		time.Sleep(500 * time.Millisecond)

	}

	return nil
}


func ( a *GrpcAdapter) SayHelloToEveryOne(stream grpc.ClientStreamingServer[hello.HelloRequest, hello.HelloResponse]) error {
	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(
				&hello.HelloResponse{
					Greet: res,
				},
			)
		}

		if err != nil {
			log.Fatalln("Error while reading from client: ", err)
		}

		greet := a.helloService.GenerateHello(req.Name)

		res += greet + " "
	}
}


func (a *GrpcAdapter) SayHelloContinuous(stream grpc.BidiStreamingServer[hello.HelloRequest, hello.HelloResponse]) error {

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalln("Error while reading from client: ", err)
		}

		greet := a.helloService.GenerateHello(req.Name)

		err = stream.Send(
			&hello.HelloResponse{
				Greet: greet,
			},
		)

		if err != nil {
			log.Fatalln("Error while sending response to client: ", err)
		}
	}
}
