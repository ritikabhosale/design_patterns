package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type LegacyServiceAdapter struct {
	Service LegacyService
}

func (a LegacyServiceAdapter) Add(student Student) {
	legacyStudent := LegacyStudent{
		FirstName:    student.FirstName,
		LastName:     student.LastName,
		EnrollmentNo: student.EnrollmentNo,
		AddressLine1: student.Address.AddressLine1,
		AddressLine2: student.Address.AddressLine2,
		City:         student.Address.City,
		State:        student.Address.State,
		PinCode:      student.Address.PinCode,
	}

	a.Service.addData(legacyStudent)
}

func (a LegacyServiceAdapter) Get() []byte {
	records := a.Service.getData()

	var students []Student
	for _, record := range records {
		enrollment, err := strconv.Atoi(record[2])

		if err != nil {
			print(fmt.Sprintf("failed to parse enrollment number into int: %v", err))

			continue
		}

		pinCode, err := strconv.Atoi(record[2])

		if err != nil {
			print(fmt.Sprintf("failed to parse pincode number into int: %v", err))

			continue
		}

		student := Student{
			FirstName:    record[0],
			LastName:     record[1],
			EnrollmentNo: int(enrollment),
			Address: Address{
				AddressLine1: record[3],
				AddressLine2: record[4],
				City:         record[5],
				State:        record[6],
				PinCode:      pinCode,
			},
		}
		students = append(students, student)
	}

	bytes, err := json.Marshal(students)

	if err != nil {
		print(fmt.Sprintf("failed to parse : %v", err))

		return nil
	}

	return bytes
}
