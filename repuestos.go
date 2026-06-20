package main

import (
	"errors"
	"fmt"
)

// Repuesto representa un producto automotriz disponible para la venta.
type Repuesto struct {
	nombre string
	precio float64
	stock  int
}

// NewRepuesto funciona como constructor para crear un repuesto.
func NewRepuesto(nombre string, precio float64, stock int) Repuesto {
	return Repuesto{
		nombre: nombre,
		precio: precio,
		stock:  stock,
	}
}

// GetNombre permite obtener el nombre del repuesto.
func (r Repuesto) GetNombre() string {
	return r.nombre
}

// GetPrecio permite obtener el precio del repuesto.
func (r Repuesto) GetPrecio() float64 {
	return r.precio
}

// GetStock permite obtener el stock del repuesto.
func (r Repuesto) GetStock() int {
	return r.stock
}

// ValidarStock verifica que exista stock disponible.
func (r Repuesto) ValidarStock() error {
	if r.stock <= 0 {
		return errors.New("no existe stock disponible del repuesto")
	}
	return nil
}

// VerificarStock valida si la cantidad solicitada está disponible.
func VerificarStock(repuesto Repuesto, cantidad int) error {
	if cantidad <= 0 {
		return errors.New("la cantidad solicitada debe ser mayor a cero")
	}

	if cantidad > repuesto.stock {
		return errors.New("stock insuficiente para realizar la orden")
	}

	return nil
}

// Mostrar imprime la información del repuesto.
func (r Repuesto) Mostrar() {
	fmt.Println("Repuesto:", r.nombre, "- Precio:", r.precio, "- Stock:", r.stock)
}