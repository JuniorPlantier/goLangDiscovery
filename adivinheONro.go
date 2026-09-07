package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

func main() {
	nroAleatorio := rand.N(100) + 1

	fmt.Printf("[DEBUG] - nro aleatório: %d\n", nroAleatorio)

	for {
		var entrada string
		fmt.Println("--")
		fmt.Print("Entre com um nro de 1 até 100: ")
		fmt.Scan(&entrada)

		nro, err := strconv.Atoi(entrada)

		if err != nil {
			fmt.Println("Erro: É permitido apenas números inteiros.")
			continue
		}

		if nro < 1 || nro > 100 {
			fmt.Println("Erro: Números permitidos entre 1 e 100")
			continue
		}

		if nro > nroAleatorio {
			fmt.Println("O nro é maior que o sorteado. ")
		} else if nro < nroAleatorio {
			fmt.Println("O nro é menor que o sorteado ")
		} else {
			fmt.Println("Você adivinhou!!")
			fmt.Printf("O número aleatório geado foi %d", nroAleatorio)
			break
		}
	}
}
