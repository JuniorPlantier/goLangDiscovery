package main

import (
	"bufio"
	"fmt"

	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Verifica se um número inteiro é primo
func ehPrimo(n int) bool {
	if n <= 1 {
		return false
	}
	// Testamos até a raiz quadrada de n para otimizar o algoritmo
	limite := int(math.Sqrt(float64(n)))
	for i := 2; i <= limite; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Processa estatísticas básicas em float64
func calcularEstatisticas(numeros []float64) (min, max, soma, media float64) {
	if len(numeros) == 0 {
		return 0, 0, 0, 0
	}

	min = numeros[0]
	max = numeros[0]

	for _, n := range numeros {
		soma += n
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	media = soma / float64(len(numeros))
	return min, max, soma, media
}

// Calcula a mediana (requer slice ordenado)
func calcularMediana(numeros []float64) float64 {
	n := len(numeros)
	if n == 0 {
		return 0
	}

	// Criamos uma cópia para não alterar o slice original durante a ordenação
	copia := make([]float64, n)
	copy(copia, numeros)
	slices.Sort(copia)

	// Se for par, faz a média dos dois do meio. Se for ímpar, pega o elemento central.
	if n%2 == 0 {
		return (copia[n/2-1] + copia[n/2]) / 2.0
	}
	return copia[n/2]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== ANALISADOR NUMÉRICO EM GO ===")
	fmt.Print("Digite números separados por espaço (ex: 10 -5 3.14 7 42 8): ")

	if !scanner.Scan() {
		fmt.Println("Erro ao ler entrada.")
		return
	}

	entrada := scanner.Text()
	partes := strings.Fields(entrada)

	if len(partes) == 0 {
		fmt.Println("Nenhum número foi inserido.")
		return
	}

	var numerosFloat []float64
	var numerosInt []int
	var contagemPares, contagemImpares int

	// Processamento e conversão de tipos
	for _, p := range partes {
		// Conversão para float64 (aceita decimais e inteiros)
		valFloat, errFloat := strconv.ParseFloat(p, 64)
		if errFloat != nil {
			fmt.Printf("⚠️  Aviso: '%s' não é um número válido e foi ignorado.\n", p)
			continue
		}
		numerosFloat = append(numerosFloat, valFloat)

		// Conversão para int (para verificações de par/ímpar e primo)
		valInt, errInt := strconv.Atoi(p)
		if errInt == nil {
			numerosInt = append(numerosInt, valInt)
			if valInt%2 == 0 {
				contagemPares++
			} else {
				contagemImpares++
			}
		}
	}

	if len(numerosFloat) == 0 {
		fmt.Println("Nenhum número válido foi processado.")
		return
	}

	// Execução das análises
	min, max, soma, media := calcularEstatisticas(numerosFloat)
	mediana := calcularMediana(numerosFloat)

	// ================= RELATÓRIO =================
	fmt.Println("\n================ RELATÓRIO ================")
	fmt.Printf("• Total de números válidos: %d\n", len(numerosFloat))
	fmt.Printf("• Menor valor (min): %.2f\n", min)
	fmt.Printf("• Maior valor (max): %.2f\n", max)
	fmt.Printf("• Soma total: %.2f\n", soma)
	fmt.Printf("• Média aritmética: %.2f\n", media)
	fmt.Printf("• Mediana: %.2f\n", mediana)

	fmt.Println("\nPROPRIEDADES DOS INTEIROS:")
	fmt.Printf("• Pares: %d | Ímpares: %d\n", contagemPares, contagemImpares)

	fmt.Print("• Números primos identificados: ")
	encontrouPrimo := false
	for _, n := range numerosInt {
		if ehPrimo(n) {
			fmt.Printf("%d ", n)
			encontrouPrimo = true
		}
	}
	if !encontrouPrimo {
		fmt.Print("Nenhum número primo encontrado.")
	}
	fmt.Println()
}
