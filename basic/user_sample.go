package basic

import (
	"log"
	"math/rand"

	"github.com/tanalam2411/grpc-demo/protogen/basic"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/genproto/googleapis/type/date"

)


func BasicUser() {

	addr := basic.Address{
		Street: "Dapodi",
		City: "Pune",
		Country: "India",
		Coordinate: &basic.Address_Coordinate{
			Latitude: 40.6848588484858484848,
			Longitude: -74.58384857738374758,			
		},
	}

	comm := randomCommunicationChannel()
	sr := map[string]uint32{"fly": 5, "speed": 5, "durability": 4}

	u := basic.User{
		Id: 99,
		Username: "user-X",
		IsActive: true,
		Password: []byte("ahsh34h3jd!@3"),
		Emails: []string{"user-x@gmail.com", "user-x@hotmail.com"},
		Gender: basic.Gender_GENDER_FEMALE,
		Address: &addr,
		CommunicationChannel: &comm,
		SkillRating: sr,
		LastLoginTimestamp: timestamppb.Now(),
		BirthDate: &date.Date{Year: 2000, Month: 5, Day: 27},

	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}

func ProtoToJsonUser() {
	p := basic.User{
		Id: 98,
		Username: "some-user",
		IsActive: true,
		Password: []byte("abcdef123"),
		Emails: []string{"user-x@gmail.com", "user-x@hotmail.com"},
		Gender: basic.Gender_GENDER_FEMALE,
	}

	jsonBytes, _ := protojson.Marshal(&p)
	log.Println(string(jsonBytes))
}

func JsonToProtoUser() {
	json := `
	{
	"id": 98,
	"username": "some-user",
	"isactive": true,
	"password": "YWJjZGVmMTIz",
	"emails": [
		"user-x@gmail.com",
		"user-x@hotmail.com"
	],
	"gender": "GENDER_FEMALE"
	}
	`

	var p basic.User

	err := protojson.Unmarshal([]byte(json), &p)

	if err != nil {
		log.Fatalln("Got error: ", err)
	}

	log.Println(&p)
}

func randomCommunicationChannel() anypb.Any {

	paperMail := basic.PaperMail{
		PaperMailAddress: "09 3 B, kon 3325",
	} 

	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "bTm",
		SocialMediaUsername: "ton.btm",
	}

	instantMessaging := basic.InstantMessaging{
		InstantMessagingProduct: "p1",
		InstantMessagingUsername: "p1.btm",
	}

	var a anypb.Any

	switch r := rand.Intn(10) % 3; r {
	case 0:
		anypb.MarshalFrom(&a, &paperMail, proto.MarshalOptions{})
	case 1:
		anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})
	default:
		anypb.MarshalFrom(&a, &instantMessaging, proto.MarshalOptions{})
	}

	return a
}

func BasicUnmarshalAnyKnown() {
	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "Platform-1",
		SocialMediaUsername: "User-1",
	}

	var a anypb.Any
	anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})

	// known type (Social Media)
	sm := basic.SocialMedia{}

	if err := a.UnmarshalTo(&sm); err != nil {
		return
	}

	jsonBytes, _ := protojson.Marshal(&sm)
	log.Println(string(jsonBytes))
}


func BasicUnmarshalAnyNotKnown() {
	a := randomCommunicationChannel()

	var unmarshaled protoreflect.ProtoMessage

	unmarshaled, err := a.UnmarshalNew()

	if err != nil {
		return
	}

	log.Println("Unmarshal as", unmarshaled.ProtoReflect().Descriptor().Name())

	jsonBytes, _ := protojson.Marshal(unmarshaled)
	log.Println(string(jsonBytes))
}

func BasicUnmarshalAnyIs() {
	a := randomCommunicationChannel()
	pm := basic.PaperMail{}

	if a.MessageIs(&pm) {
		if err := a.UnmarshalTo(&pm); err != nil {
			log.Fatalln(err)
		}

		jsonBytes, _ := protojson.Marshal(&pm)
		log.Println(string(jsonBytes))
	} else {
		log.Println("Not PaperMail, but: ", a.TypeUrl)
	}
}

func BasicOneOf() {
	// socialMedia := basic.SocialMedia{
	// 	SocialMediaPlatform: "tiktok",
	// 	SocialMediaUsername: "utube",
	// }

	// ecomm := basic.User_SocialMedia{
	// 	SocialMedia:  &socialMedia,
	// }

	instantMessaging := basic.InstantMessaging{
		InstantMessagingProduct: "wtsapp",
		InstantMessagingUsername: "u1",
	}

	ecomm := basic.User_InstantMessaging{
		InstantMessaging: &instantMessaging,
	}

	u := basic.User{
		Id: 96,
		Username: "aquaman",
		IsActive: true,
		ElectronicCommChannel: &ecomm,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}
 