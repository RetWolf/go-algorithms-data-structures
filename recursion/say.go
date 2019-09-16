package recursion

import (
	"fmt"
	"strconv"
)

func Say(str string, iter int) string {
	if iter == 0 {
		return str
	}

	fmt.Println(str)
	num, count := rune(str[0]), 0
	res := ""
	for _, digit := range str {
		if digit == num {
			count++
		} else {
			res = res + strconv.Itoa(count) + string(num)
			num = digit
			count = 1
		}
	}

	res = res + strconv.Itoa(count) + string(num)
	return Say(res, iter-1)
}
