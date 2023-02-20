package main

import "fmt"

//12. Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
//собственное множество.

func stringArrToSet(arr []string) {
	fmt.Println(arr)
	//Так как в мапе уникальные ключи, то можно записывать строки из последовательности как ключи, тогда дубликатов не будет
	//В этом случае значения нам не нужны, поэтому указываем тип interface{}
	set := make(map[string]interface{})
	for _, str := range arr {
		set[str] = struct{}{}
	}
	fmt.Println(set)
}

func main() {
	arr := []string{
		"cat", "cat", "dog", "cat", "tree",
	}
	stringArrToSet(arr)
}
