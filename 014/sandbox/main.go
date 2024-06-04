package main

import (
	"fmt"
)

// func sum(a int, b int, ch chan int) {
// 	ch <- a + b
// }

func makeFib(n int) []int {
	res := make([]int, n)
	res[0] = 0
	if n > 1 {
		res[1] = 1
	}
	for i := 2; i < n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	fmt.Println(res)
	return res
}

func main() {

	makeFib(10)

	// s := []int{1, 2, 3, 4, 5, 6}
	// c := 10
	// ch := make(chan int, 6)

	// wg := sync.WaitGroup{} // Объект для синхронизации горутин
	// wg.Add(len(s))         // Добавляем количество горутин, которые будут выполняться

	// go func() {

	// 	for _, value := range s {
	// 		value := value
	// 		go func() {
	// 			sum(value, c, ch)
	// 			wg.Done() // Уменьшаем счетчик горутин
	// 		}()
	// 	}
	// 	wg.Wait() // Ждем, пока все горутины не завершатся
	// 	close(ch) // Закрываем канал, когда все горутины завершились
	// }()

	// for d := range ch {
	// 	fmt.Printf("Result:%d\n", d)
	// }

}
