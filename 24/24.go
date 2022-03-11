package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPointer(xCord float64, yCord float64) *Point {
	p := &Point{
		x: xCord,
		y: yCord,
	}
	return p
}

func (p *Point) Getter() (x float64, y float64) {
	x = (*p).x
	y = (*p).y
	return x, y
}

func (p1 *Point) GetDistance(p2 *Point) float64 {
	x1, y1 := (*p1).Getter()
	x2, y2 := (*p2).Getter()
	distance := math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
	return distance
}

func main() {
	pointFirst := NewPointer(1, 1)
	pointSecond := NewPointer(2, 2)

	fmt.Println(pointFirst.Getter())
	fmt.Println(pointSecond.Getter())
	fmt.Println(pointFirst.GetDistance(pointSecond))
}
