package car

import (
	"log"

	"github.com/google/uuid"
	"github.com/tanalam2411/grpc-demo/protogen/car"
)

func ValidateCar() {
	car := car.Car{
		CarId: uuid.New().String(),
		Brand: "Bmwx",
		Model: "745i Sport",
		Price: 204000,
		ManufactureYear: 2037,
	}

	err := car.ValidateAll()

	if err != nil {
		log.Fatalln("Validation Failed", err)
	}

	log.Println(&car)
}

// run main.go
// 04:10:11 Validation Failed invalid Car.Brand: value does not match regex pattern "(?i)^Toyota|Honda|Ford|BMW$"; invalid Car.ManufactureYear: value must be inside range [2020, 2030]
// exit status 1
// make: *** [run] Error 1