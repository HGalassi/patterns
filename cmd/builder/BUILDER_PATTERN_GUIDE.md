# Builder Pattern em Go - Guia de Implementação

## Visão Geral

O **Builder Pattern** é um padrão de design criacional que permite construir objetos complexos passo a passo, definindo apenas as propriedades que são necessárias. Este padrão é especialmente útil quando você tem objetos com muitos parâmetros opcionais ou quando a construção do objeto requer múltiplas etapas.

## Quando Usar

- Objetos com muitos parâmetros opcionais
- Construção complexa que requer múltiplas etapas
- Necessidade de criar diferentes representações do mesmo objeto
- Quando você quer evitar construtores telescópicos (múltiplos construtores com diferentes combinações de parâmetros)

## Implementação em Go

### Estrutura Básica

```go
// 1. Definir o objeto que será construído
type dailyRoutine struct {
    familyTime     int
    work           int
    sleep          string
    eat            string
    programming    string
    hasHobby       bool
    exercise       bool
    language_study bool
}

// 2. Definir o Builder
type DailyRoutineBuilder struct {
    dailyRoutine dailyRoutine
}

// 3. Função construtora do Builder
func NewDailyRoutineBuilder() *DailyRoutineBuilder {
    return &DailyRoutineBuilder{dailyRoutine: dailyRoutine{}}
}
```

### Métodos Builder (Fluent Interface)

Cada método de configuração deve:
- Retornar o próprio builder (`*DailyRoutineBuilder`)
- Permitir method chaining
- Seguir a convenção `SetXxx()`

```go
func (b *DailyRoutineBuilder) SetFamilyTime(hours int) *DailyRoutineBuilder {
    b.dailyRoutine.familyTime = hours
    return b
}

func (b *DailyRoutineBuilder) SetWork(hours int) *DailyRoutineBuilder {
    b.dailyRoutine.work = hours
    return b
}

func (b *DailyRoutineBuilder) SetSleep(hours string) *DailyRoutineBuilder {
    b.dailyRoutine.sleep = hours
    return b
}

// ... outros métodos Set
```

### Método Build

O método `Build()` finaliza a construção e retorna o objeto final:

```go
func (b *DailyRoutineBuilder) Build() dailyRoutine {
    return b.dailyRoutine
}
```

### Função de Conveniência

Para facilitar o uso, é comum criar uma função que retorna diretamente o builder:

```go
func NewDailyRoutine() *DailyRoutineBuilder {
    return NewDailyRoutineBuilder()
}
```

## Exemplo de Uso

```go
func main() {
    // Rotina completa (weeklyRoutine)
    weeklyRoutine := NewDailyRoutine().
        SetEat("3 meals").
        SetFamilyTime(2).
        SetWork(8).
        SetSleep("7-8 hours").
        SetProgramming("2 hours").
        SetHobby(true).
        SetExercise(true).
        SetLanguageStudy(true).
        Build()

    // Rotina simples (dailyRoutine)
    dailyRoutine := NewDailyRoutine().
        SetEat("2 meals").
        SetFamilyTime(1).
        SetSleep("6 hours").
        SetProgramming("1 hour").
        SetHobby(false).
        Build()

    fmt.Printf("Weekly: %+v\n", weeklyRoutine)
    fmt.Printf("Daily: %+v\n", dailyRoutine)
}
```

## Boas Práticas em Go

### 1. **Convenções de Nomenclatura**
- Builder struct: `XxxBuilder`
- Métodos setter: `SetXxx()`
- Função construtora: `NewXxxBuilder()`
- Função de conveniência: `NewXxx()`

### 2. **Method Chaining**
- Sempre retorne `*Builder` nos métodos setter
- Permite sintaxe fluente e legível

### 3. **Encapsulamento**
- Mantenha o objeto interno (struct) com campos privados (lowercase)
- Exponha apenas através dos métodos do builder

### 4. **Validação**
```go
func (b *DailyRoutineBuilder) Build() (dailyRoutine, error) {
    if b.dailyRoutine.work < 0 {
        return dailyRoutine{}, errors.New("work hours cannot be negative")
    }
    return b.dailyRoutine, nil
}
```

### 5. **Valores Padrão**
```go
func NewDailyRoutineBuilder() *DailyRoutineBuilder {
    return &DailyRoutineBuilder{
        dailyRoutine: dailyRoutine{
            work:    8,     // valor padrão
            sleep:   "8 hours",
            exercise: true,
        },
    }
}
```

## Vantagens

✅ **Construção passo a passo**: Permite criar objetos complexos de forma incremental  
✅ **Reutilização**: O mesmo código de construção pode ser usado para diferentes representações  
✅ **Legibilidade**: Código mais limpo e expressivo  
✅ **Flexibilidade**: Parâmetros opcionais sem construtores telescópicos  
✅ **Single Responsibility**: Separa a lógica de construção da lógica de negócio  

## Desvantagens

❌ **Complexidade**: Aumenta o número de classes/structs no código  
❌ **Overhead**: Para objetos simples pode ser desnecessário  

## Relacionamento com Outros Padrões

- **Factory Pattern**: Builder pode ser usado junto com Factory para criar diferentes tipos de builders
- **Singleton**: O builder pode implementar Singleton se houver necessidade de uma única instância
- **Composite**: Útil para construir árvores complexas de objetos compostos

## Casos de Uso Comuns

### 1. **SQL Query Builder**
```go
query := NewQueryBuilder().
    Select("name", "email").
    From("users").
    Where("age > 18").
    OrderBy("name").
    Build()
```

### 2. **Configuration Builder**
```go
config := NewConfigBuilder().
    SetDatabase("postgresql://...").
    SetPort(8080).
    SetDebug(true).
    Build()
```

### 3. **HTTP Request Builder**
```go
request := NewRequestBuilder().
    SetMethod("POST").
    SetURL("https://api.example.com").
    SetHeader("Content-Type", "application/json").
    SetBody(data).
    Build()
```

## Executando o Exemplo

Para executar o exemplo fornecido:

```bash
cd cmd/builder
go run .
```

**Importante**: Use `go run .` em vez de `go run main.go` para incluir todos os arquivos `.go` do pacote.

## Considerações Finais

O Builder Pattern é uma excelente escolha para construir objetos complexos em Go, especialmente quando há muitos parâmetros opcionais. A implementação em Go se beneficia da sintaxe fluente e da simplicidade da linguagem, resultando em código limpo e expressivo.

Lembre-se de avaliar se a complexidade adicional do padrão é justificada pela complexidade do objeto que está sendo construído.

https://medium.com/@josueparra2892/builder-pattern-in-go-56605f9e7387