package main

import (
	"fmt"
	"sync"
	"time"
)

func slowTask(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)

	fmt.Println("Задача - ", n, "Завершена")

}
func main() {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go slowTask(i, &wg)

	}
	wg.Wait()
	elapsed := time.Since(start)

	fmt.Println("Всего времени:", elapsed)

}
