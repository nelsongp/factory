package main

import (
	"fmt"
	"github.com/nelsongp/factory/factory"
	"os"
)

func main() {
	var t int
	fmt.Print("Digite la conexion deseada: ")
	_, err := fmt.Scan(&t)
	if err != nil {
		fmt.Printf("error al leer la opcion: %v", err)
		os.Exit(1)
	}

	conn := factory.Factory(t)
	if conn == nil {
		fmt.Println("Motor no valido")
		os.Exit(1)
	}

	err = conn.Connect()
	if err != nil {
		fmt.Println("No se pudo realizar la conexion")
	}

	now, err := conn.GetNow()
	if err != nil {
		fmt.Printf("No se pudo obtener la fecha de la base %v", err)
	}

	fmt.Println(now.Format("2006-01-02"))

	err = conn.Close()
	if err != nil {
		fmt.Printf("No se pudo cerrar la conexion %v", err)
	}
}
