# Documentación del archivo main.go

Archivo fuente: [main.go](ejercicios/structs/main.go#L1-L74)

Descripción general
- Este programa define un tipo `personaStruct` y pide al usuario que introduzca los valores de cada campo. Utiliza reflection para iterar los campos del struct, establecer sus valores y leer las etiquetas definidas en las tags del struct.

Notas importantes
- No se ha modificado el código fuente; la documentación explica las funciones y las sentencias usadas, con enlaces a la documentación oficial de Go para facilitar la consulta.

Funciones y sentencias documentadas

1) `main()`
- Propósito: crea una instancia de `personaStruct`, solicita los datos al usuario y muestra los campos.
- Llamadas principales en el código: `completaDatosStruct(persona)` y `getStructFields(persona)`.
- Enlaces útiles:
  - Paquete `fmt`: https://pkg.go.dev/fmt
  - Paquete `reflect`: https://pkg.go.dev/reflect

2) `completaDatosStruct(persona *personaStruct)`
- Propósito: recorre cada campo del struct `persona` con reflection y asigna a cada campo el valor leído desde la entrada estándar.
- Sentencias y funciones clave usadas:
  - `reflect.ValueOf(persona).Elem()` — obtiene el valor direccionable del valor pasado (necesario para modificar campos).
    Enlace: https://pkg.go.dev/reflect#ValueOf
  - `reflect.TypeOf(persona).Elem()` — obtiene el tipo del struct apuntado.
    Enlace: https://pkg.go.dev/reflect#TypeOf
  - `personaType.NumField()` — número de campos en el tipo struct.
    Enlace: https://pkg.go.dev/reflect#Type.NumField
  - `personaValue.Field(i)` — obtiene el i-ésimo campo como `reflect.Value`.
    Enlace: https://pkg.go.dev/reflect#Value.Field
  - `campo.SetString(valor)` — asigna un valor string al campo (requiere que el campo sea asignable y de tipo string).
    Enlace: https://pkg.go.dev/reflect#Value.SetString
- Observaciones:
  - Es imprescindible pasar al reflection un puntero al struct y obtener la `Elem()` para que los campos sean direccionables y puedan ser modificados.

3) `pideValorParaCampo(fieldName string) string`
- Propósito: mostrar un prompt por consola y leer una entrada de tipo string.
- Sentencias y funciones clave:
  - `fmt.Printf("Introduzca el valor para %s: ", fieldName)` — muestra el prompt.
    Enlace: https://pkg.go.dev/fmt#Printf
  - `fmt.Scanln(&valor)` — lee un token/string de la entrada estándar hasta nueva línea.
    Enlace: https://pkg.go.dev/fmt#Scanln

4) `getStructFields(p *personaStruct)`
- Propósito: obtener información del tipo y los valores del struct y mostrar cada etiqueta (`etiqueta`) y su valor asociado.
- Sentencias y funciones clave:
  - `reflect.ValueOf(p).Elem()` — valor direccionable del struct apuntado por `p`.
    Enlace: https://pkg.go.dev/reflect#ValueOf
  - `reflect.TypeFor[personaStruct]()` — (forma genérica de obtener el `reflect.Type` para un tipo); también se pueden usar `reflect.TypeOf(personaStruct{})` o `reflect.TypeOf((*personaStruct)(nil)).Elem()`.
    Enlace: https://pkg.go.dev/reflect#TypeFor
  - `pt.NumField()` — número de campos en el tipo struct.
    Enlace: https://pkg.go.dev/reflect#Type.NumField
  - `field := pt.Field(i)` — obtiene la descripción del campo (nombre, tipo y tags).
    Enlace: https://pkg.go.dev/reflect#Type.Field
  - `field.Tag.Get("etiqueta")` — lectura de la tag con clave `etiqueta`.
    Enlace: https://pkg.go.dev/reflect#StructTag.Get
  - `pv.Field(i).String()` — obtiene el valor string del i-ésimo campo.
    Enlace: https://pkg.go.dev/reflect#Value.String
  - `fmt.Printf` y `println` — muestran la salida por consola.
    Enlace `fmt.Printf`: https://pkg.go.dev/fmt#Printf

Etiquetas de struct (struct tags)
- En el tipo `personaStruct` cada campo incluye una tag `etiqueta:"..."`. Estas tags se leen con `field.Tag.Get("etiqueta")` para mostrar textos personalizados junto a cada valor.

Consejos y riesgos comunes
- Asegúrese de pasar un puntero a `completaDatosStruct` y `getStructFields` para que `reflect.ValueOf(...).Elem()` funcione correctamente; si se pasa un valor no direccionable, las llamadas a `SetString` producirán panic.
- `fmt.Scanln` separa por espacios; si desea leer líneas completas con espacios, considere `bufio.NewReader(os.Stdin).ReadString('\n')` (fuera del alcance de este ejercicio).

Referencias generales
- Paquete reflect (documentación oficial): https://pkg.go.dev/reflect
- Paquete fmt (documentación oficial): https://pkg.go.dev/fmt

Fin de la documentación.
