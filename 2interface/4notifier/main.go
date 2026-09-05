package main

import "fmt"

type Notifier interface {
	Send(message string) error
}
type EmailNotifier struct {
	Email string
}
type SMSNotifier struct {
	Phone string
}

func (en EmailNotifier) Send(message string) error {
	fmt.Println("Отправка на email ", en.Email, ":", message)
	return nil

}
func (smsn SMSNotifier) Send(message string) error {
	fmt.Println("Отправка SMS на ", smsn.Phone, ":", message)
	return nil
}
func NotifyAll(notifiers []Notifier, message string) {
	for i := range notifiers {
		notifiers[i].Send(message)
	}

}
func main() {
	no1 := EmailNotifier{Email: "miles.ru"}
	no2 := SMSNotifier{Phone: "88005553535"}
	notifiers := make([]Notifier, 0, 2)
	notifiers = append(notifiers, no1)
	notifiers = append(notifiers, no2)
	NotifyAll(notifiers, "приветиикииии")
}
