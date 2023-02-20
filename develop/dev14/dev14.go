package main

import (
	"fmt"
	"reflect"
)

//14. Разработать программу, которая в рантайме способна определить тип
//переменной: int, string, bool, channel из переменной типа interface{}.

func main() {

	i := interface{}(1)
	s := interface{}("string")
	b := interface{}(true)

	cc := make(chan int)
	c := interface{}(cc)

	//reflect.TypeOf - позволяет узнать тип переменной из interface{}
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
}
