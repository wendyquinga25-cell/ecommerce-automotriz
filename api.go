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

// Endpoint orden de trabajo
func OrdenHandler(w http.ResponseWriter, r *http.Request) {
	cliente := NewCliente("Wendy Quinga", "1720000000", "Hyundai Tucson")
	repuesto := NewRepuesto("Filtro de aceite", 25.50, 10)
	cantidad := 2

	orden := NewOrdenTrabajo(cliente, repuesto, cantidad)

	respuesta := map[string]interface{}{
		"cliente":  cliente.Nombre,
		"cedula":   cliente.Cedula,
		"vehiculo": cliente.Vehiculo,
		"repuesto": repuesto.Nombre,
		"precio":   repuesto.Precio,
		"cantidad": cantidad,
		"total":    orden.total,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// Endpoint factura
func FacturaHandler(w http.ResponseWriter, r *http.Request) {
	cliente := NewCliente("Wendy Quinga", "1720000000", "Hyundai Tucson")
	repuesto := NewRepuesto("Filtro de aceite", 25.50, 10)
	cantidad := 2

	orden := NewOrdenTrabajo(cliente, repuesto, cantidad)
	factura := NewFactura(orden)
	factura.GenerarFactura()

	respuesta := map[string]interface{}{
		"mensaje":  "Factura generada correctamente",
		"cliente":  cliente.Nombre,
		"cedula":   cliente.Cedula,
		"vehiculo": cliente.Vehiculo,
		"producto": repuesto.Nombre,
		"precio":   repuesto.Precio,
		"cantidad": cantidad,
		"total":    orden.total,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// Endpoint cotización
func CotizacionHandler(w http.ResponseWriter, r *http.Request) {
	precio := 25.50
	cantidad := 2
	total := precio * float64(cantidad)

	respuesta := map[string]interface{}{
		"cliente":     "Wendy Quinga",
		"vehiculo":    "Hyundai Tucson",
		"producto":    "Filtro de aceite",
		"precio":      precio,
		"cantidad":    cantidad,
		"total":       total,
		"estado":      "Cotización generada",
		"descripcion": "Cotización básica para venta de repuesto automotriz",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// Endpoint servicio
func ServicioHandler(w http.ResponseWriter, r *http.Request) {
	respuesta := map[string]interface{}{
		"servicio":     "Mantenimiento preventivo",
		"vehiculo":     "Hyundai Tucson",
		"descripcion":  "Cambio de aceite, revisión de filtros y control general del vehículo",
		"duracion":     "2 horas",
		"estado":       "Disponible",
		"especialidad": "Servicio automotriz HYUNDAI y multimarca",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// Endpoint resumen del sistema
func ResumenHandler(w http.ResponseWriter, r *http.Request) {
	respuesta := map[string]interface{}{
		"proyecto":       "Sistema de Gestión E-commerce Automotriz",
		"integrante":     "Wendy Quinga",
		"lenguaje":       "Go",
		"serializacion":  "JSON",
		"servicios_web":  8,
		"funcionalidades": []string{
			"Gestión de clientes",
			"Gestión de repuestos",
			"Inventario",
			"Cotización",
			"Orden de trabajo",
			"Facturación",
			"Servicios automotrices",
			"Resumen del sistema",
		},
		"estado": "Proyecto funcional",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}