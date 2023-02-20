package main

import (
	"fmt"
	"strings"
)

//26. Разработать программу, которая проверяет, что все символы в строке
//уникальные (true — если уникальные, false etc). Функция проверки должна быть
//регистронезависимой.
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

func isCharsUnique(str string) bool {
	//переводим в нижний регистр
	s := strings.ToLower(str)
	//создаем мапу, которая в качестве ключей хранит символы
	m := make(map[rune]struct{})

	// каждый символ строки пробуем вытащить из мапы, если получилось, то данный символ не уникальный
	//иначе записываем его в мапу и проверяем дальше
	for _, r := range s {
		if _, ok := m[r]; !ok {
			m[r] = struct{}{}
		} else {
			fmt.Printf("%s - contains not unique chars \n", str)
			return false
		}
	}
	fmt.Printf("%s - contains only unique chars \n", str)
	return true
}

func main() {
	isCharsUnique("abcd")
	isCharsUnique("abCdefAaf")
	isCharsUnique("aabcd")
}
