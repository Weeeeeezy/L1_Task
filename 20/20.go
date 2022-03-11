package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"

	fmt.Println(Reverse(str))
}

func Reverse(str string) string {
	words := strings.Split(str, " ")
	left := 0
	right := len(words) - 1

	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	return strings.Join(words, " ")
}
