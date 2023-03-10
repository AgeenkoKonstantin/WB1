package main

import "fmt"

//10. Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
//15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
//градусов. Последовательность в подмножноствах не важна.

func group(arr []float64) map[int][]float64 {
	// создаем мапу, в которой ключи - левая граница в группе, инты кратные 10ти
	// а значения - слайсы температурных колебаний
	result := make(map[int][]float64)
	for _, v := range arr {
		//определяем начало диапазона в группе: обрезаем дробную часть, делим на 10, умножаем на 10
		//ПРИМЕР, число 24.5 попадет в группу от 20 до 30
		var group int
		//Для чего проверка на знак v ниже
		//если v<0, допустим -5.0, то без проверки на знак было бы group=0
		//То есть 5.0 попало бы в группу от 0 до 10, что неверно
		if v > 0 {
			group = int(v) / 10 * 10
		} else {
			group = int(v)/10*10 - 10
		}

		result[group] = append(result[group], v)
	}
	return result
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 5.0, -5.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(group(temps))
}
