package main

import "fmt"

func main() {
	// Arrays
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Println(array[0], array[1])
	fmt.Println(array)

	numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numPrimos)
	fmt.Println(numPrimos[5])
	// fatiando um array, sempre ignora o ultimo elemento 0: 4 => 2, 3, 5, 7
	fmt.Println(numPrimos[0: 4])
	// Slice
	slice := make([]string, 5)
	slice[0] = "Olá"
	slice[1] = "Mundo"
	fmt.Println(slice[0], slice[1])
	fmt.Println(slice)
	fmt.Println(len(slice))

	numPares2 := []int{2, 4, 6, 8, 10, 12}
	fmt.Println(numPares2)

	numPares2 = append(numPares2, 14, 16, 18) // só podemos usar do mesmo tipo, podemos add varios itens de uma vez


}
