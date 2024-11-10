package main

type Service interface {
	Get() []byte
	Add(Student)
}
