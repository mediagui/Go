package ejerciciosIntermedio

// Importo únicamente el paquete fmt para la entrada y salida de datos, como has pedido.
import "fmt"

/*
###########
#ENUNCIADO#
###########

Tema 3 Ejercicio 4:

Fábrica de personitas
Crear una aplicación  en go, que permita crear objetos de tipos [Persona]. Después se mostrará cada uno de los objetos creados.

Los atributos para cada persona, serán:

v  Nombre.

v  Apellidos.

v  DNI.

v  Edad.

v  Estado civil.

v  Año de nacimiento.

Por cada Persona creada, se deberán solicitar los datos pertinentes.
*/

// Defino la estructura [Persona] con los campos solicitados.
type Persona struct {
	// Campo para del nombre de la persona
	Nombre string
	// Campo para los apellidos de la persona
	Apellidos string
	// Campo para el DNI/NIE de la persona
	Dni string
	// Campo para la edad de la persona
	Edad int
	// Campo para el estado civil de la persona
	EstadoCivil string
	// Campo para el año de nacimiento de la persona
	AnoNacimiento int
}

// Función principal del ejercicio, renombrada a CrearPersonitas.
func CrearPersonitas() {
	// Declaro un [slice] en el que almacenaré cada "persona" creada.
	var personas []Persona

	// Variable de control para el bucle.
	var salir bool = false
	var respuesta string

	// Inicio un bucle for para crear las  personitas.
	for !salir {
		// En cada iteración, creo una instancia de [Persona].
		var nuevaPersona Persona

		// Solicito los datos para crear a la persona.
		fmt.Println("_-= Nueva Persona =-_")

		// Solicito el nombre.
		fmt.Print("Introduce el nombre: ")
		fmt.Scan(&nuevaPersona.Nombre)

		// Solicito el primer apellido.
		fmt.Print("Introduce el primer apellido: ")
		fmt.Scan(&nuevaPersona.Apellidos)

		// Solicito el DNI.
		fmt.Print("Introduce el DNI: ")
		fmt.Scan(&nuevaPersona.Dni)

		// Solicito el DNI.
		fmt.Print("Introduce tu edad: ")
		fmt.Scan(&nuevaPersona.Edad)

		// Solicito el Estado Civil.
		fmt.Print("Introduce tu Estado Civil: ")
		fmt.Scan(&nuevaPersona.EstadoCivil)

		// Solicito el Año de Nacimiento.
		fmt.Print("Introduce tu año de nacimiento: ")
		fmt.Scan(&nuevaPersona.AnoNacimiento)

		// Añado la "nueva persona" al slice de [personas].
		personas = append(personas, nuevaPersona)

		fmt.Println("_-= Persona creada correctamente =-_")

		// Compruebo si el usuario desea crear otra persona.
		fmt.Print("¿Desea crear otra persona? (s/n): ")
		fmt.Scan(&respuesta)

		switch respuesta {
		case "s", "S":
			// Termino la ejecución
			salir = false
		case "n", "N":
			// Termino la ejecución
			salir = true
		default:
			fmt.Println("Opción, como tu opinión, inválida.")
		}
	}

	// Al salir del bucle, listo las personas creadas.
	fmt.Println(`
#############################
#Listado de Personas Creadas#
#############################
	`)

	// Bucle for para recorrer el listado de "personas".
	for i, p := range personas {
		// Muestro el número de "persona" (sumo 1 a [i] por que los ínidces empiezan en 0)
		fmt.Printf("Persona %d;\n", i+1)
		fmt.Printf("\n#######################################\n")
		// Imprimo los datos de cada persona.
		fmt.Printf("#\tNombre: %s\n", p.Nombre)
		fmt.Printf("#\tApellido: %s\n", p.Apellidos)
		fmt.Printf("#\tDNI: %s\n", p.Dni)
		fmt.Printf("#\tEdad: %d\n", p.Edad)
		fmt.Printf("#\tEstado Civil: %s\n", p.EstadoCivil)
		fmt.Printf("#\tAño de Nacimiento: %d\n", p.AnoNacimiento)
		fmt.Printf("#######################################\n")
	}
}
