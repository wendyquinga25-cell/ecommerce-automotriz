package main

import (
	"encoding/json"
	"net/http"
)

// Endpoint cliente
func ClienteHandler(w http.ResponseWriter, r *http.Request) {
	cliente := Cliente{
		Nombre:   "Wendy Quinga",
		Cedula:   "1720000000",
		Vehiculo: "Hyundai Tucson",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cliente)
}

// Endpoint repuesto
func RepuestoHandler(w http.ResponseWriter, r *http.Request) {
	repuesto := Repuesto{
		Nombre: "Filtro de aceite",
		Precio: 25.50,
		Stock:  10,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repuesto)
}

// Endpoint inventario
func InventarioHandler(w http.ResponseWriter, r *http.Request) {
	inventario := Inventario{}

	repuesto1 := NewRepuesto("Filtro de aceite", 25.50, 10)
	repuesto2 := NewRepuesto("Pastillas de freno", 45.00, 6)
	repuesto3 := NewRepuesto("Bujías", 18.75, 12)

	inventario.AgregarRepuesto(repuesto1)
	inventario.AgregarRepuesto(repuesto2)
	inventario.AgregarRepuesto(repuesto3)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventario)
}