package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//18. Реализовать структуру-счетчик, которая будет инкрементироваться в
//конкурентной среде. По завершению программа должна выводить итоговое
//значение счетчика.

type Counter struct {
	value int64
}

// атомарно увеличивает счетчик
func (c *Counter) inc() {
	atomic.AddInt64(&c.value, 1)
}

func (c *Counter) print() {
	fmt.Println(atomic.LoadInt64(&c.value))
}

func main() {
	c := new(Counter)

	var wg sync.WaitGroup

	//создаем 5 горутин, в каждой увеличиваем счетчик на 500
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				c.inc()
			}
		}()
	}
	//дожидаемся выполнения горутин
	wg.Wait()
	c.print()

}
