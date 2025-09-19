package main

import "fmt"

type Pessoa struct {
	Nome string
	Idade int
	Profissao string
}

func (p Pessoa) ListaNomeEIdade() string {
	return fmt.Sprintf("%s tem %d anos", p.Nome, p.Idade)
}

func main() {
	pessoa := Pessoa{
		Nome: "Willian",
		Idade: 28,
		Profissao: "Dev",
	}
	segundaPessoa := Pessoa {
		Nome: "Will",
		Idade: 27,
		Profissao: "Dev Backedn",
	}

	terceiraPessoa := Pessoa {
		Nome: "Jonas",
		Idade: 25,
		Profissao: "Officeboy",
	}
	println(pessoa.ListaNomeEIdade())
	println(segundaPessoa.ListaNomeEIdade())
	println(terceiraPessoa.ListaNomeEIdade())
}