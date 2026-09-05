package main

import (
	"fmt"
	"sync"
)

func count(word string, wg *sync.WaitGroup) {
	defer wg.Done()
	cw := len(word)
	if cw%10 == 1 {
		fmt.Printf("Слово %v содержит %d символ\n", word, cw)

	} else if 1 < cw%10 && cw%10 < 5 {
		fmt.Printf("Слово %v содержит %d символа\n", word, cw)
	} else {
		fmt.Printf("Слово %v содержит %d символов\n", word, cw)
	}
}
func main() {
	var wg sync.WaitGroup
	words := []string{"i", "go", "python", "rust", "java", "kotlin", "sbcpabdspcvspbv"}
	for _, word := range words {
		wg.Add(1)
		go count(word, &wg)

	}
	wg.Wait()
	fmt.Println("Обработка завершена")
}
