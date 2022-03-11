package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format(time.UnixDate))
	Sleep(3)
	fmt.Println(time.Now().Format(time.UnixDate))
}

func Sleep(s int) {
	<-time.After(time.Duration(s) * time.Second)
}
