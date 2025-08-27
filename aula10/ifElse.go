package main

import "fmt"

func main() {
	numero := 10;

	// if(numero == 1) {
	// 	fmt.Println("O Valor é igual a 1");
	// } else {
	// 	fmt.Println("O valor não é igual a 1");
	// }

	if numero == 1 {
		fmt.Println("O numero é igual a 1")
	} else if numero == 2 {
		fmt.Println("O numero é igual a 2")
	} else {
		fmt.Println("O numero não é 1 ou 2")
	}

	if numero%2 == 0 {
		fmt.Printf("%d é par", numero);
	} else {
		fmt.Printf("%d impar", numero);
	}
}
