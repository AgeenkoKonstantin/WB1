package main

import (
	"fmt"
	"runtime"
	"sync"
)

//9. Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
//массива, во второй — результат операции x*2, после чего данные из второго
//канала должны выводиться в stdout

// arr -> fromArrtoChan1() -> chan1 -> x2fromchan1tochan2() -> chan2 -> for_range -> stdout

func main() {
	chan1 := make(chan int, 5)
	chan2 := make(chan int, 5)
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 10}

	var wg sync.WaitGroup

	//создаем горутину,  в которой пишем из массива в канал
	go func() {
		// закрываем канал после обработки всего массива
		defer close(chan1)
		for _, v := range arr {
			chan1 <- v
		}
	}()

	//создаем горутины, которые берут числа из первого канала и во второй канал отправляют x*2
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range chan1 {
				chan2 <- v * 2
			}
		}()
	}

	//горутина отслеживающая готовность группы, по готовности закрывает канал 2
	go func() {
		wg.Wait()
		close(chan2)
	}()

	// в главной горутине печатаем в stdout
	for v := range chan2 {
		fmt.Println(v)
	}

}
