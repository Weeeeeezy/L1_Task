package main

import "fmt"

func main() {
	str := "Hi!"
	num := 1
	b := true
	ch := make(chan interface{}, 1)

	fmt.Println(typeIdentify(str))
	fmt.Println(typeIdentify(num))
	fmt.Println(typeIdentify(b))

	ch <- str
	fmt.Println(typeIdentify(<-ch))
	ch <- num
	fmt.Println(typeIdentify(<-ch))
	ch <- b
	fmt.Println(typeIdentify(<-ch))

}

func typeIdentify(a interface{}) string {
	varType := fmt.Sprintf("%T", a)
	return varType
}
