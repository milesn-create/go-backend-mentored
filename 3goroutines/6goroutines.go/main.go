package main

import (
	"fmt"
	"sync"
)

func PrintNumber(number int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Это горутина номер :", number)
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go PrintNumber(i, &wg)

	}
	wg.Wait()
	fmt.Println("Все горутины отработали!")

}
