package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 7, 9, 11}
	var i int
	fmt.Println("Print indx")
	fmt.Scanln(&i)

	nums = append(nums[:i], nums[i+1:]...)
	fmt.Println(nums)
}
