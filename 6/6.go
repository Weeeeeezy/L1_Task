package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// остановка горутины через передачу флага остановки горутине через канал

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	extChan := make(chan bool)
	defer close(extChan)

	go worker(extChan, wg)

	time.Sleep(5 * time.Second)

	extChan <- true

	wg.Wait()
}

func worker(extChan <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-extChan:
			fmt.Println("worker is done")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println(rand.Int())
		}
	}
}
