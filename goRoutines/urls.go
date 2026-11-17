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

	var wg sync.WaitGroup // Cria a variável de controle. "O Gerenciador de Tráfego".

	inicio := time.Now()
	fmt.Println("Iniciando chamadas")

	for _, url := range urls {
		wg.Add(1)
		go verifyServiceStatus(url, &wg) // aqui está a goRoutine
		/*  A palavra-chave go: qdo vc coloca a palavra go na frente de uma chamada de função,
		    você está dizendo ao Go: "Inicie esta função em segundo plano (em paralelo) e
			continue a execução do loop imediatamente, sem esperar ela terminar".
		*/
	}
	wg.Wait() /* ela pausa a execução da main e diz: "Fique aqui travada esperando o contador chegar a zero".
	   Assim que a última goroutine termina e chama o Done(), a main é liberada para imprimir o tempo final. */

	fmt.Printf("\nFINALIZADO EM: %.2fs\n", time.Since(inicio).Seconds())
}

func verifyServiceStatus(url string, wg *sync.WaitGroup) {
	defer wg.Done() // "Terminei minha tarefa, pode subtrair 1 do contador".

	start := time.Now()

	res, err := http.Get(url)

	transcorrido := time.Since(start).Seconds()
	if err != nil {
		fmt.Printf("      [ERROR] %s está fora do ar é inválido! (%.2fs)\n", url, transcorrido)
	} else {
		fmt.Printf("      [OK] %s - Status: %d - Tempo: %.2fs\n", url, res.StatusCode, transcorrido)
	}
}
