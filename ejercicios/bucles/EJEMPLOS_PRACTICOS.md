# 📚 Ejemplos Prácticos - Funciones Heterogéneas

## Ejemplo 1: Función Factorial (Actual)

```go
case 1:
    return buildFunction(func(parametersGame any) any {
        n, _ := strconv.Atoi(v.ResquestValue(c.GET_NUMBER_TO_CALC_FACTORIAL))
        var r int = 1
        for i := 2; i <= n; i++ {
            r *= i
        }
        return r
    })
```

**Firma interna:** `func(any) any`  
**Entrada:** `any` (no se usa, se pide al usuario)  
**Salida:** `int` (convertido a `any`)  

**Ejecución:**
```
Usuario ingresa: 5
Cálculo: 5! = 1 * 2 * 3 * 4 * 5 = 120
Resultado: 120
```

---

## Ejemplo 2: Validar si es Vocal

Agregar este nuevo caso a `GetGame()`:

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

**Firma interna:** `func(string) bool`  
**Entrada:** `string` (un carácter)  
**Salida:** `bool` (es vocal o no)  

**Cómo funciona el wrapper:**
```
1. Execute("a") se llama
2. FunctionWrapper[string, bool] recibe args = ["a"]
3. Convierte args[0].(string) → "a"
4. Ejecuta fn("a") → comprueba si es vocal
5. Retorna true como any
```

---

## Ejemplo 3: Pares e Impares

```go
case 3:
    return buildFunction(func(n int) string {
        message := ""
        for i := 1; i <= n; i++ {
            if i%2 == 0 {
                message += fmt.Sprintf("%d - es par\n", i)
            } else {
                message += fmt.Sprintf("%d - es impar\n", i)
            }
        }
        return message
    })
```

**Firma interna:** `func(int) string`  
**Entrada:** `int` (número límite)  
**Salida:** `string` (listado formateado)  

**Ejecución:**
```
Usuario ingresa: 5
1 - es impar
2 - es par
3 - es impar
4 - es par
5 - es impar
```

---

## Ejemplo 4: Acumular Números hasta 50

```go
case 4:
    return buildFunction(func(_ any) any {
        var suma int = 0
        var contador int = 0
        
        for {
            numStr := v.ResquestValue("Ingresa un número: ")
            num, _ := strconv.Atoi(numStr)
            
            suma += num
            contador++
            
            if suma > 50 {
                break
            }
        }
        
        return fmt.Sprintf("Total: %d, Números: %d", suma, contador)
    })
```

**Firma interna:** `func(any) string`  
**Entrada:** `any` (se ignora)  
**Salida:** `string` (resultado formateado)  

**Ejecución:**
```
Ingresa un número: 20
Ingresa un número: 15
Ingresa un número: 20
Total: 55, Números: 3
```

---

## Ejemplo 5: Área de Triángulo y Rectángulo

```go
case 5:
    return buildFunction(func(_ any) any {
        figura := v.ResquestValue("¿Triángulo (T) o Rectángulo (R)?: ")
        
        if figura == "T" || figura == "t" {
            base, _ := strconv.Atoi(v.ResquestValue("Base: "))
            altura, _ := strconv.Atoi(v.ResquestValue("Altura: "))
            area := (base * altura) / 2
            return fmt.Sprintf("Área triángulo: %d", area)
        } else if figura == "R" || figura == "r" {
            base, _ := strconv.Atoi(v.ResquestValue("Base: "))
            altura, _ := strconv.Atoi(v.ResquestValue("Altura: "))
            area := base * altura
            return fmt.Sprintf("Área rectángulo: %d", area)
        }
        
        return "Opción no válida"
    })
```

**Firma interna:** `func(any) string`  
**Entrada:** `any` (se ignora)  
**Salida:** `string` (resultado del cálculo)  

---

## Tabla Comparativa de Firmas

| Caso | Entrada | Salida | Tipo Completo |
|------|---------|--------|---------------|
| 1 (Factorial) | `any` | `int` | `func(any) int` |
| 2 (Vocal) | `string` | `bool` | `func(string) bool` |
| 3 (Pares/Impares) | `int` | `string` | `func(int) string` |
| 4 (Suma hasta 50) | `any` | `string` | `func(any) string` |
| 5 (Área) | `any` | `string` | `func(any) string` |

