package main

import (
	"fmt"
	"strings"
)

//20. Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func main() {
	fmt.Println("snow dog sun")
	fmt.Println(reverse("snow dog sun"))
}

func reverse(s string) string {
	//получаем срез слов
	words := strings.Fields(s)
	l := len(words)
	//меняем местами аналогично прошлой задаче
	for i := 0; i < l/2; i++ {
		words[i], words[l-i-1] = words[l-i-1], words[i]
	}
	//соединяем обратно в единую строку
	return strings.Join(words, " ")
}
