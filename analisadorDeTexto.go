package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Analisa a contagem de vogais e consoantes
func contarVogaisEConsoantes(texto string) (int, int) {
	vogais := 0
	consoantes := 0

	for _, char := range strings.ToLower(texto) {
		if unicode.IsLetter(char) {
			switch char {
			case 'a', 'e', 'i', 'o', 'u', 'á', 'é', 'í', 'ó', 'ú', 'â', 'ê', 'ô', 'ã', 'õ':
				vogais++
			default:
				consoantes++
			}
		}
	}
	return vogais, consoantes
}

// Mapeia a frequência de cada letra
func calcularFrequencia(texto string) map[rune]int {
	frequencia := make(map[rune]int)

	for _, char := range strings.ToLower(texto) {
		if unicode.IsLetter(char) {
			frequencia[char]++
		}
	}
	return frequencia
}

// Verifica se o texto é um palíndromo
func ehPalindromo(texto string) bool {
	var caracteres []rune

	// Extrai apenas letras e números
	for _, char := range strings.ToLower(texto) {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			caracteres = append(caracteres, char)
		}
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite o texto para análise: ")
	if !scanner.Scan() {
		fmt.Println("Erro ao ler entrada.")
		return
	}

	texto := scanner.Text()

	if strings.TrimSpace(texto) == "" {
		fmt.Println("Texto vazio. Tente novamente.")
		return
	}

	// 1. Processamento básico
	totalCaracteres := len([]rune(texto)) // Conta caracteres reais (runes)
	palavras := strings.Fields(texto)     // Divide o texto por espaços em branco
	totalPalavras := len(palavras)

	// 2. Análise de vogais e consoantes
	totalVogais, totalConsoantes := contarVogaisEConsoantes(texto)

	// 3. Frequência de letras
	frequenciaLetras := calcularFrequencia(texto)

	// 4. Teste de Palíndromo
	palindromo := ehPalindromo(texto)

	// ================= RELATÓRIO =================
	fmt.Println("\n================ RELATÓRIO ================")
	fmt.Printf("• Total de caracteres (com espaços): %d\n", totalCaracteres)
	fmt.Printf("• Total de palavras: %d\n", totalPalavras)
	fmt.Printf("• Vogais: %d | Consoantes: %d\n", totalVogais, totalConsoantes)

	if palindromo {
		fmt.Println("• É palíndromo? Sim! 🎉")
	} else {
		fmt.Println("• É palíndromo? Não.")
	}

	fmt.Println("\nFREQUÊNCIA DE LETRAS:")
	for letra, qtd := range frequenciaLetras {
		fmt.Printf("  '%c': %dx\n", letra, qtd)
	}
}
