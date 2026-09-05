package main

import "fmt"

// Второе занятие - интерфес
// ЗАДАЧА 1
type Shape interface {
	Area() float64
}

// Circle и Rectangle обе честно реализуют интерфейс Shape.
type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}

func main() {
	rectangle1 := Rectangle{Width: 5,
		Height: 3}
	circle1 := Circle{Radius: 5}

	PrintArea(rectangle1)
	DescriptionShape(rectangle1)
	PrintArea(circle1)
	DescriptionShape(circle1)
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (r Rectangle) Area() float64 {
	return r.Width * r.Height

}
func PrintArea(s Shape) {
	fmt.Println("Площадь:", s.Area())
}
func DescriptionShape(s Shape) {
	switch v := s.(type) {
	case Circle:
		fmt.Println("Это круг с радиусом:", v.Radius)
	case Rectangle:
		fmt.Println("Это прямогугольник", v.Height, "x", v.Width)
	default:
		fmt.Println("Это неизвестная фигура")
	}

}
