package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	const N = 5

	wg := &sync.WaitGroup{}
	inputChan := make(chan int)
	timer := time.NewTimer(N * time.Second)

	wg.Add(1)

	go generator(inputChan, timer)
	
	go reader(inputChan, wg)

	wg.Wait()
}

func generator(inputChan chan<- int, timer *time.Timer) {
	for {
		select {

		case <-timer.C:
			fmt.Println("Time is over")
			close(inputChan)
			return
		default:
			inputChan <- rand.Int()
			time.Sleep(time.Second)
		}
	}
}

func reader(inputChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range inputChan {
		fmt.Println(data)
	}
	fmt.Println("reader stop working")
}
