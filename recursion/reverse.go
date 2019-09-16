package recursion

func ReverseRecurse(str string) string {
	if len(str) <= 1 {
		return str
	}

	return ReverseRecurse(str[1:]) + str[:1]
}

func ReverseIterative(str string) string {
	chars := []rune(str)
	for left, right := 0, len(chars)-1; left < right; left, right = left+1, right-1 {
		chars[left], chars[right] = chars[right], chars[left]
	}
	return string(chars)
}
