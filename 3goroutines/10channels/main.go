package main

import "fmt"

func worker(ch chan string) {
	ch <- "Привет из горутины!"
}
func main() {
	channel := make(chan string)
	go worker(channel)
	message := <-channel
	fmt.Println(message)

}

//ЗДесть не нужен waitgroup, потомучто канал сам синхронизируют работу горутин, main() естественным образом ждет пока горутина не пришлет сообщение
