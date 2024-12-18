package v1

import (
	"bufio"
	"fmt"
	"mysql_driver_test/internal/connection"
	"mysql_driver_test/internal/modelo"
	"os"

	"github.com/rs/zerolog/log"
)

// Listar lista todos los clientes de la base de datos.
func Listar() {

	connection.Conectar()

	sql := "SELECT id, nombre, correo, telefono, mensaje FROM clientes order by id desc;"
	datos, err := connection.Db.Query(sql)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Error al listar los clientes")
	}
	defer connection.CerrarConexion()

	/**
	 * mostrar clientes
	clientes := modelo.Clientes{}
	for datos.Next() {
		dato := modelo.Cliente{}
		datos.Scan(&dato.ID, &dato.Nombre, &dato.Correo, &dato.Telefono, &dato.Mensaje)
		clientes = append(clientes, dato)
	}
	fmt.Println(clientes)
	*/

	fmt.Println("\n Listado de clientes organizaco")
	fmt.Println("----------------------------")
	for datos.Next() {
		var dato = modelo.Cliente{}
		err := datos.Scan(&dato.ID, &dato.Nombre, &dato.Correo, &dato.Telefono, &dato.Mensaje)
		if err != nil {
			log.Logger.Error().Err(err).Msg("Error al listar los clientes")
		}
		fmt.Printf("ID: %d\nNombre: %s\nCorreo: %s\nTelefono: %s\nMensaje: %s\n\n", dato.ID, dato.Nombre, dato.Correo, dato.Telefono, dato.Mensaje)
	}
}

// ListarPorId lista un cliente por id.
func ListarPorId(id int) {
	connection.Conectar()

	sql := "SELECT id, nombre, correo, telefono, mensaje FROM clientes where id=?;"
	datos, err := connection.Db.Query(sql, id)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Error al listar los clientes")
	}
	defer connection.CerrarConexion()

	fmt.Println("\n Listado de clientes por id")
	fmt.Println("----------------------------")
	for datos.Next() {
		var dato = modelo.Cliente{}
		err := datos.Scan(&dato.ID, &dato.Nombre, &dato.Correo, &dato.Telefono, &dato.Mensaje)
		if err != nil {
			log.Logger.Error().Err(err).Msg("Error al listar los clientes")
		}
		fmt.Printf("ID: %d\nNombre: %s\nCorreo: %s\nTelefono: %s\nMensaje: %s\n\n", dato.ID, dato.Nombre, dato.Correo, dato.Telefono, dato.Mensaje)
		fmt.Println("----------------------------")
	}
}

// Insertar inserta un cliente en la base de datos.
func Insertar(cliente modelo.Cliente) {
	connection.Conectar()

	sql := "INSERT INTO clientes (nombre, correo, telefono, mensaje) VALUES (?, ?, ?, ?);"
	result, err := connection.Db.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono, cliente.Mensaje)
	fmt.Println(result)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Error al insertar un cliente")
	}

	fmt.Println("Cliente insertado correctamente")

}

// Editar modifica un cliente en la base de datos.
func Editar(id int, cliente modelo.Cliente) {
	connection.Conectar()

	sql := "UPDATE clientes SET nombre=?, correo=?, telefono=?, mensaje=? WHERE id=?;"
	result, err := connection.Db.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono, cliente.Mensaje, id)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Error al editar un cliente")
	}
	fmt.Println(result)
	fmt.Println("Cliente editado correctamente")
}

// Eliminar elimina un cliente de la base de datos.
func Eliminar(id int) {
	connection.Conectar()

	sql := "DELETE FROM clientes WHERE id=?;"
	_, err := connection.Db.Exec(sql, id)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Error al eliminar un cliente")
	}

	fmt.Println("Cliente eliminado correctamente")
}

//funciones de trabajo

var ID int
var Nombre, correo, telefono, mensaje string

// Ejecutar ejecuta la aplicación.
func Ejecutar() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("---------------------------------")
		fmt.Println("\nSeleccione una opción: \n")
		fmt.Println("1. Listar todos los registros")
		fmt.Println("2. Listar un registro por id")
		fmt.Println("3. Insertar un registro")
		fmt.Println("4. Editar un registro")
		fmt.Println("5. Eliminar un registro")
		fmt.Println("6. Salir")

		if scanner.Scan() {
			opcion := scanner.Text()
			switch opcion {
			case "1":
				Listar()
			case "2":
				fmt.Println("Ingrese el id del registro a buscar: ")
				var ID int
				fmt.Scan(&ID)
				ListarPorId(ID)
			case "3":
				fmt.Println("Ingrese el nombre del cliente: ")
				var Nombre, correo, telefono, mensaje string
				fmt.Scan(&Nombre)
				fmt.Println("Ingrese el correo del cliente: ")
				fmt.Scan(&correo)
				fmt.Println("Ingrese el teléfono del cliente: ")
				fmt.Scan(&telefono)
				fmt.Println("Ingrese el mensaje del cliente: ")
				fmt.Scan(&mensaje)
				cliente := modelo.Cliente{Nombre: Nombre, Correo: correo, Telefono: telefono, Mensaje: mensaje}
				Insertar(cliente)
			case "4":
				fmt.Println("Ingrese el id del registro a editar: ")
				var ID int
				fmt.Scan(&ID)
				fmt.Println("Ingrese el nombre del cliente: ")
				var Nombre, correo, telefono, mensaje string
				fmt.Scan(&Nombre)
				fmt.Println("Ingrese el correo del cliente: ")
				fmt.Scan(&correo)
				fmt.Println("Ingrese el teléfono del cliente: ")
				fmt.Scan(&telefono)
				fmt.Println("Ingrese el mensaje del cliente: ")
				fmt.Scan(&mensaje)
				cliente := modelo.Cliente{Nombre: Nombre, Correo: correo, Telefono: telefono, Mensaje: mensaje}
				Editar(ID, cliente)
			case "5":
				fmt.Println("Ingrese el id del registro a eliminar: ")
				var ID int
				fmt.Scan(&ID)
				Eliminar(ID)
			case "6":
				fmt.Println("Saliendo...")
				return
			default:
				fmt.Println("Opción no válida, por favor intente de nuevo.")
			}
		} else {
			fmt.Println("Error al leer la entrada. Por favor, intente de nuevo.")
		}
	}
}
