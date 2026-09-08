package main

import (
	"LAB2/mathutil"
	"LAB2/strop"
	"fmt"
)

func main() {

	var a, b int

	fmt.Print("Enter base: ")
	fmt.Scan(&a)

	fmt.Print("Enter power: ")
	fmt.Scan(&b)

	fmt.Println("Answer =", mathutil.Pow(a, b))
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	fmt.Println("Factorial =", mathutil.Fact(n))
	var str string

	fmt.Print("Enter a string: ")
	fmt.Scan(&str)
	fmt.Println("Reverse of '", str, "' is:", strop.Reverse(str))
	fmt.Println("Number of vowels in '", str, "' is:", strop.CountVowels(str))
	// fmt.Println("Factorial of 5 is:", mathutil.Fact(5))
	// fmt.Println("Power of 2 raised to 3 is:", mathutil.Pow(2, 3))

	// fmt.Println("Reverse of 'hello' is:", strop.Reverse("hello"))

	// fmt.Println("Number of vowels in 'hello' is:", strop.CountVowels("hello"))
}
