package main

import "fmt"

func zeroValue(i int) {
	i = 0 
	fmt.Println("Enderço do I dentro da função: ", &i)
}

func zeroPointer(i *int) {
	*i = 0 // mudou para 0, alterando o valor salvo no endereço
}

func main() {
	i := 1

	fmt.Println("Enderço de memoria do I: ", &i)
	a := &i

	fmt.Println("Salvando end do I em A: ", a)
	fmt.Println("Valor salvo no end do I em A: ", *a)
	fmt.Println("Endereço do A: ", &a)

	zeroValue(i) // permanece 1, endereços diferentes na memoria
	fmt.Println(i) // permanece 1, endereços diferentes na memoria
	zeroPointer(&i) // mudou para 0, alterando o valor salvo no endereço
	fmt.Println(i) // mudou para 0, alterando o valor salvo no endereço
}
