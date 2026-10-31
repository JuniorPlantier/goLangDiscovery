package main

import (
	"fmt"
	"net/http"
	"time"
)

/* consulta cada um dos serviços sem utilizar concorrência */
func main() {

	urls := []string{
		"https://github.com/",
		"https://www.drawio.com/",
		"https://www.youtube.com/",
		"https://www.amazon.com.br/",
		"https://www.google.com/",
		"https://www.einstein.br/estrutura/banco-sangue/agendamento-online",
	}

	inicio := time.Now()

	fmt.Println("Iniciando verificações...")

	for _, url := range urls {
		verifyServiceStatus(url)
	}

	fmt.Println("\n FINALIZADO EM: %.2fs\n", time.Since(inicio).Seconds())
}

func verifyServiceStatus(url string) {
	start := time.Now()
	fmt.Printf("[START] - Iniciando verificação: %s\n", url)

	res, err := http.Get(url)
	elapsed := time.Since(start).Seconds()

	if err != nil {
		fmt.Printf("[ERRO] %s está fora do ar é inválido! (%.2fs)\n", url, elapsed)
	} else {
		fmt.Printf("[OK] %s - Status: %d - Tempo: %.2fs\n", url, res.StatusCode, elapsed)
	}
}
