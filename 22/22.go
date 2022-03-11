package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)
	a.SetString("160000000000000000000000000000000000000000", 10)
	b.SetString("800000000000000000000000000000", 10)

	sum := new(big.Int)
	sum.Add(a, b)

	sub := new(big.Int)
	sub.Sub(a, b)

	div := new(big.Int)
	div.Quo(a, b)

	mult := new(big.Int)
	mult.Mul(a, b)

	fmt.Println("Sum is ", sum)
	fmt.Println("Sub is ", sub)
	fmt.Println("Div is ", div)
	fmt.Println("Mult is ", mult)
}
