package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var N int
	fmt.Println("Print number of goroutines")
	fmt.Scanf("%d", &N)

	inputChan := make(chan interface{})
	defer close(inputChan) //чтобы канал не висел открытым после интерапта

	for i := 0; i < N; i++ {
		go worker(inputChan, i)
	}

	generateNums(inputChan)
}

func worker(inputChan <-chan interface{}, i int) {
	for obj := range inputChan {
		fmt.Println(obj)
	}
}

func generateNums(inputChan chan<- interface{}) {
	for i := 0; ; i++ {
		inputChan <- rand.Int()
		time.Sleep(time.Second)
	}
}
