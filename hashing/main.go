package main

import "fmt"

func hash(key string) int {
	fmt.Println("key is ", key)

	sum := 0
	for _, char := range key {
		sum += int(char)
	}
	return sum
}

func main() {
	key := hash("1007")
	fmt.Println(key)
}
