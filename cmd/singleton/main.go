package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Teste Singleton com Struct ===")

	// Teste básico - todas devem ser a mesma instância
	s1 := GetInstance_example_1()
	s2 := GetInstance_example_2()
	s3 := GetInstance_example_3()
	s4 := GetInstance_example_4()

	fmt.Printf("Instância 1: %p\n", s1)
	fmt.Printf("Instância 2: %p\n", s2)
	fmt.Printf("Instância 3: %p\n", s3)
	fmt.Printf("Instância 4: %p\n", s4)

	// Teste dos métodos da struct
	fmt.Println("\n=== Teste dos Métodos ===")
	s3.Set("nome", "João")
	s3.Set("idade", "30")
	s3.Set("cidade", "São Paulo")

	nome, exists := s3.Get("nome")
	fmt.Printf("Nome: %s (existe: %t)\n", nome, exists)

	idade, exists := s3.Get("idade")
	fmt.Printf("Idade: %s (existe: %t)\n", idade, exists)

	fmt.Printf("Tamanho do map: %d\n", s3.Size())

	// Verificar se todas as instâncias compartilham o mesmo dado
	fmt.Println("\n=== Teste de Compartilhamento ===")
	nome1, _ := s1.Get("nome")
	nome2, _ := s2.Get("nome")
	nome4, _ := s4.Get("nome")

	fmt.Printf("Nome via s1: %s\n", nome1)
	fmt.Printf("Nome via s2: %s\n", nome2)
	fmt.Printf("Nome via s4: %s\n", nome4)

	// Teste de concorrência
	fmt.Println("\n=== Teste de Concorrência exemplo 3 ===")
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			instance := GetInstance_example_3()
			instance.Set(fmt.Sprintf("goroutine_%d", id), fmt.Sprintf("valor_%d", id))
			time.Sleep(time.Millisecond * 10)
			fmt.Printf("Goroutine %d: Instância %p, Tamanho: %d\n", id, instance, instance.Size())
		}(i)
	}
	wg.Wait()

	fmt.Println("\n=== Teste de Concorrência exemplo 4 ===")
	wg = sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			instance := GetInstance_example_4()
			instance.Set(fmt.Sprintf("goroutine_%d", id), fmt.Sprintf("valor_%d", id))
			time.Sleep(time.Millisecond * 10)
			fmt.Printf("Goroutine %d: Instância %p, Tamanho: %d\n", id, instance, instance.Size())
		}(i)
	}
	wg.Wait()

	fmt.Println("=== Fim dos Testes ===")
}
