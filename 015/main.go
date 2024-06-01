package main

import (
	"fmt"
	"sync"
)

// type currecnyExchange struct {
// 	mu    sync.Mutex
// 	value int
// }

// func (ce *currecnyExchange) read() int {
// 	ce.mu.Lock()
// 	defer ce.mu.Unlock()
// 	return ce.value
// }

// func (ce *currecnyExchange) add() {
// 	ce.mu.Lock()
// 	defer ce.mu.Unlock()
// 	ce.value++
// }

func main() {
	//Atomic
	// var j int32
	// for i := 0; i < 1000; i++ {
	// 	go func() {
	// 		atomic.AddInt32(&j, 1)
	// 	}()
	// }
	// time.Sleep(time.Second)
	// fmt.Println(j)

	// Mutex
	// var n int
	// var mu sync.Mutex
	// for i := 0; i < 1000; i++ {
	// 	go func() {
	// 		mu.Lock()
	// 		defer mu.Unlock()
	// 		n++
	// 	}()
	// }
	// time.Sleep(2 * time.Second)
	// mu.Lock()
	// defer mu.Unlock()
	// fmt.Println(n)

	// Exchange

	// ce := currecnyExchange{}
	// for i := 0; i < 1000; i++ {
	// 	go func() {
	// 		ce.read()
	// 	}()
	// 	go func() {
	// 		ce.add()
	// 	}()
	// }
	// time.Sleep(2 * time.Second)
	// fmt.Println(ce.read())

	// Wait Group
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		i := i // Замыкание в анонимной функции
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			fmt.Println("горутина No", i)
		}(i)
		wg.Wait()
	}

}
