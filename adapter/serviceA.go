package main

import (
	"encoding/json"
	"fmt"
)

type serviceA struct {
	students []Student
}

func (s *serviceA) Add(student Student) {
	s.students = append(s.students, student)
}

func (s serviceA) Get() []byte {
	bytes, err := json.Marshal(s.students)

	if err != nil {
		println(fmt.Sprintf("failed to marshal %v", err))

		return nil
	}

	return bytes
}
