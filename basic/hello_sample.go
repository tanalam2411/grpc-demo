package basic

import (
	"log"

	"github.com/tanalam2411/grpc-demo/protogen/basic"

)

func BasicHello() {
	h := basic.Hello{
		Name: "Max",
	}

	log.Println(&h)
}