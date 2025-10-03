package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 1. Processamento sequencial (sem goroutines)
	fmt.Println("1. 🐌 SEQUENCIAL - Sem goroutines:")
	exemploSequencial()

	fmt.Println()
	fmt.Println("--------------------------------------------------")
	fmt.Println()

	// 2. Goroutines com race condition (rápido mas incorreto)
	fmt.Println("2. ⚡❌ GOROUTINES - Com race condition (rápido mas incorreto):")
	exemploRaceCondition()

	fmt.Println()
	fmt.Println("--------------------------------------------------")
	fmt.Println()

	// 3. Goroutines com channels (rápido e correto)
	fmt.Println("3. ⚡✅ CHANNELS - Goroutines seguras:")
	exemploChannels()

	fmt.Println()
	fmt.Println("--------------------------------------------------")
	fmt.Println()

	// 4. Concorrência inteligente (sem race condition por design)
	fmt.Println("4. 🧠✅ CONCORRÊNCIA INTELIGENTE - Trabalho independente:")
	exemploConcorrenciaInteligente()
}

// 🐌 Método sequencial (sem goroutines)
func exemploSequencial() {
	start := time.Now()
	counter := 0

	fmt.Println("   Processando sequencialmente...")

	// Simula 5 "workers" processando sequencialmente
	for worker := 1; worker <= 5; worker++ {
		for i := 0; i < 10000000; i++ { // 5 × 10 milhões = 50 milhões
			counter++
			// Simula um pequeno processamento
			if i%50000 == 0 {
				_ = i * 10
			}
		}
	}

	duration := time.Since(start)
	fmt.Printf("\n   🐌 Resultado: %d (correto!)\n", counter)
	fmt.Printf("   ⏱️  Tempo: %v\n", duration)
}

// ❌ Método que demonstra race condition (com 500k incrementos)
func exemploRaceCondition() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup

	fmt.Println("   5 goroutines competindo pela mesma variável...")

	// 5 goroutines trabalhadoras
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			// Cada goroutine faz 10 milhões de incrementos (5 × 10 milhões = 50 milhões)
			for j := 0; j < 10000000; j++ {
				// PROBLEMA: Lê, processa, escreve sem proteção
				temp := counter
				counter = temp + 1

				// Mostra progresso ocasionalmente
				if i%50000 == 0 {
					_ = i * 10
				}
			}
		}(i)
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("\n   ⚡❌ Resultado: %d (deveria ser 50.000.000, mas não é!)\n", counter)
	fmt.Printf("   ⏱️  Tempo: %v\n", duration)
	fmt.Printf("   📝 %d incrementos foram perdidos devido ao race condition\n", 50000000-counter)
}

// ✅ Método que resolve com channels (versão simples e didática)
func exemploChannels() {
	start := time.Now()
	// Canal simples - sem buffer, mais didático
	// incrementCh: Canal para workers enviarem "sinais" de incremento
	// - Tipo bool: só precisamos do sinal, não importa o valor
	// - Sem buffer: comunicação síncrona (worker espera contador processar)
	incrementCh := make(chan bool)

	// doneCh: Canal para receber o resultado final do contador
	// - Tipo int: carrega o valor final do counter
	// - Permite main() aguardar o contador terminar completamente
	doneCh := make(chan int)

	counter := 0

	// Goroutine "contadora" - única que mexe no counter
	go func() {
		for range incrementCh { // para cada mensagem que chegar
			counter++ // processe, ou.. no nosso caso.. incremente
		}
		doneCh <- counter
	}()

	var wg sync.WaitGroup

	// 5 goroutines trabalhadoras - só enviam "sinais" de incremento
	fmt.Println("   Criando 5 goroutines...")
	start_creation := time.Now()

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			// Cada uma faz 10 milhões de incrementos
			for j := 0; j < 10000000; j++ {
				incrementCh <- true // Sinal: "incremente por favor!"
				_ = j * 10
			}
		}(i)
	}

	_ = time.Since(start_creation)
	fmt.Printf("   ⚡ Go Routines criadas!!\n")
	fmt.Println("   📋 Agora as goroutines executam seus loops em PARALELO...") // Espera todas terminarem de enviar sinais
	wg.Wait()
	close(incrementCh) // Fecha o canal - goroutine contadora vai parar

	// Pega o resultado final
	finalResult := <-doneCh
	duration := time.Since(start)

	fmt.Printf("\n   ⚡✅ Resultado: %d (sempre correto!)\n", finalResult)
	fmt.Printf("   ⏱️  Tempo: %v\n", duration)
	fmt.Println("   📝 Canal simples: workers enviam 'sinais', contador processa")
}

// 🧠 Método com concorrência inteligente (sem race condition por design)
func exemploConcorrenciaInteligente() {
	start := time.Now()
	fmt.Println("   Cada goroutine trabalha independentemente...")

	// Canal para receber resultados individuais de cada goroutine
	// - Cada goroutine envia seu próprio contador final
	// - Não há variável compartilhada = não há race condition!
	resultsCh := make(chan int, 5) // Buffer para 5 resultados

	var wg sync.WaitGroup

	// 5 goroutines trabalhadoras - cada uma com SEU PRÓPRIO contador
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			// 🔑 CHAVE: Cada goroutine tem sua PRÓPRIA variável local
			localCounter := 0

			// Cada uma conta de 1 a 10 milhões independentemente
			for j := 1; j <= 10000000; j++ {
				localCounter = j // Simula contagem: 1, 2, 3, ..., 10M
				_ = j * 10000
			}

			fmt.Printf("   ✅ Goroutine %d terminou: contador local = %d\n", id, localCounter)

			// Envia SEU resultado via canal (sem competição!)
			resultsCh <- localCounter
		}(i)
	}

	// Aguarda todas as goroutines terminarem
	wg.Wait()
	close(resultsCh)

	// Agrega os resultados (soma todos os contadores individuais)
	totalSum := 0
	resultsReceived := 0

	fmt.Println("   📊 Agregando resultados:")
	for result := range resultsCh {
		totalSum += result
		resultsReceived++
		fmt.Printf("   📥 Resultado %d: %d (soma acumulada: %d)\n", resultsReceived, result, totalSum)
	}

	duration := time.Since(start)
	fmt.Printf("\n   🧠✅ Resultado final: %d (sempre correto!)\n", totalSum)
	fmt.Printf("   ⏱️  Tempo: %v\n", duration)
	fmt.Println("   📝 Concorrência inteligente: cada goroutine trabalha independentemente")
	fmt.Printf("   🔍 Verificação: 5 goroutines × 10.000.000 = %d ✅\n", 5*10000000)
}
