package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

//4. Реализовать постоянную запись данных в канал (главный поток). Реализовать
//набор из N воркеров, которые читают произвольные данные из канала и
//выводят в stdout. Необходима возможность выбора количества воркеров при
//старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
//способ завершения работы всех воркеров.

// Способ завершения:
// оздаем  context.WithCancel
// в отдельной горутине запускаем функцию отслеживающую сигнал os.Interrupt, если этот сигнал приходит, то вызываем cancel()
// контекс завершен, теперь в функциях записи и чтения каналов в select{} нужно отслежить это завершение контекста и выйти из функции в этом случае
func main() {

	var numWorkers int
	fmt.Println("Type number of workers: ")
	fmt.Scanf("%d\n", &numWorkers)

	ch := make(chan int, 5)

	ctx, cancel := context.WithCancel(context.Background())
	go handleSignals(cancel)

	//workers start
	for i := 0; i < numWorkers; i++ {
		go worker(ctx, ch, i)
	}

	writeToMainChan(ctx, ch)
	fmt.Println("main end")

}

func handleSignals(cancel context.CancelFunc) {
	//создаем канал для сигналов ОС
	sigChan := make(chan os.Signal)
	//говорим, чтобы сигнал os.Interrupt перехватывался и попадал в канал sigchan
	signal.Notify(sigChan, os.Interrupt)
	for {
		//если сигнал - os.Interrupt, то завершаем контекс, все горутины, в которые мы добавили отслеживание заверщения контекста также завершаться
		sig := <-sigChan
		switch sig {
		case os.Interrupt:
			cancel()
			fmt.Println("Context done")
			return
		}
	}
}

func writeToMainChan(ctx context.Context, ch chan<- int) {
	for {
		select {
		//проверяем не завершен ли контекст, если да, то выходим из функци
		case <-ctx.Done():
			fmt.Println("writeToMainChan end")
			return
			//иначе пишем в канал произвольные данные
		case ch <- rand.Intn(100):
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func worker(ctx context.Context, ch <-chan int, workerId int) {
	for {
		select {
		//проверяем не завершен ли контекст, если да, то выходим из функции
		case <-ctx.Done():
			fmt.Printf("worker: %d end\n", workerId)
			return
			//иначе выводим в stdout значение из канала
		case v := <-ch:
			fmt.Printf("worker: %d, value: %d\n", workerId, v)
		}
	}
}
