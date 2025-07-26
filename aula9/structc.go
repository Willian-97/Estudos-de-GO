package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

type Profissao struct {
	Pessoa
	Tipo string
}

func main() {

	fmt.Println(Pessoa{"Willian", 28})
	fmt.Println(Pessoa{Nome: "Willian", Idade: 28})
	fmt.Println(Pessoa{Nome: "Will"})

	p1 := Pessoa{"Ines", 27}
	fmt.Println(p1.Nome)
	fmt.Println(p1.Idade)
	p1.Idade = 28
	// fmt.Println(p1.Idade)
	p2 := Pessoa{"Willian Silva", 29}

	// Struct + Slice (listas)
	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	fmt.Println(pessoas)

	// Struct + map + slice

	// alunos := map[string][]Pessoa{}
	// alunos["Programação"] = pessoas
	// fmt.Println(alunos)

	var alunos = map[string][]Pessoa{
		"Programação": {{Nome: "Will", Idade: 28}, {Nome: "Willian", Idade: 29}},
		"Enfermagem":  {{Nome: "ines", Idade: 27}, {Nome: "Ines", Idade: 28}},
	}
	fmt.Println(alunos)

	// Herdando struct
	prof := Profissao{p1, "Enfermeira"}
	prof2 := Profissao{p2, "Dev"}

	fmt.Println(prof)
	fmt.Println(prof2)
	fmt.Println(prof2.Pessoa.Nome)
	fmt.Println(prof2.Pessoa.Idade)
	fmt.Println(prof2.Tipo)


}
