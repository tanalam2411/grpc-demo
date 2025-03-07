package basic

import (
	"log"

	"github.com/tanalam2411/grpc-demo/protogen/basic"
	"google.golang.org/protobuf/encoding/protojson"
)

func BasicWriteUserContentV1() {
	uc := basic.UserContent{
		UserContentId: 1001,
		Slug: "/this-is-v1",
		// Title: "10 Strongest People",
		HtmlContent: "<p>Dummy para</p>",
		// AuthorId: 99,
	}
	WriteProtoToFile(&uc, "user_content_v1.bin")
}

func BasicReadUserContentV1() {
	log.Println("Read V1")

	uc := basic.UserContent{}
	ReadProtoFromFile("user_content_v1.bin", &uc)

	log.Println(&uc)

	ucJson, _ := protojson.Marshal(&uc)
	log.Println(string(ucJson))
	log.Println()
}

func BasicWriteUserContentV2() {
	uc := basic.UserContent{
		UserContentId: 1002,
		Slug: "/this-is-v2",
		// Title: "10 Strongest People",
		HtmlContent: "<p>Dummy para</p>",
		// AuthorId: 99,
		Category: "News",
	}
	WriteProtoToFile(&uc, "user_content_v2.bin")
}

func BasicReadUserContentV2() {
	log.Println("Read V2")

	uc := basic.UserContent{}
	ReadProtoFromFile("user_content_v2.bin", &uc)

	log.Println(&uc)

	ucJson, _ := protojson.Marshal(&uc)
	log.Println(string(ucJson))
	log.Println()
}


func BasicWriteUserContentV3() {
	uc := basic.UserContent{
		UserContentId: 1003,
		Slug: "/this-is-v3",
		// Title: "10 Strongest People",
		HtmlContent: "<p>Dummy para</p>",
		// AuthorId: 99,
		Category: "News",
		// SubCategory: "PEOPLE",

	}
	WriteProtoToFile(&uc, "user_content_v3.bin")
}

func BasicReadUserContentV3() {
	log.Println("Read V3")

	uc := basic.UserContent{}
	ReadProtoFromFile("user_content_v3.bin", &uc)

	log.Println(&uc)

	ucJson, _ := protojson.Marshal(&uc)
	log.Println(string(ucJson))
	log.Println()
}