package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}
type Employee struct {
	ID       int
	Position string
	Person   Person
}

func (e Employee) PrintEmployee() {
	fmt.Println(e)
}
func main() {
	person := Person{1, "Diego", "11/11/1999"}
	employee := Employee{123, "developer", person}
	employee.PrintEmployee()
}
