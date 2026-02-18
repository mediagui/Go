# Documentación de la Arquitectura - Funciones Heterogéneas en Go

## 📋 Problema Original

En Go, **no se pueden retornar funciones con diferentes firmas directamente** de una misma función. Por ejemplo, esto NO es posible:

```go
// ❌ INCORRECTO
func GetGame(option int) ??? {
    case 1:
        return func(int) int { ... }      // firma: func(int) int
    case 2:
        return func(string) bool { ... }  // firma: func(string) bool
    // ERROR: tipos inconsistentes
}
```

El compilador requiere que el tipo de retorno sea consistente.

---

## ✅ Solución Implementada

Se utilizó una **interfaz común** que actúa como contrato universal, permitiendo encapsular funciones con diferentes firmas:

```
┌─────────────────────────────────────────────────────────────────┐
│                     ARQUITECTURA GENERAL                        │
└─────────────────────────────────────────────────────────────────┘

                         ┌──────────────┐
                         │ GameFunction │ (Interfaz)
                         │  Interface   │
                         └──────┬───────┘
                                │
                    ┌───────────┼───────────┐
                    │           │           │
              ┌─────▼──────┐   │      ┌────▼──────┐
              │ Wrapper[   │   │      │ Wrapper[  │
              │  int,int]  │   │      │ string,   │
              │            │   │      │  bool]    │
              │ Execute()  │   │      │           │
              │            │   │      │Execute()  │
              └────────────┘   │      └───────────┘
                                │
                    ┌───────────▼──────────┐
                    │   GetGame() retorna  │
                    │   GameFunction (OK)  │
                    └──────────────────────┘
```

---

## 🏗️ Componentes Principales

### 1. **Interfaz GameFunction**
```go
type GameFunction interface {
    Execute(args ...any) any
}
```
- Define un contrato único: método `Execute()`
- Permite que diferentes tipos implementen esta interfaz
- Proporciona una forma uniforme de ejecutar funciones heterogéneas

### 2. **Estructura FunctionWrapper[T, R any]**
```go
type FunctionWrapper[T, R any] struct {
    fn func(T) R
}
```
- **T**: Tipo del parámetro de entrada
- **R**: Tipo del valor de retorno
- Implementa la interfaz GameFunction
- Cada wrapper encapsula una función con firma específica

**Ejemplo de tipos diferentes:**
```
FunctionWrapper[int, int]      → func(int) int
FunctionWrapper[string, bool]  → func(string) bool
FunctionWrapper[any, any]      → func(any) any
```

### 3. **Función buildFunction[T, R any]**
```go
func buildFunction[T any, R any](f func(T) R) GameFunction {
    return &FunctionWrapper[T, R]{fn: f}
}
```
- Factory que crea wrappers genericos
- Transforma una función específica en una GameFunction
- Mantiene type safety gracias a los genéricos

### 4. **Función GetGame(optionSelected int)**
```go
func GetGame(optionSelected int) GameFunction {
    switch optionSelected {
    case 1:
        return buildFunction(func(p any) any { ... })
    case 2:
        return buildFunction(func(p any) any { ... })
    default:
        return buildFunction(func(p any) any { ... })
    }
}
```
- Punto de entrada principal
- Retorna siempre una GameFunction (interfaz)
- Cada caso puede contener una función con firma diferente internamente

---

## 🔄 Flujo de Ejecución

```
┌──────────────────────────────────────────────────────────────┐
│ FLUJO COMPLETO: Desde main() hasta la ejecución del juego    │
└──────────────────────────────────────────────────────────────┘

1. main() → Muestra menú
              │
2.           Lectura de opción → selectedOption = 1
              │
3.           Validación
              │
4. playGame(1)
   │
5. game := games.GetGame(1)
   │
   ├─→ GetGame(1) evalúa switch
   │   │
   │   └─→ case 1:
   │       ├─→ Crea función lambda: func(p any) any { ... }
   │       ├─→ Llama buildFunction()
   │       │   │
   │       │   └─→ buildFunction crea:
   │       │       &FunctionWrapper[any, any]{fn: lambda}
   │       │
   │       └─→ Retorna GameFunction (el wrapper)
   │
6. game es ahora GameFunction (interfaz) que encapsula
   FunctionWrapper[any, any] que contiene la función lambda
   │
7. result := game.Execute(nil)
   │
   ├─→ Llama al método Execute() de FunctionWrapper
   │   │
   │   ├─→ Convierte args a tipo esperado (any)
   │   ├─→ Ejecuta fw.fn(arg) ← Ejecuta la función lambda
   │   │   │
   │   │   └─→ Función pide número al usuario
   │   │       Calcula factorial
   │   │       Retorna resultado
   │   │
   │   └─→ Retorna resultado como any
   │
8. Imprime resultado en consola
```

