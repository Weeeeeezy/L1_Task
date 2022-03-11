package main

import (
	"fmt"
)

func main() {
	var num int64
	var i uint8
	var sgn int

	fmt.Println("Print number and index")
	fmt.Scan(&num, &i)

	if num < 0 {
		sgn = -1
	} else {
		sgn = 1
	}

	num = int64(sgn) * num

	if num&(1<<i) > 0 {
		num = num & ^(1 << i)
	} else {
		num = num | (1 << i)
	}

	num = int64(sgn) * num
	fmt.Println(num)
}
