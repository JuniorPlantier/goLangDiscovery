package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {

	urls := []string{
		"https://github.com/",
		"https://www.drawio.com/",
		"https://www.youtube.com/",
		"https://www.amazon.com.br/",
		"https://www.google.com/",
		"https://www.einstein.br/estrutura/banco-sangue/agendamento-online",
	}

	var wg sync.WaitGroup

	inicio := time.Now()
	fmt.Println("Iniciando chamadas")

	for _, url := range urls {
		wg.Add(1)
		go verifyServiceStatus(url, &wg)
	}
	wg.Wait()

	fmt.Printf("\nFINALIZADO EM: %.2fs\n", time.Since(inicio).Seconds())
}

func verifyServiceStatus(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()

	res, err := http.Get(url)

	transcorrido := time.Since(start).Seconds()
	if err != nil {
		fmt.Printf("      [ERROR] %s está fora do ar é inválido! (%.2fs)\n", url, transcorrido)
	} else {
		fmt.Printf("      [OK] %s - Status: %d - Tempo: %.2fs\n", url, res.StatusCode, transcorrido)
	}
}
