package main

import (
	"log"
	app "github.com/tanalam2411/grpc-demo/internal/application"
	grpc "github.com/tanalam2411/grpc-demo/internal/adapter/grpc"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	hs := &app.HelloService{}

	grpcAdapter := grpc.NewGrpcAdapter(hs, 9090)

	grpcAdapter.Run()
}