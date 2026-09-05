package main

import "fmt"

// Задача 1. Напиши функцию, которая принимает срез целых чисел и возвращает новый срез, где каждый элемент — квадрат исходного.
func outputResult(slice_1 []int) []int {
	slice_2 := make([]int, 0, len(slice_1))
	for _, v := range slice_1 {
		slice_2 = append(slice_2, v*v)
		//тут обяхательномы должны прировнять так как иначе append создает новый массив и мы потеряем данные

	}
	return slice_2

}

// Задача 2. Напиши структуру Person с полями Name и Age, и метод IsAdult(), который возвращает true, если возраст ≥ 18.
type Person struct {
	Name string
	Age  int
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// Задача 3
func main() {
	m := make(map[string]int)
	m["a"] = 1
	modify(m)
	fmt.Println(m["a"])
}
func modify(m map[string]int) {
	m["a"] = 100
}
