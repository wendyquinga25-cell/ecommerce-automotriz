# Sistema de Gestión E-commerce Automotriz

## Información General

**Proyecto:** Sistema de Gestión E-commerce Automotriz

**Aplicación:** Servicio automotriz especializado HYUNDAI y multimarca, orientado a la venta de repuestos y gestión de servicios de mantenimiento.

**Integrante:** Wendy Quinga

**Fecha:** Junio 2026

**Asignatura:** Programación Orientada a Objetos

**Carrera:** Ingeniería en Software

**Universidad:** Universidad Internacional del Ecuador - UIDE

---

## Objetivo del Proyecto

Desarrollar un sistema de gestión de e-commerce automotriz que permita registrar clientes, administrar repuestos, controlar inventario, generar cotizaciones, crear órdenes de trabajo y simular la facturación de servicios o productos, utilizando programación funcional y conceptos de Programación Orientada a Objetos en lenguaje Go.

---

## Descripción del Sistema

El sistema permite gestionar procesos básicos de un taller automotriz y venta de repuestos, facilitando el registro de información de clientes, productos y servicios, así como la generación de documentos de control para el seguimiento de trabajos realizados.

El proyecto fue desarrollado aplicando conceptos de estructuras, métodos, encapsulación, interfaces, manejo de errores, organización modular del código y servicios web mediante una API REST básica.

Además, el sistema permite visualizar información en formato JSON desde el navegador, utilizando endpoints que representan diferentes funcionalidades del aplicativo.

---

## Funcionalidades Implementadas

### Gestión de Clientes

* Registro de clientes.
* Almacenamiento de nombre, cédula y vehículo.
* Visualización de información del cliente.
* Consulta del cliente mediante servicio web.

### Gestión de Repuestos

* Registro de repuestos.
* Control de precios.
* Control de stock disponible.
* Validación de disponibilidad.
* Consulta de repuestos mediante servicio web.

### Gestión de Inventario

* Registro de repuestos en inventario.
* Consulta de productos disponibles.
* Visualización del inventario.
* Presentación del inventario en formato JSON.

### Órdenes de Trabajo

* Creación de órdenes de trabajo.
* Asociación entre cliente y repuesto.
* Control de cantidades solicitadas.
* Cálculo automático del valor total.
* Consulta de orden de trabajo mediante servicio web.

### Facturación

* Generación de factura básica.
* Cálculo de totales.
* Presentación de información de compra.
* Simulación de factura mediante endpoint web.

### Interfaces

* Implementación de interfaces para la impresión de elementos del sistema.
* Aplicación de polimorfismo.
* Reutilización de métodos para mostrar información de clientes, repuestos y órdenes.

### Manejo de Errores

* Validación de stock.
* Control de cantidades inválidas.
* Gestión de errores mediante la interfaz `error` de Go.
* Uso de mensajes de error para evitar operaciones incorrectas.

### Servicios Web

* Implementación de API REST básica.
* Uso del paquete `net/http`.
* Serialización de datos mediante JSON.
* Consulta de información desde el navegador web.

---

## Tecnologías Utilizadas

* Lenguaje de Programación: Go / Golang
* Visual Studio Code
* Git
* GitHub
* Programación Orientada a Objetos
* Programación Funcional
* API REST
* JSON
* Paquete `net/http`
* Paquete `encoding/json`
* Navegador web Firefox
* Sistema operativo macOS

---

## Conceptos Aplicados

Durante el desarrollo del proyecto se aplicaron los siguientes conceptos:

* Estructuras de datos en Go.
* Métodos asociados a estructuras.
* Constructores.
* Métodos getter y setter.
* Encapsulación.
* Interfaces.
* Polimorfismo.
* Manejo de errores.
* Validación de datos.
* Organización modular del código.
* Servicios web.
* Serialización JSON.
* Control de versiones con Git y GitHub.

---

## Estructura del Proyecto

```text
ecommerce-automotriz/
│
├── main.go
├── clientes.go
├── repuestos.go
├── inventario.go
├── cotizacion.go
├── ordenes.go
├── facturacion.go
├── interfaces.go
├── api.go
├── server.go
├── go.mod
└── README.md
```

---

## Descripción de Archivos Principales

### `main.go`

Archivo principal del sistema. Ejecuta las funcionalidades generales del programa e inicia el servidor web local.

### `clientes.go`

Define la estructura `Cliente`, sus atributos, constructor y métodos para acceder o modificar información del cliente.

### `repuestos.go`

Define la estructura `Repuesto`, sus atributos, constructor, métodos de consulta, validación de stock y manejo de errores.

### `inventario.go`

