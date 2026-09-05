package main

import (
	"fmt"
	"time"
)

func worker(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "зайка ты такой загадочный"

}
func main() {
	channel1 := make(chan string)
	go worker(channel1)
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		select {
		case messsage := <-channel1:
			fmt.Println(messsage)
		default:
			fmt.Println("сообщений нет")
		}

	}

}
