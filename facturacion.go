package main

import "fmt"

// Facturable es una interfaz que obliga a implementar el método GenerarFactura.
type Facturable interface {
	GenerarFactura()
}

// Factura representa el cierre de una venta o servicio.
type Factura struct {
	orden OrdenTrabajo
}

// NewFactura crea una nueva factura.
func NewFactura(orden OrdenTrabajo) Factura {
	return Factura{
		orden: orden,
	}
}

// GenerarFactura imprime la factura del servicio.
func (f Factura) GenerarFactura() {
	fmt.Println("Factura generada correctamente")
	fmt.Println("Cliente:", f.orden.cliente.GetNombre())
	fmt.Println("Producto:", f.orden.repuesto.GetNombre())
	fmt.Println("Total a pagar:", f.orden.total)
}