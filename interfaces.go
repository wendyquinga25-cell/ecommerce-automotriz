package main

// Mostrable define una interfaz para los elementos que pueden mostrar información.
type Mostrable interface {
	Mostrar()
}

// ImprimirElemento recibe cualquier elemento que implemente la interfaz Mostrable.
func ImprimirElemento(elemento Mostrable) {
	elemento.Mostrar()
}