package main

import "fmt"

/*
7.1 Необходимо создать массив строк размером в 5 элементов.
Результат вывести в консоль.
*/

func stringArray() {
	myArray := [5]string{"Apple", "Banana", "Cherry", "Date", "Fig"}
	fmt.Println(myArray)
}

/*
7.2 Необходимо создать массив, содержащий значения: «яблоко», «груша», «слива», «абрикос».
Результат вывести в консоль.
*/

func fruitArray() {
	myArray := [4]string{"Apple", "Pear", "Plum", "Apricot"}
	fmt.Println(myArray)
}

/*
7.3 Необходимо создать массив, содержащий значения: «яблоко», «груша», «помидор», «абрикос».
Далее «помидор» заменить на «персик». Результат вывести в консоль.
*/

func changeArray() {
	myArray := [4]string{"Apple", "Pear", "Plum", "Apricot"}
	myArray[2] = "Tomato"
	fmt.Println(myArray)
}

/*
7.4 Необходимо создать целочисленный срез со значениями: 5, 2, 8, 3, 1, 9.
Результат вывести в консоль.
*/
func makeSlice() {
	mySlice := []int{5, 2, 8, 3, 1, 9}
	fmt.Println(mySlice)
}

/*
7.5 Необходимо создать пустой срез ёмкостью 10.
Результат вывести в консоль.
*/

func makeEmptySlice(n int) {
	mySlice := make([]int, 0, n)
	fmt.Println(mySlice)
}

/*
7.6 Необходимо создать пустой срез ёмкостью 10. Далее добавить в него значения: 4, 1, 8, 9.
Результат вывести в консоль.
*/

func addToSlice() {
	mySlice := make([]int, 0, 10)
	mySlice = append(mySlice, 4, 1, 8, 9)
	fmt.Println(mySlice)
}

/*
7.7 Необходимо создать целочисленные срезы со значениями: 1, 2, 3 и 4, 5, 6.
Далее объединить эти срезы в результирующий срез.
Результирующий срез вывести в консоль.
*/

func multipleSlices() {
	mySlice1 := []int{1, 2, 3}
	mySlice2 := []int{4, 5, 6}
	resultingSlice := append(mySlice1, mySlice2...)
	fmt.Println(resultingSlice)
}

/*
7.8 Необходимо создать целочисленный срез: 1, 2, 3, 4, 5, 6.
Далее удалить элемент «4» и вывести результат в консоль.
*/

func deleteValue(n int) {

	mySlice := []int{1, 2, 3, 4, 5, 6}
	indexToRemove := -1

	for i, value := range mySlice {

		if value == n {
			indexToRemove = i
			break
		}
	}
	if indexToRemove != -1 {
		mySlice = append(mySlice[:indexToRemove], mySlice[indexToRemove+1:]...)

	}

	fmt.Println(mySlice)
}

func main() {
	stringArray()      // 7.1
	fruitArray()       // 7.2
	changeArray()      // 7.3
	makeSlice()        // 7.4
	makeEmptySlice(10) // 7.5
	addToSlice()       //7.6
	multipleSlices()   // 7.7
	deleteValue(4)     // 7.8
}
