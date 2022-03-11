package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// остановка горутины через context.WithCancel

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go worker(ctx, wg)

	time.Sleep(5 * time.Second)

	cancel()

	wg.Wait()
}

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker is done")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println(rand.Int())
		}
	}
}
