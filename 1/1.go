package main

import "fmt"

type Human struct {
	name string
	age  uint
	city string
}

func (h Human) getName() string {
	return h.name
}

func (h *Human) changeName(name string) {
	h.name = name
}

type Action struct {
	Human
}

func main() {
	a := Action{
		Human{
			name: "Mansur",
			age:  47,
			city: "Moscow"},
	}

	fmt.Println(a.getName())
	a.changeName("Raphael")
	fmt.Println(a.getName())
}
