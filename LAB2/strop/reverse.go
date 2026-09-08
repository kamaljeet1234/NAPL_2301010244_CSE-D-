package strop

func Reverse(str string) string {
	reverse := ""

	for i := len(str) - 1; i >= 0; i-- {
		reverse = reverse + string(str[i])
	}

	return reverse
}
