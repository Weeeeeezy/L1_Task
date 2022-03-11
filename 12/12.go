package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make([]string, 1)
	m := make(map[string]struct{})

	for _, word := range words {
		m[word] = struct{}{}
	}

	for word, _ := range m {
		set = append(set, word)
	}

	fmt.Println(set)
}
