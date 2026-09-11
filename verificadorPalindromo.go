package main

import (
	"fmt"
	"strings"
)

func main() {

	// var
	palavras := []string{
		"arara",
		"Ovo",
		"São Paulo",
		"anA",
	}

	for _, p := range palavras {
		if EhPalindromo(p) {
			fmt.Printf("\"%s\" eh um palíndromo\n", p)
		} else {
			fmt.Printf("\"%s\" não eh um palíndromo\n", p)
		}
	}
}

func EhPalindromo(texto string) bool {
	texto = strings.ToLower(texto)

	/* rune?
	Iterar sobre a string convertendo para rune garante que caracteres
	especiais e acentuados sejam tratados como um único caractere.
	*/
	var caracteres []rune
	for _, char := range texto {
		caracteres = append(caracteres, char)
	}

	inicio := 0
	fim := len(caracteres) - 1

	for inicio < fim {
		if caracteres[inicio] != caracteres[fim] {
			return false
		}
		inicio++
		fim--
	}
	return true
}
