package main

import "fmt"

type Triangle struct {
	height float64
	base   float64
}

type Square struct {
	sideLength float64
}

type Shape interface {
	getArea() float64
}

func (t Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s Shape) {
	fmt.Println(fmt.Sprintf("The area is %f units", s.getArea()))
}

func main() {
	t := Triangle{height: 34, base: 12}
	s := Square{sideLength: 45}

	printArea(t)
	printArea(s)
}