Administra la lista de repuestos disponibles mediante una estructura de inventario.

### `cotizacion.go`

Contiene funciones relacionadas con el cálculo de valores y cotizaciones de productos o servicios.

### `ordenes.go`

Permite crear una orden de trabajo relacionando cliente, repuesto, cantidad y total.

### `facturacion.go`

Simula la generación de una factura a partir de una orden de trabajo.

### `interfaces.go`

Contiene la implementación de interfaces para imprimir información de diferentes elementos del sistema.

### `api.go`

Contiene los controladores de los servicios web y permite devolver información en formato JSON.

### `server.go`

Configura y ejecuta el servidor web local en el puerto 8080.

---

## Servicios Web Implementados

El sistema cuenta con los siguientes endpoints disponibles:

| Endpoint      | Descripción                                                           |
| ------------- | --------------------------------------------------------------------- |
| `/cliente`    | Muestra la información de un cliente registrado.                      |
| `/repuesto`   | Muestra la información de un repuesto disponible.                     |
| `/inventario` | Muestra una lista de repuestos disponibles en inventario.             |
| `/orden`      | Muestra una orden de trabajo con cliente, repuesto, cantidad y total. |
| `/factura`    | Muestra la simulación de una factura generada correctamente.          |

---

## Ejemplos de Endpoints

### Cliente

```text
http://localhost:8080/cliente
```

### Repuesto

```text
http://localhost:8080/repuesto
```

### Inventario

```text
http://localhost:8080/inventario
```

### Orden de Trabajo

```text
http://localhost:8080/orden
```

### Factura

```text
http://localhost:8080/factura
```

---

## Ejemplo de Respuesta JSON

Ejemplo de respuesta del endpoint `/factura`:

```json
{
  "cantidad": 2,
  "cedula": "1720000000",
  "cliente": "Wendy Quinga",
  "mensaje": "Factura generada correctamente",
  "precio": 25.5,
  "producto": "Filtro de aceite",
  "total": 51,
  "vehiculo": "Hyundai Tucson"
}
```

---

## Cómo Ejecutar el Proyecto

### 1. Clonar el repositorio

```bash
git clone https://github.com/wendyquinga25-cell/ecommerce-automotriz.git
```

### 2. Ingresar al directorio

```bash
cd ecommerce-automotriz
```

### 3. Ejecutar el proyecto

```bash
go run .
```

Después de ejecutar el comando, el sistema mostrará en la terminal:

```text
Servidor ejecutándose en http://localhost:8080
```

Posteriormente, se pueden abrir los endpoints en el navegador web utilizando las rutas correspondientes.

---

## Resultados Alcanzados

Durante el desarrollo del proyecto se implementaron conceptos fundamentales de Programación Orientada a Objetos en Go, incluyendo:

* Encapsulación.
* Interfaces.
* Métodos.
* Constructores.
* Manejo de errores.
* Organización modular del código.
* Gestión de clientes.
* Gestión de repuestos.
* Gestión de inventario.
* Generación de órdenes de trabajo.
* Facturación básica.
* Implementación de servicios web.
* Serialización de información en formato JSON.
* Publicación del código en GitHub.

El sistema constituye una base funcional para futuras mejoras como la implementación de una interfaz web, conexión con base de datos, autenticación de usuarios, generación de reportes y comercio electrónico en línea.

---

## Evidencias de Funcionamiento

El sistema fue probado localmente mediante el navegador web, accediendo a los endpoints implementados en el servidor local.

Los endpoints `/cliente`, `/repuesto`, `/inventario`, `/orden` y `/factura` permiten visualizar información del sistema en formato JSON.

---

## Conclusión

El desarrollo de este proyecto permitió aplicar los conocimientos adquiridos en la asignatura de Programación Orientada a Objetos, integrando estructuras, métodos, interfaces, encapsulación, manejo de errores y servicios web en lenguaje Go.

El sistema desarrollado representa una solución básica para la gestión de un servicio automotriz especializado, permitiendo registrar información, administrar repuestos, controlar inventario, generar órdenes de trabajo y simular facturación.

Además, la implementación de una API REST permitió ampliar el alcance del proyecto, facilitando la consulta de información desde el navegador mediante respuestas en formato JSON.

Este proyecto demuestra cómo el lenguaje Go puede ser utilizado para construir soluciones organizadas, funcionales y orientadas a necesidades reales de gestión empresarial.

---

## Repositorio

https://github.com/wendyquinga25-cell/ecommerce-automotriz

---

## Autora

**Wendy Quinga**

Ingeniería en Software

Universidad Internacional del Ecuador - UIDE

Junio 2026
