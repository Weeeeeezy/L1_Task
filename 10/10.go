package main

import "fmt"

func main() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	grTemp := make(map[int][]float64)

	for _, value := range temp {
		if uint(value/10) > 0 {
			grTemp[int(value/10)*10] = append(grTemp[int(value/10)*10], value)
		} else if value >= 0 {
			grTemp[0] = append(grTemp[0], value)
		} else {
			grTemp[-1] = append(grTemp[-1], value)
		}
	}

	fmt.Println(grTemp)
}
