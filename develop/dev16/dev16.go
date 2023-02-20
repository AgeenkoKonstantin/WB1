package main

import "fmt"

//16. Реализовать быструю сортировку массива (quicksort) встроенными методами
//языка.

func partition(arr []int, low, high int) ([]int, int) {
	//выбираем опорный элемент
	pivot := arr[high]
	i := low
	//обходим подмассив от low до high
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			//если текущий элемент меньше опроного, то меняем местами элементы
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	// переставляем первый и последний элементы
	arr[i], arr[high] = arr[high], arr[i]
	//возвращаем массив и индекс опорного элемента
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		//опорный элемент
		var p int
		// раздеделение массива на подмассивы
		arr, p = partition(arr, low, high)
		//рекурсивно сортируем левый подмассив
		arr = quickSort(arr, low, p-1)
		//рекурсивно сортируем правый подмассив
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func main() {
	fmt.Println(quickSortStart([]int{17, 6, 5, 5, 6, 7, 2, 1, 0, 15, -10, 2, 333, -50}))
}
