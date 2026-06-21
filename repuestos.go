package main

import (
	"errors"
	"fmt"
)

type Repuesto struct {
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
	Stock  int     `json:"stock"`
}

func NewRepuesto(nombre string, precio float64, stock int) Repuesto {
	return Repuesto{
		Nombre: nombre,
		Precio: precio,
		Stock:  stock,
	}
}

func (r Repuesto) GetNombre() string {
	return r.Nombre
}

func (r Repuesto) GetPrecio() float64 {
	return r.Precio
}

func (r Repuesto) GetStock() int {
	return r.Stock
}

func (r Repuesto) ValidarStock() error {
	if r.Stock <= 0 {
		return errors.New("no existe stock disponible del repuesto")
	}
	return nil
}

func VerificarStock(repuesto Repuesto, cantidad int) error {
	if cantidad <= 0 {
		return errors.New("la cantidad solicitada debe ser mayor a cero")
	}

	if cantidad > repuesto.Stock {
		return errors.New("stock insuficiente para realizar la orden")
	}

	return nil
}

func (r Repuesto) Mostrar() {
	fmt.Println("Repuesto:", r.Nombre, "- Precio:", r.Precio, "- Stock:", r.Stock)
}