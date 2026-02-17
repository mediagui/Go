// Indico el paquete al que pertenece el ejercicio.
package ejerciciosIntermedio

/*
###########
#ENUNCIADO#
###########

Tema 3 Ejercicio 6: Gestor de Contactos
Se pide desarrollar una aplicación que permita a los usuarios guardar y administrar contactos en un archivo. Hay que recordar manejar los posibles errores que puedan ocurrir.
*/

import (
	"bufio"         // Importo bufio para leer el texto desde la entrada por consola de forma más cómoda.
	"encoding/json" // Importo json para poder codificar y decodificar contactos en formato json.
	"fmt"           // Importo fmt para imprimir mensajes en la consola.
	"os"            // Importo os para interactuar con archivos y el sistema operativo.
)

// Estructura para los contatos. Indico que los archivos se van a guardar con un formato JSON
type Contact struct {
	Nombre string `json:"name"`  // Campo para almacenar el nombre del contacto
	Email  string `json:"email"` // Campo para almacenar el email del contacto
	Tfno   string `json:"tfno"`  // Campo para almacenar el teléfono del contacto
}

func ManejoErrores() {
	fmt.Println(`
###################
#MANEJO DE ERRORES#
###################
	`)
	// Creo un [slice] de tipo [Contact], para ir guradando cada uno de estos.
	var contacts []Contact

	// Llamo a la función que imprime la lista de contactos.
	err := cargarContactosArchivo(&contacts) // Intento cargar los contactos desde el archivo JSON, en caso de que ocurra un erro, lo imprimo.
	if err != nil {
		fmt.Println("ERROR al cargar los contactos")
	}

	//  Creo una instancia de [bufio], para leer la entrada por consola.
	rd := bufio.NewReader(os.Stdin)

	// Bucle infinito para mostrar el menú y ejecutar cada acción correspondiente.
	for {
		// Muestro el menú.
		fmt.Print(`=== GESTOR DE CONTACTOS ===
1. Añadir contacto
2. Mostrar todos los contactos
3. Modificar un contacto
4. Eliminar un contacto
5. Salir
============================
Escoja una opción: `)
		// Leo la opción del usuario.
		var opcion int
		_, err := fmt.Scan(&opcion)
		// En caso de que exista un error, lo indico.
		if err != nil {
			fmt.Println("ERROR al leer la opción: ", err)
			return
		}
		// Consumo un salto de línea para evitar que se salte un paso.
		rd.ReadString('\n')

		// Ejecuto la opción introducida por el usuario.
		switch opcion {
		case 1: // Ingresar y crear contacto.
			// Creo una instancia de Contact.
			var c Contact

			// Solicito cada dato del nuevo contacto.
			fmt.Print("Nombre: ")
			c.Nombre, _ = rd.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = rd.ReadString('\n')
			fmt.Print("Teléfono: ")
			c.Tfno, _ = rd.ReadString('\n')

			// Añado un contacto nuevo al slice.
			contacts = append(contacts, c)

			// Guardo la información en un archivo json.
			if err := guardarContactoArchivo(contacts); err != nil {
				// Intento guardar los contactos e ingormo de si hay algún fallo.
				fmt.Println("Error al guardar el contacto: ", err)
			}
		case 2: // Mostrar todos los contactos.
			fmt.Println("=====================================")
			// Recorro el slice de contactos.
			for index, contacto := range contacts {
				fmt.Printf("%d. Nombre: %s Email: %s Teléfono: %s\n", index+1, contacto.Nombre, contacto.Email, contacto.Tfno)
			}
			fmt.Println("=====================================")
		case 3: // Modificar un contacto.
			fmt.Print("Introduce el id del contacto a modificar: ")
			// Variable para almacenar el indice escogido por el usuario.
			var index int
			// Leo el número escrito por el usuario.
			_, err := fmt.Scan(&index)
			// Si hay un error lo indicio.
			if err != nil {
				fmt.Println("Error al leer el índice: ", err)
				break
			}
			// Limpio el buffer.
			rd.ReadString('\n')

			// Ajusto el índice introducido por el usuario, para que corresponda con el del slice.
			index--

			// Compruebo si el número introducido, está dentro de los límites de la lista.
			if index >= 0 && index < len(contacts) {
				// Muestro el nombre actual y solicito el nuevo.
				fmt.Print("Nombre actual [" + contacts[index].Nombre + "], nuevo nombre: ")
				// Actualizo el nombre.
				contacts[index].Nombre, _ = rd.ReadString('\n')

				// Muestro el email actual y solicito el nuevo.
				fmt.Print("Email actual [" + contacts[index].Email + "], nuevo email: ")
				// Actualizo el email.
				contacts[index].Email, _ = rd.ReadString('\n')

				// Muestro el teléfono actual y solicito el nuevo.
				fmt.Print("Teléfono actual [" + contacts[index].Tfno + "], nuevo teléfono: ")
				// Actualizo el teléfono.
				contacts[index].Tfno, _ = rd.ReadString('\n')

				// Intento guardar todos los cambios en el archivo.
				if err := guardarContactoArchivo(contacts); err != nil {
					// Si hay un fallo, lo aviso.
					fmt.Println("Error al guardar los cambios: ", err)
				} else {
					// Si no hay fallo, imprimo un mensaje de confirmación.
					fmt.Println("Contacto modificado con éxito.")
				}
			} else {
				// Si el índice introducido no existe, lo aviso.
				fmt.Println("Índice no válido")
			}
		case 4: // Eliminar un contacto.
			// Solicito el índice del contacto a eliminar.
			fmt.Print("Introduce el número del contacto a eliminar: ")
			var index int
			// Leo el número del contacto a eliminar.
			_, err := fmt.Scan(&index)
			// Verifico que no hubiese errores de lectura.
			if err != nil {
				fmt.Println("Error al leer el índice: ", err)
				break
			}
			// Limpio el buffer.
			rd.ReadString('\n')

			// Ajusto el índice a base 0.
			index--

			// Verifico que el índice exista en mi lista de contactos actual.
			if index >= 0 && index < len(contacts) {
				// Para eliminar, creo un nuevo slice uniendo lo que hay antes y después del contacto a eliminar.
				contacts = append(contacts[:index], contacts[index+1:]...)
				// Guardo inmediatamente la nueva lista, con el contacto eliminado.
				if err := guardarContactoArchivo(contacts); err != nil {
					// Si falla la escritura, lo notifico.
					fmt.Println("Error al guardar los cambios: ", err)
				} else {
					// Confirmamos que el contacto ha sido eliminado.
					fmt.Println("Contacto eliminado con éxito")
				}
			} else { // En caso de que el índice no sea válido, lo indico.
				fmt.Println("Índice no válido, compruebe la lista.")
			}
		case 5: // Salir del sistema.
			return
		default: // Opción no coincide con nigún caso.
			fmt.Println("Opción inválida")
		}
	}
}

