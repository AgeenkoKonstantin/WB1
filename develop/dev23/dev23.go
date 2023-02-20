package main

import "fmt"

//23. Удалить i-ый элемент из слайса.

// сохраняет порядок, работает медленнее
func del(i int) {
	users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	fmt.Println(users)
	users = append(users[:i], users[i+1:]...) // формируем новый срез  [1] ... [i-1]  [i+1] ... [len-1]
	fmt.Println(users)
}

// не сохраняет порядок, работает быстрее
func del1(i int) {
	users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	fmt.Println(users)
	users[i] = users[len(users)-1] //заменяем i-ый элемент последним
	users = users[:len(users)-1]   // обрезаем последний элемент
	fmt.Println(users)
}

func main() {
	del(3)
	del1(3)
}