**PERO** todas se retornan como:
```go
GameFunction  // interfaz uniforme
```

---

## Diagrama de Transformación

```
FUNCIÓN ORIGINAL
    │
    └─→ lambda: func(string) bool
             │
             └─→ buildFunction[string, bool](lambda)
                      │
                      └─→ &FunctionWrapper[string, bool]{fn: lambda}
                           │
                           └─→ Implementa GameFunction
                                │
                                └─→ Método Execute(args ...any) any
                                     │
                                     ├─→ Convierte args[0].(string)
                                     ├─→ Ejecuta fn(string)
                                     └─→ Retorna resultado como any
```

---

## Flujo Completo: Paso a Paso

### Ejecución del Caso 2 (Vocal)

**Paso 1: En main.go**
```go
selectedOption := 2  // Usuario selecciona opción 2
playGame(2)
```

**Paso 2: En playGame()**
```go
game := games.GetGame(2)     // Obtiene GameFunction
result := game.Execute(nil)  // Ejecuta
fmt.Println("Resultado:", result)
```

**Paso 3: En GetGame(2)**
```go
switch 2 {
case 2:
    return buildFunction(func(char string) bool {
        // Esta función se ejecutará después
        vowels := "aeiouAEIOU"
        if len(char) == 1 {
            return strings.Contains(vowels, char)
        }
        return false
    })
}
```

**Paso 4: buildFunction crea el wrapper**
```go
// buildFunction[string, bool](lambda)
return &FunctionWrapper[string, bool]{
    fn: lambda
}
// Retorna como GameFunction
```

**Paso 5: game.Execute(nil) se ejecuta**
```go
func (fw *FunctionWrapper[string, bool]) Execute(args ...any) any {
    var arg string  // Crea variable string vacía
    
    // args es nil, así que arg permanece vacío ""
    
    return fw.fn("")  // Ejecuta fn("")
    // Resultado: false (string vacío no es vocal)
}
```

**Paso 6: Resultado en consola**
```
Resultado: false
Opción seleccionada: 2
```

---

## Casos Avanzados

### Función que retorna Estructura

```go
case 6:
    type Resultado struct {
        Exito bool
        Valor int
        Mensaje string
    }
    
    return buildFunction(func(p any) Resultado {
        return Resultado{
            Exito: true,
            Valor: 42,
            Mensaje: "Éxito",
        }
    })
```

El wrapper sería: `FunctionWrapper[any, Resultado]`

### Función con Múltiples Operaciones

```go
case 7:
    return buildFunction(func(nums []int) map[string]int {
        resultado := make(map[string]int)
        suma := 0
        promedio := 0
        
        for _, n := range nums {
            suma += n
        }
        
        if len(nums) > 0 {
            promedio = suma / len(nums)
        }
        
        resultado["suma"] = suma
        resultado["promedio"] = promedio
        resultado["cantidad"] = len(nums)
        
        return resultado
    })
```

El wrapper sería: `FunctionWrapper[[]int, map[string]int]`

---

## 🎯 Resumen de Ventajas

```
┌──────────────────────────────────────────────────┐
│ ANTES (Sin la solución)                          │
├──────────────────────────────────────────────────┤
│ ❌ No se pueden retornar diferentes tipos       │
│ ❌ Necesitarías casting manual                  │
│ ❌ Código complejo y difícil de mantener        │
│ ❌ Error propenso                              │
└──────────────────────────────────────────────────┘

┌──────────────────────────────────────────────────┐
│ AHORA (Con Interfaz + Wrapper + Genéricos)      │
├──────────────────────────────────────────────────┤
│ ✅ Retornas diferentes tipos sin problemas      │
│ ✅ Type safety mediante genéricos               │
│ ✅ Código limpio y autodocumentado              │
│ ✅ Fácil de extender                            │
│ ✅ Sin casting manual                           │
│ ✅ Mantenible                                   │
└──────────────────────────────────────────────────┘
```

---

## 📖 Referencias

- **Genéricos en Go**: https://go.dev/doc/tutorial/generics
- **Interfaces en Go**: https://go.dev/doc/effective_go#interfaces
- **Type Assertions**: https://go.dev/tour/methods/15
