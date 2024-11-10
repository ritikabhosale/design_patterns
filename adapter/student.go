package main

import "strconv"

type Address struct {
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	PinCode      int
}

func (a Address) ToString() string {
	return a.AddressLine1 + ", " + a.AddressLine2 + ", " + a.City + ", " + a.State + ", " + strconv.Itoa(a.PinCode)
}

type Student struct {
	FirstName    string
	LastName     string
	EnrollmentNo int
	Address      Address
}

func (s Student) ToString() string {
	return "Name: " + s.FirstName + " " + s.LastName +
		"\nEnrollment No: " + strconv.Itoa(s.EnrollmentNo) +
		"\nAddress: " + s.Address.ToString()
}
