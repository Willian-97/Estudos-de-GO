package main

import (
	"fmt"
	"time"
)

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i;
	}
	fmt.Println(sum);

	// loop infinito
	for {
		fmt.Println("Golang do Zero");
		time.Sleep(2 * time.Second);
	}

	frutas := []string{"Laranja", "maça", "banana", "uva"};

	for i, fruta := range frutas { // podemos escoder o indice i usando _ 
		fmt.Println("Futra", fruta, ", Indice", i);
	}
}
