package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(UniqChecker("abcd"))
	fmt.Println(UniqChecker("abCdefAaf"))
	fmt.Println(UniqChecker("aabcd"))
}

func UniqChecker(s string) bool {
	m := make(map[rune]bool)

	s = strings.ToLower(s)
	for _, c := range s {
		if m[c] {
			return false
		}

		m[c] = true
	}
	return true
}
