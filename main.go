package main

import "fmt"

func main() {
	fmt.Println("Sistema de Gestión E-commerce Automotriz")
	fmt.Println("----------------------------------------")

	cliente := NewCliente("Wendy Quinga", "1720000000", "Hyundai Tucson")
	repuesto := NewRepuesto("Filtro de aceite", 25.50, 10)

	// Uso de la interfaz
	ImprimirElemento(cliente)
	ImprimirElemento(repuesto)

	cantidad := 2

	// Manejo de errores al verificar stock
	err := VerificarStock(repuesto, cantidad)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	orden := NewOrdenTrabajo(cliente, repuesto, cantidad)

	// Uso de la interfaz con orden de trabajo
	ImprimirElemento(orden)

	factura := NewFactura(orden)
	factura.GenerarFactura()
}