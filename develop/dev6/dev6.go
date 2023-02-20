package main

import (
	"context"
	"fmt"
	"time"
)

// завершение горутины через какие-либо контексты
// context.WithCancel  - я использовал в 4ом задании
// context.WithTimeout - в 5ом
// ниже пример context.WithDeadline - тоже самое, что и context.WithTimeout
func contextStop() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(100*time.Millisecond))
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("gorutina: contextStop end")
				return
			}
		}
	}()
}

// через  отправку значения в канал, Periodic polling channel
func chanStop() {
	ch := make(chan int)
	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("gorutina: chanStop end")
				return
			}
		}
	}()

	time.Sleep(time.Second * 2)
	ch <- 1
}

// через закрытие канала
func closeChanStop() {
	ch := make(chan int)
	go func() {
		//можно заменить на for range
		for {
			v, ok := <-ch
			if !ok {
				fmt.Println("gorutina: closeChanStop end")
				return
			}
			fmt.Println(v)
		}
	}()
	time.Sleep(time.Second)
	close(ch)

}

func main() {
	contextStop()
	chanStop()
	closeChanStop()
	time.Sleep(time.Second)
}
