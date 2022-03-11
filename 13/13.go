package main

import "fmt"

func main() {
	a := 5
	b := 10
	a, b = b, a
	fmt.Println(a, b)

	c := -5
	d := 5
	c = c - d
	d = c + d
	c = -c + d
	fmt.Println(c, d)

	x := 12
	y := 75
	x = (x | y) & ^(x & y)
	y = (y | x) & ^(y & x)
	x = (x | y) & ^(x & y)
	fmt.Println(x, y)
}
