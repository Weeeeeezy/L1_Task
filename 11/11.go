package main

import "fmt"

type void struct{}

func main() {
	var null void

	firstSet := make(map[int]void)
	secondSet := make(map[int]void)
	result := make(map[int]void)

	{
		firstSet[1] = null
		firstSet[0] = null
		firstSet[-1] = null
		firstSet[3] = null
		firstSet[4] = null
	}

	{
		secondSet[1] = null
		secondSet[0] = null
		secondSet[-3] = null
		secondSet[3] = null
		secondSet[7] = null
		secondSet[-6] = null
	}

	for key, _ := range firstSet {
		if _, exist := secondSet[key]; exist {
			result[key] = null
		}
	}

	for key, _ := range result {
		fmt.Print(key, " ")
	}
}
