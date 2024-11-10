package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type LegacyStudent struct {
	FirstName    string
	LastName     string
	EnrollmentNo int
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	PinCode      int
}

type LegacyService struct{}

func (s LegacyService) addData(student LegacyStudent) {
	info := []string{
		student.FirstName,
		student.LastName,
		strconv.Itoa(student.EnrollmentNo),
		student.AddressLine1,
		student.AddressLine2,
		student.City,
		student.State,
		strconv.Itoa(student.PinCode)}

	file, err := os.OpenFile("students.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		println(fmt.Sprintf("error while opening file : %v", err))

		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	err = writer.Write(info)

	if err != nil {
		println(fmt.Sprintf("error while writing data : %v", err))
	}
}

func (s LegacyService) getData() [][]string {
	file, err := os.Open("students.csv")

	if err != nil {
		println(fmt.Sprintf("error while opening file : %v", err))

		return nil
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		println(fmt.Sprintf("error while reading file : %v", err))

		return nil
	}

	return records
}
