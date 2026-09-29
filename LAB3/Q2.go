package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Marks float64
}

func modifyVariable(num *int) {
	var temp int
	fmt.Println("Enter the new value for modification")
	fmt.Scan(&temp)
	*num = temp

}

func addDataIntoStructure(s1 *Student) {
	fmt.Print("Enter name: ")
	fmt.Scan(&s1.Name)
	fmt.Print("Enter age: ")
	fmt.Scan(&s1.Age)
	fmt.Print("Enter marks: ")
	fmt.Scan(&s1.Marks)
}
func main() {
	a := 10
	p := &a
	fmt.Println("Adress of a: ", p)

	fmt.Println(*p)

	// part 2
	fmt.Println("To modify the value of a, we can use pointer")
	var input int
	fmt.Print("Enter a new value: ")
	fmt.Scan(&input)

	modifyVariable(&input)
	fmt.Println("value of variable after modification: ", input)

	// Part 3
	fmt.Println("To make the student struct ")
	s1 := new(Student)
	addDataIntoStructure(s1)
	fmt.Println(*s1)
}

