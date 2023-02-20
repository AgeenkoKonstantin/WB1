package main

import (
	"fmt"
	"math/big"
)

//22. Разработать программу, которая перемножает, делит, складывает, вычитает две
//числовых переменных a,b, значение которых > 2^20.

//2^20 = 1 048576

func main() {
	// работаем с пакетом math/big
	//создаем переменные типа биг инт
	var a = new(big.Int)
	var b = new(big.Int)
	//сетим значения, указывая систему счисления
	a.SetString("3333333333333333333333", 10)
	b.SetString("2222222222223333333", 10)
	//для каждой операции создаем свой объект биг инт, потом у него вызываем нужный метод
	fmt.Printf("mul: %v \n", new(big.Int).Mul(a, b))
	fmt.Printf("div: %v \n", new(big.Int).Div(a, b))
	fmt.Printf("add: %v \n", new(big.Int).Add(a, b))
	fmt.Printf("sub: %v \n", new(big.Int).Sub(a, b))
}
