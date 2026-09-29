package main

import (
	"fmt"
)

type person struct {
	name   string
	age    int
	job    string
	salary float64
}

func (p *person) readData() {
	fmt.Print("Enter name: ")
	fmt.Scan(&p.name)
	fmt.Print("Enter age: ")
	fmt.Scan(&p.age)
	fmt.Print("Enter job: ")
	fmt.Scan(&p.job)
	fmt.Print("Enter salary: ")
	fmt.Scan(&p.salary)
}
func (p person) printData() {
	fmt.Println("\n-- Formatted person data --")
	fmt.Printf("Name: %s\n", p.name)
	fmt.Printf("Age: %d\n", p.age)
	fmt.Printf("Job: %s\n", p.job)
	fmt.Printf("Salary: %s\n", p.salary)
	fmt.Println("---------------------------")
}
func main() {
	var person1 person
	var person2 person

	fmt.Println("===Input data for person 1===")
	person1.readData()
	fmt.Println("===Input data for person 2===")
	person2.readData()
	person1.printData()
	person2.printData()
}
