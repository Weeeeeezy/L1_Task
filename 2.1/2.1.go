package main

//решение через пул воркеров с синхронизацией каналами
import (
	"fmt"
	"log"
	"runtime"
	"time"
)

const goroutinesNum = 3 //количество запускаемых воркеров

func main() {
	defer timeTrack(time.Now(), "main")
	/*
		хотел затестить изменение рантайма в зависимости от наличия Goshed
		при  наличии трекает всегда около 500 ns, при отсутствии иногда 500 иногда 0
	*/

	nums := [...]int{2, 4, 6, 8, 10}
	inputChan := make(chan int, 1) //канал который передает элементы массива в воркеры
	doneChan := make(chan bool)    //с помощью канала проверяем все ли корутины завершили работу

	for i := 0; i < goroutinesNum; i++ {
		go startWorker(inputChan, doneChan)
	}

	for _, number := range nums {
		inputChan <- number //записываем элементы в канал для передачи воркерам
	}

	close(inputChan)

	for i := 0; i < goroutinesNum; i++ {
		<-doneChan //получаем блокировку main пока все воркеры не закончат работу
	}
}

func startWorker(inputChan <-chan int, doneChan chan<- bool) {
	for num := range inputChan {
		fmt.Println(num * num)
		runtime.Gosched() // ??
	}
	doneChan <- true
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
