package main

import "fmt"

func main() {
	// Varivael
	var nome string
	nome = "Bento"
	fmt.Println(nome)

	nome = "Will"
	fmt.Println(nome)
	// Outra forma de declarar variaveis 	
	var idade int
	idade = 28
	fmt.Println(idade)
	// inferencia de tipo (Go entende o tipo)
	var a = "Willian"
	fmt.Println(a)
	// decalrando varias variaveis com o mesmo tipo
	var b, c int = 10, 5
	fmt.Println(b, c)

	// Forma mais usada, usando inferencia
	meuNome := "Willian"
	fmt.Println(meuNome)
	numero := 10
	fmt.Println(numero)

	// Constante
	const minhaIdade = 28
	fmt.Println(minhaIdade)
	const estudou = true
	fmt.Println(estudou)
}
