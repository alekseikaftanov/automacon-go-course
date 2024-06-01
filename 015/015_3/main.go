package main

import (
	"fmt"
	"sync"
	"time"
)

/*
15.3
Используя Mutex, необходимо реализовать счётчик,
с которым параллельно могут работать несколько горутин
*/

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Decrement() {
	c.mu.Lock()
	c.count--
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {

	counter := Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
			time.Sleep(time.Millisecond * 1000)
		}()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Decrement()
			time.Sleep(time.Millisecond * 1000)
		}()
	}

	wg.Wait()
	fmt.Printf("Counter value: %d\n", counter.Value())

}
