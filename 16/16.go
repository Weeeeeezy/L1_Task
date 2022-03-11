package main

import (
	"fmt"
	"math/rand"
)

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a) //выбираем опорный элемент

	a[pivot], a[right] = a[right], a[pivot] //меняем правый и опорный

	// все что меньше опорного скидываем влево
	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left] // ставим опорный на место после последнего элемента меньшего опорного

	//по итогу разбили массив на две части, слева меньше опорного, справа больше
	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

func main() {
	fmt.Println(quicksort([]int{5, 6, 7, 2, 1, 0}))
}
