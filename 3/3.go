// объединение подходов из 2 решения, контролируем количество воркеров с синхронизацией через WaitGroup
package main


import (
	"fmt"
	"sync"
)

const goroutinesNum = 3

func main() {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	result := 0
	nums := [...]int{2, 4, 6, 8, 10}
	inputChan := make(chan int, 1)

	for i := 0; i < goroutinesNum; i++ {
		wg.Add(1)
		go worker(inputChan, wg, mu, &result)
	}

	for _, n := range nums {
		inputChan <- n
	}

	close(inputChan)

	fmt.Println(result)
}

func worker(inputChan <-chan int, wg *sync.WaitGroup, mu *sync.Mutex, result *int) {
	defer wg.Done()
	for n := range inputChan {
		mu.Lock() // лочим мьютекс воркером, чтобы не было гонки
		*result += n * n
		mu.Unlock()
	}
}
