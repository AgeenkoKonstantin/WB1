package main

import "fmt"

//21. Реализовать паттерн «адаптер» на любом примере.

//Адаптер - позволяет выдать недоступному для модификации классу необходимый интерфейс

// Недоступный класс
type inaccessibleClass struct {
	msg string
}

// медот недоступного класса
func (c *inaccessibleClass) sendMessage() {
	fmt.Println(c.msg)
}

// Сам адаптер, он реализовывает интерфейс, который мы хотим выдать недоступному классу
type adapter struct {
	*inaccessibleClass
}

// В методе нужного интерфейса мы вызываем метод недоступного класса
func (a *adapter) send() {
	a.sendMessage()
}

// Нужный интерфейс
type Sender interface {
	send()
}

func main() {
	adapt := &adapter{
		inaccessibleClass: &inaccessibleClass{
			msg: "im inaccessibleClass, but impelements Sender interface",
		},
	}
	//Пример использования, можем вызвать метод inaccessibleClass, от адаптера, который реализовывает интерфейс Sender
	var send Sender = adapt
	send.send()
}
