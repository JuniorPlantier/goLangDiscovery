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

	fmt.Printf("Iniciando chamandas...\n")

	for _, url := range urls {
		verifyServiceStatus(url)
	}

	fmt.Printf("\nFINALIZADO EM: %.2fs\n", time.Since(inicio).Seconds())
}

func verifyServiceStatus(url string) {
	start := time.Now()
	fmt.Printf("[START] - Iniciando verificação: %s\n", url)

	res, err := http.Get(url)
	transcorrido := time.Since(start).Seconds()

	if err != nil {
		fmt.Printf("      [ERROR] %s está fora do ar é inválido! (%.2fs)\n", url, transcorrido)
	} else {
		fmt.Printf("      [OK] %s - Status: %d - Tempo: %.2fs\n", url, res.StatusCode, transcorrido)
	}
}
