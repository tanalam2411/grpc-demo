package basic

import (
	"log"

	"github.com/tanalam2411/grpc-demo/protogen/basic"
	"google.golang.org/protobuf/encoding/protojson"
)


func BasicUserGroup() {

	u1 := basic.User{
		Id: 97,
		Username: "batman",
		IsActive: true,
		Password: []byte("bat132"),
		Gender: basic.Gender_GENDER_MALE,
	}

	u2 := basic.User{
		Id: 96,
		Username: "nightwing",
		IsActive: true,
		Password: []byte("nightwing132"),
		Gender: basic.Gender_GENDER_MALE,
	}

	u3 := basic.User{
		Id: 95,
		Username: "robin",
		IsActive: true,
		Password: []byte("robin132"),
		Gender: basic.Gender_GENDER_FEMALE,
	}

	groups := basic.UserGroup{
		GroupId: 999,
		GroupName: "Alpha",
		Users: []*basic.User{&u1, &u2, &u3},
		Description: "Alpha Group",
	}

	jsonBytes, _ := protojson.Marshal(&groups)
	log.Println(string(jsonBytes))
}