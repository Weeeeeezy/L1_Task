package main

import "fmt"

func BinarySearch(arr []int, el int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := right - (right-left)/2
		if el < arr[mid] {
			right = mid - 1
		} else if el > arr[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	arr := []int{0, 1, 20, 30, 34}
	el := 20
	if i := BinarySearch(arr, el); i != -1 {
		fmt.Printf("%d\n", i)
	} else {
		fmt.Println("Can't found")
	}
}
