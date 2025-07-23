package main 

import "fmt"

func main() {
	somaDosValores := soma(50, 10) // salvando o retorno em uma var
	fmt.Println(soma(10, 25)) // printando o retorno
	fmt.Println(somaDosValores)

	sub := subtracao(50, 25)
	fmt.Println(sub)

	fmt.Println(nome("Willian"))
	nome1, idade := printNomeEIdade("Will", 28)
	fmt.Println(nome1, idade)

}
// função do tipo int que retorna a soma de dois numeros do tipo int
func soma(x int, y int) int {
	return x + y
}

func subtracao(x int, y int) int {
	return x - y
}

func nome(n string) string {
	return  "Olá " + n 
}
// função com dois tipos de retornos
func printNomeEIdade(nome string, idade int) (string, int) {
	return nome, idade
}