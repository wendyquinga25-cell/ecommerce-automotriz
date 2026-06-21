package main

import "fmt"

// Inventario administra la lista de repuestos disponibles.
type Inventario struct {
	Repuestos []Repuesto `json:"repuestos"`
}

// AgregarRepuesto incorpora un repuesto al inventario.
func (i *Inventario) AgregarRepuesto(repuesto Repuesto) {
	i.Repuestos = append(i.Repuestos, repuesto)
}

// MostrarInventario imprime todos los repuestos registrados.
func (i Inventario) MostrarInventario() {
	fmt.Println("Inventario disponible:")

	for _, repuesto := range i.Repuestos {
		repuesto.Mostrar()
	}
}