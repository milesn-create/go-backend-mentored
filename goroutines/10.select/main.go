package main

import (
	"fmt"
	"time"
)

func worker(name string, ch chan string, delay time.Duration) {
	time.Sleep(delay)
	ch <- name + "закончил работу"

}
func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	go worker("Воркер 1 ", channel1, 1*time.Second)
	go worker("Воркер 2 ", channel2, 2*time.Second)
	select {
	case message1 := <-channel1:
		fmt.Println("Сообщение от воркера 1: ", message1)
	case message2 := <-channel2:
		fmt.Println("Сообщение от воркера 2: ", message2)

	}

}
