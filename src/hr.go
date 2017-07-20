// filename: hr.go
package main

import "errors"

const MaxAge int = 70

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string) (*Person, error) {
	if name == "" {
		return nil, errors.New("Name is required")
	} else {
		return &Person{Name: name}, nil
	}
}
