// package main

// import "fmt"

// // Собственная функция Append. Принимает слайс и произвольное кол-во элментов
// func customAppend(slice []int, elements ...int) {
// 	lenght := len(slice)                // Определим переменную как длину слайса (3)
// 	newLenght := lenght + len(elements) // Определим новую длину как сумму текущей и кол-во элементов (6)

// 	if newLenght > cap(slice) { // Если новая длина превысит емкость слайса (это так, 6  > 3)
// 		newSlice := make([]int, (newLenght+1)*2) // Создаем новый слайс через make выделив память равной двойной длине слайса
// 		copy(newSlice, slice)                    // Копируем элементы из нового слайса
// 		slice = newSlice                         // Получаем пустой слай с нужными значениями
// 	}
// 	slice = slice[:newLenght]      // Нихуя не понимаем
// 	copy(slice[lenght:], elements) // Нихуя не понимаем
// 	fmt.Println(slice)
// }

// func main(){
// 	customAppend({1,2,3}, 2)
// }