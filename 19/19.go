package main

import "fmt"

func main() {
	str := "главрыба"

	fmt.Println(Reverse(str))
}

func Reverse(str string) string {
	ch := []rune(str)
	left := 0
	right := len(ch) - 1

	for left < right {
		ch[left], ch[right] = ch[right], ch[left]
		left++
		right--
	}

	return string(ch)
}
