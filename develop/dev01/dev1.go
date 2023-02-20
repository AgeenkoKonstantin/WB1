package main

import "fmt"

//1. Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры
//Human (аналог наследования).

type Human struct {
	name string
}

func (h *Human) getName() string {
	return h.name
}

type Action struct {
	//встроили анонимное поле
	*Human
}

func main() {
	h := &Human{
		name: "geralt from rivia",
	}
	//Создали объект типа Action, полю *Human присволи значение
	action := &Action{
		Human: h,
	}
	//Появилась возможность вызывать функции *Human от переменной Action
	fmt.Println(action.getName())
}
