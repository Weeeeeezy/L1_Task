package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	sync.Mutex
}

func (c *Counter) Increase() {
	c.Lock()
	c.count++
	c.Unlock()
}

func New() *Counter {
	return &Counter{}
}

func (c *Counter) Getter() int {
	return c.count
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(5)

	c := New()

	for i := 0; i < 5; i++ {
		go func(c *Counter) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				c.Increase()
			}
		}(c)
	}

	wg.Wait()

	fmt.Println(c.Getter())
}
