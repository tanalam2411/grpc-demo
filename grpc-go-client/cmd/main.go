package main

import (
	"context"
	"log"

	"github.com/tanalam2411/grpc-demo/internal/adapter/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("localhost:9090", opts...)

	if err != nil{
		log.Fatalln("Can not connect to gRPC server: ", err)
	}

	defer conn.Close()
 
	helloAdapter, err := hello.NewHelloAdapter(conn)

	if err != nil {
		log.Fatalln("Can not create HelloAdapter: ", err)
	}

	// runSayHello(helloAdapter, "Bruce  W.")
	// runSayManyHellos(helloAdapter, "Stream msg ...")
	// runSayHelloToEveryOne(helloAdapter, []string{"Tan", "San", "Nab", "Man"})
	runSayHelloContinuous(helloAdapter, []string{"Tan", "San", "Nab", "Man"})

}


func runSayHello(adapter *hello.HelloAdapter, name string) {
	greet, err := adapter.SayHello(context.Background(), name)

	if err != nil {
		log.Fatalln("Can not call SayHello: ", err)
	}

	log.Println(greet.Greet)
}

func runSayManyHellos(adapter *hello.HelloAdapter, name string) {
	adapter.SayManyHellos(context.Background(), name)
}

func runSayHelloToEveryOne(adapter *hello.HelloAdapter, names []string) {
	adapter.SayHelloToEveryOne(context.Background(), names)
}

func runSayHelloContinuous(adapter *hello.HelloAdapter, names []string) {
	adapter.SayHelloContinuous(context.Background(), names)
}