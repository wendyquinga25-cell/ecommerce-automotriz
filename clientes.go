package main

import "fmt"

// Cliente representa a una persona que solicita productos o servicios.
type Cliente struct {
	nombre   string
	cedula   string
	vehiculo string
}

// NewCliente funciona como constructor para crear un nuevo cliente.
func NewCliente(nombre string, cedula string, vehiculo string) Cliente {
	return Cliente{
		nombre:   nombre,
		cedula:   cedula,
		vehiculo: vehiculo,
	}
}

// GetNombre permite acceder al nombre del cliente.
func (c Cliente) GetNombre() string {
	return c.nombre
}

// SetNombre permite modificar el nombre del cliente.
func (c *Cliente) SetNombre(nombre string) {
	c.nombre = nombre
}

// Mostrar imprime la información del cliente.
func (c Cliente) Mostrar() {
	fmt.Println("Cliente:", c.nombre, "- Cédula:", c.cedula, "- Vehículo:", c.vehiculo)
}