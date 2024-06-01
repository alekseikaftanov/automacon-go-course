package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Задача 15.2
Используя пакет atomic, необходимо реализовать счётчик, с которым
параллельно могут работать несколько горутин
*/

func main() {
	var counter int64 // Счётчик типа int64
	var wg sync.WaitGroup

	numGoroutines := 5
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
			fmt.Printf("Горутина %d завершена\n", id)
		}(i)
	}
	wg.Wait()
}
