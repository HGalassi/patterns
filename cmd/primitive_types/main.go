package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	fmt.Println("=== TIPOS PRIMITIVOS EM GO ===")
	fmt.Println()

	// 1. TIPOS BOOLEANOS
	fmt.Println("1. TIPO BOOLEANO (bool)")
	fmt.Println("----------------------------")
	var isActive bool = true
	var isCompleted bool = false
	var defaultBool bool // valor zero: false

	fmt.Printf("isActive: %v (tipo: %T)\n", isActive, isActive)
	fmt.Printf("isCompleted: %v (tipo: %T)\n", isCompleted, isCompleted)
	fmt.Printf("defaultBool: %v (tipo: %T)\n", defaultBool, defaultBool)
	fmt.Printf("Tamanho em bytes: %d\n\n", unsafe.Sizeof(isActive))

	// 2. TIPOS INTEIROS COM SINAL
	fmt.Println("2. TIPOS INTEIROS COM SINAL")
	fmt.Println("---------------------------")

	var int8Var int8 = 127                   // -128 a 127
	var int16Var int16 = 32767               // -32768 a 32767
	var int32Var int32 = 2147483647          // -2147483648 a 2147483647
	var int64Var int64 = 9223372036854775807 // -9223372036854775808 a 9223372036854775807
	var intVar int = 42                      // Tamanho dependente da arquitetura (32 ou 64 bits)

	fmt.Printf("int8: %d (tipo: %T, tamanho: %d bytes)\n", int8Var, int8Var, unsafe.Sizeof(int8Var))
	fmt.Printf("int16: %d (tipo: %T, tamanho: %d bytes)\n", int16Var, int16Var, unsafe.Sizeof(int16Var))
	fmt.Printf("int32: %d (tipo: %T, tamanho: %d bytes)\n", int32Var, int32Var, unsafe.Sizeof(int32Var))
	fmt.Printf("int64: %d (tipo: %T, tamanho: %d bytes)\n", int64Var, int64Var, unsafe.Sizeof(int64Var))
	fmt.Printf("int: %d (tipo: %T, tamanho: %d bytes)\n\n", intVar, intVar, unsafe.Sizeof(intVar))

	// 3. TIPOS INTEIROS SEM SINAL
	fmt.Println("3. TIPOS INTEIROS SEM SINAL")
	fmt.Println("---------------------------")

	var uint8Var uint8 = 255                    // 0 a 255
	var uint16Var uint16 = 65535                // 0 a 65535
	var uint32Var uint32 = 4294967295           // 0 a 4294967295
	var uint64Var uint64 = 18446744073709551615 // 0 a 18446744073709551615
	var uintVar uint = 100                      // Tamanho dependente da arquitetura
	var uintptrVar uintptr = 0x1234             // Para armazenar ponteiros

	fmt.Printf("uint8: %d (tipo: %T, tamanho: %d bytes)\n", uint8Var, uint8Var, unsafe.Sizeof(uint8Var))
	fmt.Printf("uint16: %d (tipo: %T, tamanho: %d bytes)\n", uint16Var, uint16Var, unsafe.Sizeof(uint16Var))
	fmt.Printf("uint32: %d (tipo: %T, tamanho: %d bytes)\n", uint32Var, uint32Var, unsafe.Sizeof(uint32Var))
	fmt.Printf("uint64: %d (tipo: %T, tamanho: %d bytes)\n", uint64Var, uint64Var, unsafe.Sizeof(uint64Var))
	fmt.Printf("uint: %d (tipo: %T, tamanho: %d bytes)\n", uintVar, uintVar, unsafe.Sizeof(uintVar))
	fmt.Printf("uintptr: 0x%x (tipo: %T, tamanho: %d bytes)\n\n", uintptrVar, uintptrVar, unsafe.Sizeof(uintptrVar))

	// 4. ALIASES PARA TIPOS INTEIROS
	fmt.Println("4. ALIASES PARA TIPOS INTEIROS")
	fmt.Println("------------------------------")

	var byteVar byte = 255 // alias para uint8
	var runeVar rune = 'A' // alias para int32, usado para caracteres Unicode

	fmt.Printf("byte: %d (tipo: %T, tamanho: %d bytes)\n", byteVar, byteVar, unsafe.Sizeof(byteVar))
	fmt.Printf("rune: %d '%c' (tipo: %T, tamanho: %d bytes)\n\n", runeVar, runeVar, runeVar, unsafe.Sizeof(runeVar))

	// 5. TIPOS DE PONTO FLUTUANTE
	fmt.Println("5. TIPOS DE PONTO FLUTUANTE")
	fmt.Println("---------------------------")

	var float32Var float32 = 3.14159265359     // Precisão simples (32 bits)
	var float64Var float64 = 3.141592653589793 // Precisão dupla (64 bits)
	var defaultFloat = 2.718281828             // Inferido como float64

	fmt.Printf("float32: %.10f (tipo: %T, tamanho: %d bytes)\n", float32Var, float32Var, unsafe.Sizeof(float32Var))
	fmt.Printf("float64: %.15f (tipo: %T, tamanho: %d bytes)\n", float64Var, float64Var, unsafe.Sizeof(float64Var))
	fmt.Printf("defaultFloat: %.15f (tipo: %T, tamanho: %d bytes)\n\n", defaultFloat, defaultFloat, unsafe.Sizeof(defaultFloat))

	// 6. TIPOS COMPLEXOS
	fmt.Println("6. TIPOS COMPLEXOS")
	fmt.Println("------------------")

	var complex64Var complex64 = 3 + 4i    // Parte real e imaginária float32
	var complex128Var complex128 = 5 + 12i // Parte real e imaginária float64
	var defaultComplex = 1 + 2i            // Inferido como complex128

	fmt.Printf("complex64: %v (tipo: %T, tamanho: %d bytes)\n", complex64Var, complex64Var, unsafe.Sizeof(complex64Var))
	fmt.Printf("  Parte real: %.2f, Parte imaginária: %.2f\n", real(complex64Var), imag(complex64Var))
	fmt.Printf("complex128: %v (tipo: %T, tamanho: %d bytes)\n", complex128Var, complex128Var, unsafe.Sizeof(complex128Var))
	fmt.Printf("  Parte real: %.2f, Parte imaginária: %.2f\n", real(complex128Var), imag(complex128Var))
	fmt.Printf("defaultComplex: %v (tipo: %T, tamanho: %d bytes)\n\n", defaultComplex, defaultComplex, unsafe.Sizeof(defaultComplex))

	// 7. STRINGS
	fmt.Println("7. STRINGS")
	fmt.Println("----------")

	var stringVar string = "Olá, Go!"
	var emptyString string // valor zero: ""
	var unicodeString = "Texto com emoji: 🚀"
	var rawString = `String literal
com múltiplas linhas
e caracteres especiais: \n \t`

	fmt.Printf("stringVar: %q (tipo: %T, tamanho: %d bytes)\n", stringVar, stringVar, unsafe.Sizeof(stringVar))
	fmt.Printf("emptyString: %q (tipo: %T, tamanho: %d bytes)\n", emptyString, emptyString, unsafe.Sizeof(emptyString))
	fmt.Printf("unicodeString: %q (tipo: %T)\n", unicodeString, unicodeString)
	fmt.Printf("rawString: %q (tipo: %T)\n", rawString, rawString)
	fmt.Printf("Comprimento de stringVar: %d caracteres\n\n", len(stringVar))

	// 8. DEMONSTRAÇÕES PRÁTICAS
	fmt.Println("8. DEMONSTRAÇÕES PRÁTICAS")
	fmt.Println("-------------------------")

	// Conversões de tipo
	fmt.Println("Conversões de tipo:")
	var a int = 42
	var b float64 = float64(a)
	var c string = fmt.Sprintf("%d", a)
	fmt.Printf("int %d -> float64 %.2f -> string %q\n", a, b, c)

	// Operações com diferentes tipos
	fmt.Println("\nOperações matemáticas:")
	fmt.Printf("Soma inteiros: %d + %d = %d\n", int32Var, intVar, int32Var+int32(intVar))
	fmt.Printf("Multiplicação float: %.2f * %.2f = %.2f\n", float32Var, float64(float32Var), float64(float32Var)*float64(float32Var))
	fmt.Printf("Módulo complexo: |%v| = %.2f\n", complex128Var, realMagnitude(complex128Var))

	// Verificação de tipos com reflection
	fmt.Println("\nVerificação com reflection:")
	checkType := func(v interface{}) {
		t := reflect.TypeOf(v)
		fmt.Printf("Valor: %v, Tipo: %v, Kind: %v\n", v, t, t.Kind())
	}

	checkType(int8Var)
	checkType(float64Var)
	checkType(stringVar)
	checkType(complex128Var)

	// 9. VALORES ZERO DOS TIPOS
	fmt.Println("\n9. VALORES ZERO DOS TIPOS")
	fmt.Println("-------------------------")
	demonstrateZeroValues()

	// 10. LIMITES DOS TIPOS
	fmt.Println("\n10. LIMITES DOS TIPOS")
	fmt.Println("---------------------")
	demonstrateLimits()
}

