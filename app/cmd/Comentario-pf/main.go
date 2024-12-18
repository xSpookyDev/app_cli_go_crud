package main

import (
	v1 "mysql_driver_test/api/v1"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339, NoColor: false})

	//v1.Listar() // Descomentar para listar todos los registros
	//v1.ListarPorId(2) // Descomentar para listar un registro por id
	//cliente := modelo.Cliente{Nombre: "Juanes", Correo: "faselos@gmail.com", Telefono: "123456780", Mensaje: "Hola como estas???"}
	//v1.Insertar(cliente)  Descomentar para insertar un registro
	//v1.Editar(3, cliente) // Descomentar para editar un registro

	//v1.Eliminar(3) // Descomentar para eliminar un registro
	//v1.Listar()
	v1.Ejecutar()
	log.Info().Msg("Fin de la app")
}

//Funciones de trabajo
