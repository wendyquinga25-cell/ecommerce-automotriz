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