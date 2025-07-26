package main 

import "fmt"

func main() {
	idade := map[string]int{}
	idade["Willian"] = 28
	fmt.Println(idade)
	// acessando o valor especifico de uma chave
	fmt.Println(idade["Willian"])

	anoDeNasc := map[string]int{
		"Willian": 1997,
		"Bento": 2008,
	}
	fmt.Println(anoDeNasc)
	fmt.Println(anoDeNasc["Bento"])
	// add mais elementos no map
	anoDeNasc["GoLangDoZero"] = 2024
	fmt.Println(anoDeNasc)
}