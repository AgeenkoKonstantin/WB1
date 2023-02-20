package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

//5. Разработать программу, которая будет последовательно отправлять значения в
//канал, а с другой стороны канала — читать. По истечению N секунд программа
//должна завершаться.

func main() {
	ch := make(chan int)
	//по истечению таймаута контекст завершается, запись в канал завершается, программа завершается
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func(cont context.Context) {
		for {
			select {
			//пишем в канал случайное число
			case ch <- rand.Intn(100):
				time.Sleep(50 * time.Millisecond)
			// если контекс завершен, то есть таймаут вышел, то закрываем канал и выходим из функции
			case <-cont.Done():
				close(ch)
				fmt.Println("Context done, channel closed")
				return
			}
		}
	}(ctx)

	// в главной горутине читаем из канала, пока он не закрыт
	for v := range ch {
		fmt.Println(v)
	}

}
