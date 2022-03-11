package main

/*
из-за того что глобальный слайс содержит в себе указатель на срез массива,
сборщик мусора не будет чистить строку v, соответственно будут расходованы лишние ресурсы
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	tmp := make([]byte, 100)
	copy(tmp, v)
	justString = string(tmp)
}

func main() {
	someFunc()
}
