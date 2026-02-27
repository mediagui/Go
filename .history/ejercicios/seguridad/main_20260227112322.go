package main

// Las ctes en Go, por convención no se deben escribir en mayúsculas ya que
// las convertimos en públicas automaticamente
// para evitar el acceso desde fuera de la aplicación, habría que definirlas
// en un módulo llamado "internal"
const FILE_NAME = "auditoria.txt"

type Acceso struct {
	Id        uint
	Usuario   string
	Riesgo    int
	Categoria string
}

// Crea e Inicializa un array con 5 elementos de tipo Acceso
var eventos [5]Acceso = [5]Acceso{
	{Id: 1, Usuario: "admin", Riesgo: 0, Categoria: "Administrador"},
	{Id: 2, Usuario: "usuario1", Riesgo: 1, Categoria: "Usuario Regular"},
	{Id: 3, Usuario: "usuario2", Riesgo: 2, Categoria: "Usuario Regular"},
	{Id: 4, Usuario: "usuario3", Riesgo: 3, Categoria: "Usuario Regular"},
	{Id: 5, Usuario: "usuario4", Riesgo: 4, Categoria: "Usuario Regular"},
}

func main() {

}
