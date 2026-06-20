package main

import "fmt"

// OrdenTrabajo representa una solicitud de mantenimiento o compra.
type OrdenTrabajo struct {
	cliente  Cliente
	repuesto Repuesto
	cantidad int
	total    float64
}

// NewOrdenTrabajo crea una orden de trabajo.
func NewOrdenTrabajo(cliente Cliente, repuesto Repuesto, cantidad int) OrdenTrabajo {
	total := CalcularTotal(repuesto.GetPrecio(), cantidad)

	return OrdenTrabajo{
		cliente:  cliente,
		repuesto: repuesto,
		cantidad: cantidad,
		total:    total,
	}
}

// Mostrar imprime la orden de trabajo.
func (o OrdenTrabajo) Mostrar() {
	fmt.Println("Orden de trabajo")
	fmt.Println("Cliente:", o.cliente.GetNombre())
	fmt.Println("Repuesto:", o.repuesto.GetNombre())
	fmt.Println("Cantidad:", o.cantidad)
	fmt.Println("Total:", o.total)
}