package main

import (
	"fmt"
	"sync"
)

func square(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- n * n
}
func main() {
	var wg sync.WaitGroup
	numbers := []int{1, 2, 3, 4, 5}
	channel := make(chan int, 5)
	for _, num := range numbers {
		wg.Add(1)
		go square(num, channel, &wg)

	}
	go func() {
		wg.Wait()
		close(channel)

	}()
	for res := range channel {
		fmt.Println(res)
	}

}
