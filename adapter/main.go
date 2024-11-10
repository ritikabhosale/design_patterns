package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	student := Student{
		FirstName:    "John",
		LastName:     "Doe",
		EnrollmentNo: 12345,
		Address: Address{
			AddressLine1: "123 Main St",
			AddressLine2: "Apt 4B",
			City:         "Springfield",
			State:        "IL",
			PinCode:      62701,
		},
	}

	legacy := LegacyService{}
	adpater := LegacyServiceAdapter{Service: legacy}

	service := serviceA{}
	service.Add(student)

	adpater.Add(student)
	services := []Service{&service, adpater}

	for _, service := range services {
		students := make([]Student, 0)
		err := json.Unmarshal(service.Get(), &students)

		if err != nil {
			println(fmt.Sprintf("failed to unmarshal : %v", err))

			continue
		}

		for _, student := range students {
			println(student.ToString())
		}
		// can use students in further implementation
	}
}
