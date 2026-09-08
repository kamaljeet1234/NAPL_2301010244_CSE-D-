package mathutil

func Pow(a, b int) int {
	result := 1

	for i := 1; i <= b; i++ {
		result = result * a
	}

	return result
}

// func main() {
// 	var a, b int

// 	fmt.Print("Enter base: ")
// 	fmt.Scan(&a)

// 	fmt.Print("Enter power: ")
// 	fmt.Scan(&b)

// 	fmt.Println("Answer =", Pow(a, b))
// }
