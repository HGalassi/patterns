package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

// Definindo structs para demonstração
type Person struct {
	Name    string
	Age     int
	Email   string
	Address Address
}

type Address struct {
	Street  string
	City    string
	ZipCode string
	Country string
}

// Definindo interfaces
type Speaker interface {
	Speak() string
	GetLanguage() string
}

type Worker interface {
	Work() string
	GetSalary() float64
}

// Implementando interfaces
type Human struct {
	Name     string
	Language string
	Job      string
	Salary   float64
}

func (h Human) Speak() string {
	return fmt.Sprintf("%s says hello in %s", h.Name, h.Language)
}

func (h Human) GetLanguage() string {
	return h.Language
}

func (h Human) Work() string {
	return fmt.Sprintf("%s works as %s", h.Name, h.Job)
}

func (h Human) GetSalary() float64 {
	return h.Salary
}

// Tipo de função personalizado
type MathOperation func(int, int) int
type StringProcessor func(string) string

// Métodos em tipos customizados
type Temperature float64

func (t Temperature) Celsius() float64 {
	return float64(t)
}

func (t Temperature) Fahrenheit() float64 {
	return float64(t)*9/5 + 32
}

func (t Temperature) Kelvin() float64 {
	return float64(t) + 273.15
}

func main() {
	fmt.Println("=== TIPOS NÃO PRIMITIVOS EM GO ===")
	fmt.Println()

	// 1. ARRAYS
	fmt.Println("1. ARRAYS (Tamanho Fixo)")
	fmt.Println("-------------------------")

	var intArray [5]int = [5]int{1, 2, 3, 4, 5}
	var stringArray [3]string = [3]string{"Go", "Python", "Java"}
	var boolArray [4]bool = [4]bool{true, false, true, false}

	// Array com inicialização automática de tamanho
	autoArray := [...]int{10, 20, 30, 40}

	fmt.Printf("intArray: %v (tipo: %T, tamanho: %d)\n", intArray, intArray, len(intArray))
	fmt.Printf("stringArray: %v (tipo: %T, tamanho: %d)\n", stringArray, stringArray, len(stringArray))
	fmt.Printf("boolArray: %v (tipo: %T, tamanho: %d)\n", boolArray, boolArray, len(boolArray))
	fmt.Printf("autoArray: %v (tipo: %T, tamanho: %d)\n", autoArray, autoArray, len(autoArray))
	fmt.Printf("Tamanho em bytes do intArray: %d\n", unsafe.Sizeof(intArray))
	fmt.Println()

	// 2. SLICES
	fmt.Println("2. SLICES (Arrays Dinâmicos)")
	fmt.Println("----------------------------")

	var nilSlice []int                   // slice nil
	var emptySlice = []int{}             // slice vazio
	var intSlice = []int{1, 2, 3, 4, 5}  // slice com valores
	var makeSlice = make([]string, 3, 5) // make(tipo, tamanho, capacidade)

	// Operações com slices
	intSlice = append(intSlice, 6, 7, 8) // adicionando elementos
	subSlice := intSlice[2:5]            // slice de slice

	fmt.Printf("nilSlice: %v (len: %d, cap: %d, tipo: %T)\n", nilSlice, len(nilSlice), cap(nilSlice), nilSlice)
	fmt.Printf("emptySlice: %v (len: %d, cap: %d, tipo: %T)\n", emptySlice, len(emptySlice), cap(emptySlice), emptySlice)
	fmt.Printf("intSlice: %v (len: %d, cap: %d, tipo: %T)\n", intSlice, len(intSlice), cap(intSlice), intSlice)
	fmt.Printf("makeSlice: %v (len: %d, cap: %d, tipo: %T)\n", makeSlice, len(makeSlice), cap(makeSlice), makeSlice)
	fmt.Printf("subSlice: %v (len: %d, cap: %d, tipo: %T)\n", subSlice, len(subSlice), cap(subSlice), subSlice)
	fmt.Println()

	// 3. MAPS
	fmt.Println("3. MAPS (Chave-Valor)")
	fmt.Println("---------------------")

	var nilMap map[string]int          // map nil
	var emptyMap = map[string]int{}    // map vazio
	var makeMap = make(map[string]int) // criado com make

	// Map com valores iniciais
	studentGrades := map[string]int{
		"Alice":   95,
		"Bob":     87,
		"Charlie": 92,
		"Diana":   88,
	}

	// Operações com maps
	studentGrades["Eve"] = 90               // adicionando
	grade, exists := studentGrades["Alice"] // verificando existência
	delete(studentGrades, "Bob")            // removendo

	fmt.Printf("nilMap: %v (len: %d, tipo: %T)\n", nilMap, len(nilMap), nilMap)
	fmt.Printf("emptyMap: %v (len: %d, tipo: %T)\n", emptyMap, len(emptyMap), emptyMap)
	fmt.Printf("makeMap: %v (len: %d, tipo: %T)\n", makeMap, len(makeMap), makeMap)
	fmt.Printf("studentGrades: %v (len: %d, tipo: %T)\n", studentGrades, len(studentGrades), studentGrades)
	fmt.Printf("Alice's grade: %d (exists: %t)\n", grade, exists)
	fmt.Println()

	// 4. STRUCTS
	fmt.Println("4. STRUCTS (Tipos Compostos)")
	fmt.Println("----------------------------")

	// Diferentes formas de criar structs
	var emptyPerson Person // struct zero

	personLiteral := Person{
		Name:  "João Silva",
		Age:   30,
		Email: "joao@email.com",
		Address: Address{
			Street:  "Rua das Flores, 123",
			City:    "São Paulo",
			ZipCode: "01234-567",
			Country: "Brasil",
		},
	}

	personPartial := Person{
		Name: "Maria Santos",
		Age:  25,
	}

	// Struct anônima
	anonymousStruct := struct {
		ID   int
		Name string
	}{
		ID:   1,
		Name: "Produto Teste",
	}

	fmt.Printf("emptyPerson: %+v (tipo: %T)\n", emptyPerson, emptyPerson)
	fmt.Printf("personLiteral: %+v (tipo: %T)\n", personLiteral, personLiteral)
	fmt.Printf("personPartial: %+v (tipo: %T)\n", personPartial, personPartial)
	fmt.Printf("anonymousStruct: %+v (tipo: %T)\n", anonymousStruct, anonymousStruct)
	fmt.Printf("Tamanho do struct Person: %d bytes\n", unsafe.Sizeof(personLiteral))
	fmt.Println()

	// 5. POINTERS
	fmt.Println("5. POINTERS (Ponteiros)")
	fmt.Println("-----------------------")

	var intPtr *int // ponteiro nil
	var x int = 42
	intPtr = &x // endereço de x

	var personPtr *Person = &personLiteral // ponteiro para struct

	// Dereferenciamento
	fmt.Printf("x: %d (endereço: %p)\n", x, &x)
	fmt.Printf("intPtr: %p (valor apontado: %d, tipo: %T)\n", intPtr, *intPtr, intPtr)
	fmt.Printf("personPtr: %p (nome: %s, tipo: %T)\n", personPtr, personPtr.Name, personPtr)
	fmt.Printf("Tamanho do ponteiro: %d bytes\n", unsafe.Sizeof(intPtr))

	// Modificando através do ponteiro
	*intPtr = 100
	fmt.Printf("Após modificação via ponteiro - x: %d\n", x)
	fmt.Println()

	// 6. FUNCTIONS (Funções como Tipos)
	fmt.Println("6. FUNCTIONS (Funções como Tipos)")
	fmt.Println("---------------------------------")

	// Definindo variáveis de função
	var add MathOperation = func(a, b int) int {
		return a + b
	}

	var multiply MathOperation = func(a, b int) int {
		return a * b
	}

	var toUpper StringProcessor = func(s string) string {
		return fmt.Sprintf("*** %s ***", s)
	}

	// Função que retorna função
	createMultiplier := func(factor int) MathOperation {
		return func(a, b int) int {
			return (a + b) * factor
		}
	}

	double := createMultiplier(2)

	fmt.Printf("add: %T\n", add)
	fmt.Printf("multiply: %T\n", multiply)
	fmt.Printf("toUpper: %T\n", toUpper)
	fmt.Printf("add(5, 3) = %d\n", add(5, 3))
	fmt.Printf("multiply(4, 7) = %d\n", multiply(4, 7))
	fmt.Printf("toUpper(\"golang\") = %s\n", toUpper("golang"))
	fmt.Printf("double(3, 4) = %d\n", double(3, 4))
	fmt.Println()

	// 7. INTERFACES
	fmt.Println("7. INTERFACES")
	fmt.Println("-------------")

	human := Human{
		Name:     "Ana Costa",
		Language: "Português",
		Job:      "Desenvolvedora",
		Salary:   8000.50,
	}

	// Interface como tipo
	var speaker Speaker = human
	var worker Worker = human

	// Interface vazia
	var emptyInterface interface{} = "Pode ser qualquer tipo"

	fmt.Printf("speaker: %T\n", speaker)
	fmt.Printf("worker: %T\n", worker)
	fmt.Printf("emptyInterface: %v (tipo: %T)\n", emptyInterface, emptyInterface)
	fmt.Printf("Speaker says: %s\n", speaker.Speak())
	fmt.Printf("Worker info: %s\n", worker.Work())
	fmt.Printf("Salary: R$ %.2f\n", worker.GetSalary())

	// Type assertion
	if h, ok := speaker.(Human); ok {
		fmt.Printf("Type assertion successful: %s\n", h.Name)
	}
	fmt.Println()

	// 8. CHANNELS
	fmt.Println("8. CHANNELS (Comunicação entre Goroutines)")
	fmt.Println("------------------------------------------")

	// Diferentes tipos de channels
	var nilChannel chan int               // channel nil
	unbufferedChan := make(chan string)   // channel sem buffer
	bufferedChan := make(chan int, 3)     // channel com buffer
	readOnlyChan := make(<-chan bool)     // channel só leitura
	writeOnlyChan := make(chan<- float64) // channel só escrita

	fmt.Printf("nilChannel: %v (tipo: %T)\n", nilChannel, nilChannel)
	fmt.Printf("unbufferedChan: %v (tipo: %T)\n", unbufferedChan, unbufferedChan)
	fmt.Printf("bufferedChan: %v (tipo: %T, cap: %d)\n", bufferedChan, bufferedChan, cap(bufferedChan))
	fmt.Printf("readOnlyChan: %v (tipo: %T)\n", readOnlyChan, readOnlyChan)
	fmt.Printf("writeOnlyChan: %v (tipo: %T)\n", writeOnlyChan, writeOnlyChan)

	// Demonstração prática com goroutines
	resultChan := make(chan string, 2)

	go func() {
		time.Sleep(100 * time.Millisecond)
		resultChan <- "Resultado da goroutine 1"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		resultChan <- "Resultado da goroutine 2"
	}()

	// Recebendo resultados
	result1 := <-resultChan
	result2 := <-resultChan

	fmt.Printf("Recebido: %s\n", result1)
	fmt.Printf("Recebido: %s\n", result2)
	fmt.Println()

	// 9. MÉTODOS
	fmt.Println("9. MÉTODOS (Associados a Tipos)")
	fmt.Println("-------------------------------")

	temp := Temperature(25.0)

	fmt.Printf("Temperatura: %.1f°C (tipo: %T)\n", temp.Celsius(), temp)
	fmt.Printf("Em Fahrenheit: %.1f°F\n", temp.Fahrenheit())
	fmt.Printf("Em Kelvin: %.1fK\n", temp.Kelvin())
	fmt.Println()

	// 10. DEMONSTRAÇÕES AVANÇADAS
	fmt.Println("10. DEMONSTRAÇÕES AVANÇADAS")
	fmt.Println("---------------------------")

	// Slice de interfaces
	speakers := []Speaker{
		Human{Name: "Carlos", Language: "Espanhol", Job: "Professor", Salary: 5000},
		Human{Name: "Marie", Language: "Francês", Job: "Chef", Salary: 6000},
	}

	fmt.Println("Slice de interfaces:")
	for i, s := range speakers {
		fmt.Printf("  [%d] %s\n", i, s.Speak())
	}

	// Map de functions
	operations := map[string]MathOperation{
		"add":      add,
		"multiply": multiply,
		"subtract": func(a, b int) int { return a - b },
	}

	fmt.Println("\nMap de funções:")
	for name, op := range operations {
		fmt.Printf("  %s(10, 5) = %d\n", name, op(10, 5))
	}

	// Channel de structs
	personChan := make(chan Person, 2)
	personChan <- personLiteral
	personChan <- personPartial
	close(personChan)

	fmt.Println("\nChannel de structs:")
	for person := range personChan {
		fmt.Printf("  Pessoa: %s, Idade: %d\n", person.Name, person.Age)
	}

	// 11. REFLECTION E TYPE ASSERTION
	fmt.Println("\n11. REFLECTION E TYPE CHECKING")
	fmt.Println("------------------------------")

	values := []interface{}{
		42,
		"Hello",
		[]int{1, 2, 3},
		map[string]int{"key": 123},
		Person{Name: "Test", Age: 30},
		func() { fmt.Println("Function") },
		make(chan int),
	}

	for i, v := range values {
		t := reflect.TypeOf(v)
		fmt.Printf("[%d] Valor: %v, Tipo: %v, Kind: %v\n", i, v, t, t.Kind())
	}

	// 12. COMPARAÇÃO DE TAMANHOS
	fmt.Println("\n12. COMPARAÇÃO DE TAMANHOS EM MEMÓRIA")
	fmt.Println("-------------------------------------")

	fmt.Printf("Array [5]int: %d bytes\n", unsafe.Sizeof([5]int{}))
	fmt.Printf("Slice []int: %d bytes (header)\n", unsafe.Sizeof([]int{}))
	fmt.Printf("Map map[string]int: %d bytes (header)\n", unsafe.Sizeof(map[string]int{}))
	fmt.Printf("Struct Person: %d bytes\n", unsafe.Sizeof(Person{}))
	fmt.Printf("Ponteiro *int: %d bytes\n", unsafe.Sizeof((*int)(nil)))
	fmt.Printf("Function func(): %d bytes\n", unsafe.Sizeof(func() {}))
	fmt.Printf("Interface interface{}: %d bytes\n", unsafe.Sizeof(interface{}(nil)))
	fmt.Printf("Channel chan int: %d bytes\n", unsafe.Sizeof(make(chan int)))
}
