package main

import "fmt"

func square(n int, ch chan int) {
	ch <- n * n
}
func main() {
	var n int
	fmt.Println("Введите целое число:")
	fmt.Scan(&n)
	channel := make(chan int)
	go square(n, channel)
	result := <-channel
	fmt.Println("Квадрат числа = ", result)

}
