package main

import "fmt"

// 11. Реализовать пересечение двух неупорядоченных множеств.

// Записываем в  мапу ключи из множеств при этом увеличиваем значение при данном ключе, значение - счетчик вхождений
func intersection(set1, set2 map[string]struct{}) (result []string) {
	midMap := make(map[string]int) // хранит ключ : число_вхождений
	for key, _ := range set1 {     // добавляем ключи из первого сета
		midMap[key] = 1
	}
	for key, _ := range set2 { // добавляем ключи из второго сета
		midMap[key]++
	}
	// Если вхождения 2, то данных ключ есть  в обоих множествах, в ином случае удаляем его
	for key, val := range midMap {
		if val != 2 {
			delete(midMap, key)
		} else {
			result = append(result, key)
		}
	}
	return
}

func main() {
	var set1 = map[string]struct{}{
		"Tom":   struct{}{},
		"Sam":   struct{}{},
		"Alice": struct{}{},
	}

	var set2 = map[string]struct{}{
		"Tom": struct{}{},
		"Bob": struct{}{},
		"Sam": struct{}{},
	}

	fmt.Println(intersection(set1, set2))

}