---

## 📊 Diagrama de Tipos

```
CASO 1: Función Factorial
┌──────────────────────────────────────────┐
│ Función Lambda                           │
│ func(any) any { calcular factorial() }   │
│                                          │
│ Entrada: any (no se usa)                │
│ Salida: int (convertido a any)          │
└────────────────────┬─────────────────────┘
                     │
                     ▼
         ┌──────────────────────┐
         │ buildFunction()      │
         │ [T=any, R=any]       │
         └────────────┬─────────┘
                      │
                      ▼
         ┌──────────────────────────┐
         │ FunctionWrapper[any,any] │
         │ {                        │
         │   fn: lambda             │
         │ }                        │
         │ Execute() → any          │
         └────────────┬─────────────┘
                      │
                      ▼
         ┌──────────────────────────┐
         │ GameFunction (interfaz)  │
         │ ✓ Compatible             │
         └──────────────────────────┘
```

---

## 🎯 Casos de Uso Actuales

### Caso 1: Factorial
```go
case 1:
    return buildFunction(func(parametersGame any) any {
        // Pide número al usuario
        // Calcula: 5! = 1 * 2 * 3 * 4 * 5 = 120
        // Retorna el factorial
    })
```

### Casos 2-5: Placeholders
```go
case 2, 3, 4, 5:
    v.ShowMockText(c.SELECTED)
    fallthrough
default:
    return buildFunction(func(parametersGame any) any {
        return nil
    })
```

---

## 🚀 Cómo Agregar Nuevos Ejercicios

### Ejemplo: Agregar validación de vocales

```go
case 2:
    return buildFunction(func(char string) bool {
        vowels := "aeiouAEIOU"
        if len(char) == 1 {
            return strings.Contains(vowels, char)
        }
        return false
    })
```

**Ventajas:**
1. Firma diferente: `func(string) bool` en lugar de `func(any) any`
2. buildFunction() crea `FunctionWrapper[string, bool]`
3. No requiere cambios en `main.go` o `playGame()`
4. Funciona automáticamente con `game.Execute()`

---

## 💡 Ventajas de Esta Arquitectura

| Aspecto | Ventaja |
|--------|---------|
| **Type Safety** | Los genéricos mantienen seguridad de tipos |
| **Extensibilidad** | Fácil agregar nuevas funciones con diferentes firmas |
| **Interfaz Unificada** | `main.go` no necesita cambios |
| **Limpieza** | Sin casting innecesario |
| **Mantenibilidad** | Código autodocumentado con genéricos |
| **Flexibilidad** | Cada función puede tener su propia lógica |

---

## 🔑 Conceptos Clave

### Genéricos en Go
```go
buildFunction[T any, R any](f func(T) R) GameFunction
```
- `[T any, R any]`: Parámetros genéricos
- `T`: Tipo de entrada (puede ser int, string, bool, etc.)
- `R`: Tipo de retorno (puede ser int, string, bool, etc.)

### Type Assertion
```go
arg = args[0].(T)  // Convierte any a tipo T
```
- Necesaria para pasar de `any` al tipo específico

### Interfaz Empty
```go
type GameFunction interface {
    Execute(args ...any) any
}
```
- Define el contrato mínimo necesario
- Cualquier tipo que implemente Execute() es GameFunction

---

## 📝 Resumen

La solución implementa un **patrón Factory + Wrapper con Genéricos** que permite:

1. ✅ Retornar funciones con diferentes firmas
2. ✅ Mantener type safety
3. ✅ Proporcionar interfaz uniforme
4. ✅ Facilitar extensión del código
5. ✅ Reducir duplicación en el código cliente

**En una línea:** Usar una interfaz como contrato común para encapsular funciones genéricas con diferentes tipos.
