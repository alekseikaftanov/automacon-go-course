package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Задача 16.3 Нужнозапустить 5 горутин и остановить в некоторое время,
которое рассчитывается по формуле: текущий момент + 2 секунды
*/

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					fmt.Printf("goroutine %d is running\n", i)
				}
			}
		}(i)
	}
}
