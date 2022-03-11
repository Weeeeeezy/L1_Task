package main

import (
	"fmt"
)

func main() {
	nums := [...]int{1, 3, 5, 7, 9, 11}
	firstInput := make(chan int, 1)
	secondInput := make(chan int)

	go func() {
		for _, n := range nums {
			firstInput <- n
		}
		close(firstInput)
	}()

	go func() {
		for n := range firstInput {
			secondInput <- n * n
		}
		close(secondInput)
	}()

	for num := range secondInput {
		fmt.Println(num)
	}
}
