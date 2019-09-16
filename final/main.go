package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	str := "is2 Thi1s T4est 3a"
	str = stringSort(str)
	fmt.Println(str)
}

func stringSort(s string) string {
	// check for empty input string
	if s == "" {
		return ""
	}

	res := make([]string, len(s))
	temp := ""
	order := 0
	for i, char := range s {
		if i == len(s)-1 {
			temp += string(char)
			res[order] = temp
			break
		}

		if unicode.IsSpace(char) {
			res[order] = temp
			order = 0
			temp = ""
		}

		if unicode.IsDigit(char) {
			order, _ = strconv.Atoi(string(char))
		}

		temp += string(char)
	}
	return strings.Join(res, " ")
}
