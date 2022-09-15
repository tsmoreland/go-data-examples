package models

import "fmt"

type Person struct {
	Id        int64
	Firstname string
	Lastname  string
}

func (p Person) FullName() string {
	return fmt.Sprintf("%v %v", p.Firstname, p.Lastname)
}

func (p Person) SortingName() string {
	return fmt.Sprintf("%v, %v", p.Lastname, p.Firstname)
}
