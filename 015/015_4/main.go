package main

import (
	"fmt"
	"sync"
)

/*
15.5
Необходимо создать функцию start, которая в консоль выводит некоторое сообщение.
Необходимо запустить 10 горутин, которые будут запускать функцию
start и выводить в консоль факт своего запуска, причём необходимо
обеспечить однократный запуск функции start.
*/

func start() {
	fmt.Println("Привет!")
}

func main() {

	wg := sync.WaitGroup{}
	var once sync.Once

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {

			defer wg.Done()
			once.Do(start)
			fmt.Println("Горутина №", i, "запущена")
		}(i)
	}
	wg.Wait()
}
