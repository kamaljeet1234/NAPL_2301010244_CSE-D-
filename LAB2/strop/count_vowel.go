package strop

func CountVowels(str string) int {
	count := 0

	for i := 0; i < len(str); i++ {
		if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' || str[i] == 'u' {
			count++
		}
	}
	return count
}
