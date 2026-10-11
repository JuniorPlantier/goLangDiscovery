package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Verifica se o ano é bissexto
func ehAnoBissexto(ano int) bool {
	return (ano%4 == 0 && ano%100 != 0) || (ano%400 == 0)
}

// Retorna o trimestre do ano com base no mês
func obterTrimestre(mes time.Month) int {
	switch {
	case mes >= 1 && mes <= 3:
		return 1
	case mes >= 4 && mes <= 6:
		return 2
	case mes >= 7 && mes <= 9:
		return 3
	default:
		return 4
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== ANALISADOR DE DATAS EM GO ===")
	fmt.Print("Digite uma data (formato DD/MM/AAAA): ")

	if !scanner.Scan() {
		fmt.Println("Erro ao ler entrada.")
		return
	}

	entrada := strings.TrimSpace(scanner.Text())

	// Em Go, a referência mágica para parsing de datas é sempre: 02/01/2006
	layout := "02/01/2006"
	data, err := time.Parse(layout, entrada)
	if err != nil {
		fmt.Printf("⚠️ Aviso: '%s' não é uma data válida ou não está no formato DD/MM/AAAA.\n", entrada)
		return
	}

	// Extração de dados básicos
	dia := data.Day()
	mes := data.Month()
	ano := data.Year()
	diaSemana := data.Weekday()
	diaDoAno := data.YearDay()
	bissexto := ehAnoBissexto(ano)
	trimestre := obterTrimestre(mes)

	// Cálculo de dias restantes no ano
	totalDiasAno := 365
	if bissexto {
		totalDiasAno = 366
	}
	diasRestantes := totalDiasAno - diaDoAno

	// Dicionários para tradução de dias e meses para o português
	diasSemanaPT := map[time.Weekday]string{
		time.Sunday:    "Domingo",
		time.Monday:    "Segunda-feira",
		time.Tuesday:   "Terça-feira",
		time.Wednesday: "Quarta-feira",
		time.Thursday:  "Quinta-feira",
		time.Friday:    "Sexta-feira",
		time.Saturday:  "Sábado",
	}

	mesesPT := map[time.Month]string{
		time.January: "Janeiro", time.February: "Fevereiro", time.March: "Março",
		time.April: "Abril", time.May: "Maio", time.June: "Junho",
		time.July: "Julho", time.August: "Agosto", time.September: "Setembro",
		time.October: "Outubro", time.November: "Novembro", time.December: "Dezembro",
	}

	// Comparação com o momento atual do sistema
	hoje := time.Now()
	// Zeramos as horas para comparar apenas os dias calendários com precisão
	hojeZero := time.Date(hoje.Year(), hoje.Month(), hoje.Day(), 0, 0, 0, 0, hoje.Location())
	dataZero := time.Date(ano, mes, dia, 0, 0, 0, 0, data.Location())

	diferencaDias := int(dataZero.Sub(hojeZero).Hours() / 24)

	// ================= RELATÓRIO =================
	fmt.Println("\n================ RELATÓRIO ================")
	fmt.Printf("• Data completa: %02d de %s de %d\n", dia, mesesPT[mes], ano)
	fmt.Printf("• Dia da semana: %s\n", diasSemanaPT[diaSemana])
	fmt.Printf("• Trimestre: %dº Trimestre\n", trimestre)
	
	if bissexto {
		fmt.Println("• Ano bissexto? Sim! 🟢")
	} else {
		fmt.Println("• Ano bissexto? Não. 🔴")
	}

	fmt.Printf("• Posição no ano: %dº dia (restam %d dias para o fim do ano)\n", diaDoAno, diasRestantes)

	fmt.Println("\nRELAÇÃO COM HOJE:")
	if diferencaDias == 0 {
		fmt.Println("• Esta data é HOJE! 📍")
	} else if diferencaDias > 0 {
		fmt.Printf("• É uma data FUTURA (daqui a %d dias).\n", diferencaDias)
	} else {
		fmt.Printf("• É uma data PASSADA (há %d dias atrás).\n", -diferencaDias)
	}
}