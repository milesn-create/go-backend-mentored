package main

import "fmt"

func square(n int, ch chan int) {
	ch <- n * n
}
func main() {
	channel := make(chan int)
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		go square(num, channel)
	}
	for i := 0; i < len(numbers); i++ {
		result := <-channel
		fmt.Println(result)
	}
}
