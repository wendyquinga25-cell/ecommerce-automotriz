package main

func main() {

	cliente := NewCliente("Wendy Quinga", "1720000000", "Hyundai Tucson")
	repuesto := NewRepuesto("Filtro de aceite", 25.50, 10)

	ImprimirElemento(cliente)
	ImprimirElemento(repuesto)

	orden := NewOrdenTrabajo(cliente, repuesto, 2)
	ImprimirElemento(orden)

	factura := NewFactura(orden)
	factura.GenerarFactura()

	IniciarServidor()
}