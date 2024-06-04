package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Задача 16.2 Нужнозапустить 5 горутини о становить через 2 секунды.

func main() {
	// Создаем контекст с тайм-аутом в 2 секунды
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Инициализируем WaitGroup
	var wg sync.WaitGroup

	// Запускаем 5 горутин
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("goroutine %d is stopped\n", i)
					return
				default:
					// Симуляция работы горутины
					fmt.Printf("goroutine %d is running\n", i)
					time.Sleep(500 * time.Millisecond)
				}
			}
		}(i)
	}

	// Ожидание завершения всех горутин
	wg.Wait()
	fmt.Println("All goroutines are stopped")
}
