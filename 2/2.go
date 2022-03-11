package main
//решение с помощью синхронизации потоков через WaitGroup
import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(5) //кол-во горутин
	nums := [...]int{2, 4, 6, 8, 10}

	for _, n := range nums {
		go func(n int) {
			defer wg.Done() //декрементим счетчик группы
			fmt.Println(n * n)
		}(n)
	}

	wg.Wait() //ждем обнуления счетчика, только потом завершаем main
}
