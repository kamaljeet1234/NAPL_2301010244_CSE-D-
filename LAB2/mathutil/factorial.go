package mathutil

func Fact(a int) int {
	fact := 1

	for i := 1; i <= a; i++ {
		fact = fact * i
	}

	return fact
}

// func main() {
// 	var n int

// 	fmt.Print("Enter a number: ")
// 	fmt.Scan(&n)

// 	fmt.Println("Factorial =", Fact(n))
// }
