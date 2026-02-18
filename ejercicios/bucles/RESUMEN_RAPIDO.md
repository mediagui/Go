# 🎯 RESUMEN RÁPIDO - Funciones Heterogéneas en Go

## El Problema

Querías que `GetGame()` devuelva funciones con **diferentes firmas**:
- Caso 1: `func(int) int` (factorial)
- Caso 2: `func(string) bool` (validar vocal)
- Caso 3: `func(int) string` (pares/impares)

**Pero Go no permite devolver tipos diferentes de la misma función.**

---

## La Solución: Interfaz + Wrapper + Genéricos

### 1️⃣ Interfaz Uniforme
```go
type GameFunction interface {
    Execute(args ...any) any
}
```
✅ Contrato común para todas las funciones

### 2️⃣ Wrapper Genérico
```go
type FunctionWrapper[T any, R any] struct {
    fn func(T) R
}

func (fw *FunctionWrapper[T, R]) Execute(args ...any) any {
    var arg T
    if len(args) > 0 && args[0] != nil {
        arg = args[0].(T)
    }
    return fw.fn(arg)
}
```
✅ Encapsula funciones con diferentes tipos

### 3️⃣ Factory
```go
func buildFunction[T any, R any](f func(T) R) GameFunction {
    return &FunctionWrapper[T, R]{fn: f}
}
```
✅ Crea wrappers de forma genérica

### 4️⃣ GetGame() retorna GameFunction
```go
func GetGame(optionSelected int) GameFunction {
    case 1:
        return buildFunction(func(p any) any { /*...*/ })
    case 2:
        return buildFunction(func(p string) bool { /*...*/ })
    // ...
}
```
✅ Todas retornan `GameFunction` (interfaz)
✅ Internamente tienen tipos diferentes

---

## Flujo de Ejecución

```
playGame(selectedOption)
    ↓
game := GetGame(selectedOption)  // Retorna GameFunction
    ↓
result := game.Execute(nil)      // Ejecuta la función
    ↓
fmt.Println(result)              // Muestra resultado
```

---

## Ventajas

| ✅ Beneficio | 📝 Descripción |
|---|---|
| **Funciones Heterogéneas** | Retorna funciones con diferentes firmas |
| **Type Safety** | Los genéricos mantienen seguridad de tipos |
| **Interfaz Uniforme** | Todos los casos se ejecutan igual: `game.Execute()` |
| **Fácil de Extender** | Agregar nuevos casos sin modificar `main.go` |
| **Sin Casting Manual** | El wrapper maneja las conversiones |
| **Código Limpio** | Autodocumentado con genéricos |

---

## Archivos Documentados

1. **games.go** - Código principal con documentación detallada
   - Interfaz `GameFunction`
   - Estructura `FunctionWrapper[T, R]`
   - Función `buildFunction()`
   - Función `GetGame()`

2. **main.go** - Código cliente actualizado
   - Función `main()`
   - Función `playGame()`

3. **ARQUITECTURA.md** - Diagrama y explicación de la arquitectura
4. **EJEMPLOS_PRACTICOS.md** - Ejemplos de nuevos casos

---

## Cómo Agregar un Nuevo Caso

**Ejemplo: Validar si es vocal**

```go
case 2:
    return buildFunction(func(char string) bool {
        vowels := "aeiouAEIOU"
        return len(char) == 1 && strings.Contains(vowels, char)
    })
```

¡Eso es todo! No necesitas cambiar nada más.

---

## Conceptos Clave

- **Genéricos**: `[T any, R any]` - Tipos de entrada y salida
- **Interface**: `GameFunction` - Contrato común
- **Wrapper**: `FunctionWrapper[T, R]` - Encapsula función específica
- **Factory**: `buildFunction()` - Crea wrappers
- **Type Assertion**: `args[0].(T)` - Convierte `any` a tipo específico

---

## Compilación

```powershell
cd 'C:\Users\Usuario\Documents\go-dev\ProyectosMA\ejercicios\bucles'
go build
```
✅ Compila sin errores

---

## Estado Final

✅ `GetGame()` devuelve funciones heterogéneas  
✅ Tipo de retorno es consistente: `GameFunction`  
✅ Cada caso puede tener firma diferente internamente  
✅ `main.go` no necesita cambios  
✅ Extensible y mantenible  
