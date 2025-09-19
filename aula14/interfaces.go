package main

import "fmt"

type Pessoa struct {
	Nome string
	Idade int
	Prof string
	TempoDeProfissao int
}

func (p Pessoa) FalaBomDia() string {
	return fmt.Sprintf("%s deseja bom dia para você", p.Nome)
}

func (p Pessoa) Profissao() string {
	return fmt.Sprintf("%s tem %d anos como %s", p.Nome, p.TempoDeProfissao, p.Prof)
}

type Adulto interface {
	FalaBomDia() string
	Profissao() string
}

func BomDia(adulto Adulto) string {
	return adulto.FalaBomDia()
}

func ProfissaoDaPessoa(adulto Adulto) string {
	return adulto.Profissao()
} 

func main() {
	adulto := Pessoa {
		Nome: "Willian",
		Idade: 28,
		Prof: "Dev",
		TempoDeProfissao: 2,
	}
	fmt.Println(BomDia(adulto))
	fmt.Println(ProfissaoDaPessoa(adulto))
}