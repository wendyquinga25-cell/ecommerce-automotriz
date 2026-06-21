package main

import "fmt"

// Cliente representa a una persona que solicita productos o servicios.
type Cliente struct {
	Nombre   string `json:"nombre"`
	Cedula   string `json:"cedula"`
	Vehiculo string `json:"vehiculo"`
}

// NewCliente funciona como constructor para crear un nuevo cliente.
func NewCliente(nombre string, cedula string, vehiculo string) Cliente {
	return Cliente{
		Nombre:   nombre,
		Cedula:   cedula,
		Vehiculo: vehiculo,
	}
}

// GetNombre permite acceder al nombre del cliente.
func (c Cliente) GetNombre() string {
	return c.Nombre
}

// SetNombre permite modificar el nombre del cliente.
func (c *Cliente) SetNombre(nombre string) {
	c.Nombre = nombre
}

// Mostrar imprime la información del cliente.
func (c Cliente) Mostrar() {
	fmt.Println(
		"Cliente:", c.Nombre,
		"– Cédula:", c.Cedula,
		"– Vehículo:", c.Vehiculo,
	)
}