// Función para añadir los contactos al archivo JSON
func guardarContactoArchivo(contacts []Contact) error {
	// Intento crear (abrir si ya existe) el archivo "contactos.json"
	file, err := os.Create("contactos.json")
	// Si ocurre algún error, lo indico.
	if err != nil {
		return err
	}
	// Programo el cierre del archivo.
	defer file.Close()

	// Creo un [encoder] específico que sabe escribir json en el archivo.
	encoder := json.NewEncoder(file)
	// Solicito al encoder que codifique en formato json la lista de contactos y la escriba en el archivo.
	err = encoder.Encode(contacts)
	// En caso de que ocurra un error, lo muestro.
	if err != nil {
		return err
	}
	// Si todo ha ido bien, devuelvo "nil" para indicar el éxito.
	return nil
}

// Función para cargar los contactos del archivo.
func cargarContactosArchivo(contactos *[]Contact) error {
	// Intento arbrir el archivo 'contactos.json' en modo lectura.
	file, err := os.Open("contactos.json")
	// Si el archivo no existe o no se puede abrir, devuelvo un error.
	if err != nil {
		return err
	}
	// Cierro el archivo.
	defer file.Close()

	//  Creo  un decoder que sabe leer json desde el archivo.
	decoder := json.NewDecoder(file)
	//  Le pido a decoder que lea el json y rellene la lista de contactos.
	err = decoder.Decode(&contactos)
	// Si el formato del archivo no es correcto o hay problemas de lectura, devuelvo un error.
	if err != nil {
		return err
	}

	// Si todo ha salido bien, devuelvo 'nil'
	return nil
}
