package main

import (
	"fmt"
	"math/rand"
)

//15. К каким негативным последствиям может привести данный фрагмент кода, и как
//это исправить? Приведите корректный пример реализации.
//var justString string
//func someFunc() {
//v := createHugeString(1 << 10)
//justString = v[:100]
//}
//func main() {
//someFunc()
//}

//Желательно не использовать глобальную переменную juststring, а создавать ее внутри функции

//строки - это рид онли слайс байтов https://go.dev/blog/strings
// поэтому justString := v[:100] дает не то поведение, которое нам требуется, в исходном варианте
//так как в переменную juststring попадут именно байты, а не символы
// поэтому нужно работать со строкой как со слайсом рун

var chars = []rune("维基百科关于中文维基百科")

func someFunc() []rune {
	v := *createHugeString(1 << 10)
	justString := v[:100]
	return justString
}

func createHugeString(l int) *[]rune {
	b := make([]rune, l)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return &b
}

func main() {
	r := someFunc()
	fmt.Println(len(r))         // массив рун длина 100, то есть 100 символов
	fmt.Println(len(string(r))) // длина строки 300, так как эта функция показывает число байт
	fmt.Println(string(r))

}
