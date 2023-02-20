package main

import (
	"fmt"
	"strconv"
	"sync"
)

//Создаем свой класс, в методы чтения и записи добавляем блокировки мьютекса
// можно также использовать sync.Map

type ConcurrentMap struct {
	mutex sync.RWMutex
	m     map[string]string
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		m: make(map[string]string),
	}
}
func (c *ConcurrentMap) Get(key string) (string, bool) {
	c.mutex.RLock()
	val, ok := c.m[key]
	c.mutex.RUnlock()
	return val, ok
}

func (c *ConcurrentMap) Put(key string, value string) {
	c.mutex.Lock()
	c.m[key] = value
	c.mutex.Unlock()
}

func main() {
	mymapTest()
	fmt.Println("")
	syncmapTest()
}

func syncmapTest() {

	var sm sync.Map
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for i := 0; i < 100; i++ {
			sm.Store("G 1 "+strconv.Itoa(i), "G 1 "+strconv.Itoa(i))
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 100; i++ {
			sm.Store("G 2 "+strconv.Itoa(i), "G 2 "+strconv.Itoa(i))
		}
		wg.Done()
	}()

	wg.Wait()

	sm.Range(func(k, v interface{}) bool {
		fmt.Println(k, ":", v)
		return true
	})
}

func mymapTest() {
	mymap := NewConcurrentMap()
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for i := 0; i < 100; i++ {
			mymap.Put("G 1 "+strconv.Itoa(i), "G 1 "+strconv.Itoa(i))
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 100; i++ {
			mymap.Put("G 2 "+strconv.Itoa(i), "G 2 "+strconv.Itoa(i))
		}
		wg.Done()
	}()
	wg.Wait()
	for k, v := range mymap.m {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}
}
