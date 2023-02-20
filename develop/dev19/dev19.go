package main

import "fmt"

//19. Разработать программу, которая переворачивает подаваемую на ход строку
//(например: «главрыба — абырвалг»). Символы могут быть unicode.

func main() {
	fmt.Println("abcdef")
	fmt.Println(reverse("abcdef"))
}

func reverse(s string) string {
	//преобразуем в массив рун
	r := []rune(s)
	l := len(r)
	for i := 0; i < l/2; i++ {
		//меняем местами первый и последний символы, второй и предпоследний и тд...
		r[i], r[l-1-i] = r[l-1-i], r[i]
	}
	return string(r)
}
