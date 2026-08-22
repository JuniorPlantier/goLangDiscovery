package main

import "fmt"

type Pessoa struct {
	Nome      string
	Idade     uint8
	Documento string
	Endereco  string
	CEP       string
}

var idade uint8 = 10
var documento string = "00000-00"

var (
	endereco string = "Morumbi"
	ativo    bool   = true
)

func main() {
	nome := "Lina"
	p := Pessoa{
		Nome:      "Jonathan Calleri",
		Idade:     30,
		Documento: "789.652.365-55",
		Endereco:  "Av. Jorge Joao Saad",
	}

	fmt.Println(nome, idade, documento, endereco, ativo)
	fmt.Println(p.Nome)

}
