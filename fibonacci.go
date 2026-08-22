package main

import (
	"fmt"
	"strconv"
)

func main() {
	var entrada string
	var fib int = 0
	var auxAnt int = 1

	fmt.Print("Digite um nro: ")
	fmt.Scanln(&entrada)

	// converte para inteiro ou captura um erro de conversão
	nro, err := strconv.Atoi(entrada)

	if err != nil {
		fmt.Println("Erro: Permitido apenas número")
		return
	}

	if nro < 0 {
		fmt.Println("Erro: Não é permitido número negativo")
		return
	}

	for i := 0; i < nro; i++ {

		fib = auxAnt + fib
		auxAnt = fib - auxAnt

		fmt.Println(fib)
	}

}
