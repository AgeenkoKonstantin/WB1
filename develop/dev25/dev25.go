package main

import (
	"fmt"
	"time"
)

// time.after - возвращает канал, в который через указанное время будет отправлено текущее время
// поэтому в строке 11 до появления значения в канале будет блокировка, то есть горутина уснет на указанное время
func mysleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	mysleep(time.Second)
	fmt.Println("After 1 sec")
}
