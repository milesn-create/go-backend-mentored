package main

import (
	"fmt"
	"sync"
)

func plus_1000(counter *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {

		mu.Lock()
		*counter++
		mu.Unlock()
	}

}
func main() {
	counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go plus_1000(&counter, &wg, &mu)

	}
	wg.Wait()
	fmt.Println(counter)

}
