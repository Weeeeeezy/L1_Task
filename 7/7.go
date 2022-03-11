package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	goroutinesNum := 5

	wg := &sync.WaitGroup{}
	wg.Add(goroutinesNum)

	mu := &sync.Mutex{}

	m := make(map[string]int)

	for i := 0; i < goroutinesNum; i++ {
		go insertToMap(&m, wg, mu, i)
	}

	wg.Wait()
	fmt.Print(m)
}

func insertToMap(mp *map[string]int, wg *sync.WaitGroup, mu *sync.Mutex, i int) {
	defer wg.Done()
	mu.Lock()
	// блокируем доступ к мапе, т.к. при добавлении элемента может поменяться хешмапа
	(*mp)[fmt.Sprintf("worker num %d", i)] = rand.Intn(10)
	mu.Unlock()
}
