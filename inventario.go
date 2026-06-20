package main

import "fmt"

// Inventario administra la lista de repuestos disponibles.
type Inventario struct {
	repuestos []Repuesto
}

// AgregarRepuesto incorpora un repuesto al inventario.
func (i *Inventario) AgregarRepuesto(repuesto Repuesto) {
	i.repuestos = append(i.repuestos, repuesto)
}

// MostrarInventario imprime todos los repuestos registrados.
func (i Inventario) MostrarInventario() {
	fmt.Println("Inventario disponible:")

	for _, repuesto := range i.repuestos {
		repuesto.Mostrar()
	}
}