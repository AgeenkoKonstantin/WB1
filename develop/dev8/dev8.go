package main

import "fmt"

//8. Дана переменная int64. Разработать программу которая устанавливает i-й бит в
//1 или 0

func changeBit(n int64, bitindex int, c bool) int64 {
	var mask int64 = 1 << bitindex
	fmt.Printf("mask : %b \n", mask)
	if c {
		fmt.Printf("number before |  : %b \n", n)
		// устанавливаем i-ый бит равный 1 через ИЛИ
		n = n | mask
		fmt.Printf("number after  |  : %b \n", n)
	} else {
		fmt.Printf("number before &^ : %b \n", n)
		//сбрасываем i-ый бит через И НЕ
		n = n &^ mask
		//можно было сделать иначе
		//mask = ^mask  побитовое НЕ для маски, то есть i-ый бит равен 0, остальные 1
		//n = n & mask через побитовое И устанавливаем i-ый бит равен 0, остальные биты не меняются
		fmt.Printf("number after  &^ : %b \n", n)
	}
	return n
}

func main() {
	var number int64 = 64
	number = changeBit(number, 2, true)
	number = changeBit(number, 2, false)
}
