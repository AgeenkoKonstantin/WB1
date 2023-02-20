package main

import "fmt"

//3. Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
//квадратов с использованием конкурентных вычислений.

func main() {
	arr := []int{2, 4, 6, 8, 10}
	ch1 := make(chan int)
	sum := 0
	//записываем квадраты в канал
	for _, v := range arr {
		go func(x int) {
			ch1 <- x * x
		}(v)
	}

	for i := 0; i < len(arr); i++ {
		//читаем из канала и суммируем
		sum += <-ch1
	}
	fmt.Println(sum)
}