// Função auxiliar para calcular magnitude de número complexo
func realMagnitude(c complex128) float64 {
	r, i := real(c), imag(c)
	return float64(r*r + i*i)
}

// Demonstra valores zero de todos os tipos
func demonstrateZeroValues() {
	var (
		zeroBool       bool
		zeroInt        int
		zeroInt8       int8
		zeroInt16      int16
		zeroInt32      int32
		zeroInt64      int64
		zeroUint       uint
		zeroUint8      uint8
		zeroUint16     uint16
		zeroUint32     uint32
		zeroUint64     uint64
		zeroUintptr    uintptr
		zeroByte       byte
		zeroRune       rune
		zeroFloat32    float32
		zeroFloat64    float64
		zeroComplex64  complex64
		zeroComplex128 complex128
		zeroString     string
	)

	fmt.Printf("bool zero value: %v\n", zeroBool)
	fmt.Printf("int zero value: %v\n", zeroInt)
	fmt.Printf("int8 zero value: %v\n", zeroInt8)
	fmt.Printf("int16 zero value: %v\n", zeroInt16)
	fmt.Printf("int32 zero value: %v\n", zeroInt32)
	fmt.Printf("int64 zero value: %v\n", zeroInt64)
	fmt.Printf("uint zero value: %v\n", zeroUint)
	fmt.Printf("uint8 zero value: %v\n", zeroUint8)
	fmt.Printf("uint16 zero value: %v\n", zeroUint16)
	fmt.Printf("uint32 zero value: %v\n", zeroUint32)
	fmt.Printf("uint64 zero value: %v\n", zeroUint64)
	fmt.Printf("uintptr zero value: %v\n", zeroUintptr)
	fmt.Printf("byte zero value: %v\n", zeroByte)
	fmt.Printf("rune zero value: %v\n", zeroRune)
	fmt.Printf("float32 zero value: %v\n", zeroFloat32)
	fmt.Printf("float64 zero value: %v\n", zeroFloat64)
	fmt.Printf("complex64 zero value: %v\n", zeroComplex64)
	fmt.Printf("complex128 zero value: %v\n", zeroComplex128)
	fmt.Printf("string zero value: %q\n", zeroString)
}

// Demonstra os limites de cada tipo numérico
func demonstrateLimits() {
	fmt.Println("Limites dos tipos inteiros:")
	fmt.Printf("int8: %d a %d\n", int8(-128), int8(127))
	fmt.Printf("uint8: %d a %d\n", uint8(0), uint8(255))
	fmt.Printf("int16: %d a %d\n", int16(-32768), int16(32767))
	fmt.Printf("uint16: %d a %d\n", uint16(0), uint16(65535))

	fmt.Println("\nComparação de tamanhos:")
	fmt.Printf("Arquitetura: %d bits\n", unsafe.Sizeof(uintptr(0))*8)
	fmt.Printf("int: %d bytes\n", unsafe.Sizeof(int(0)))
	fmt.Printf("uint: %d bytes\n", unsafe.Sizeof(uint(0)))
	fmt.Printf("uintptr: %d bytes\n", unsafe.Sizeof(uintptr(0)))
}
