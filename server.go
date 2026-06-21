package main

import (
	"fmt"
	"net/http"
)

func IniciarServidor() {
	http.HandleFunc("/cliente", ClienteHandler)
	http.HandleFunc("/repuesto", RepuestoHandler)
	http.HandleFunc("/inventario", InventarioHandler)
	http.HandleFunc("/orden", OrdenHandler)
	http.HandleFunc("/factura", FacturaHandler)
	http.HandleFunc("/cotizacion", CotizacionHandler)
	http.HandleFunc("/servicio", ServicioHandler)
	http.HandleFunc("/resumen", ResumenHandler)

	fmt.Println("Servidor ejecutándose en http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error del servidor:", err)
	}
}