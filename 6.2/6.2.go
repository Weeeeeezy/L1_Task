package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// остановка горутины через context.WithDeadline

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go worker(ctx, wg)

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
