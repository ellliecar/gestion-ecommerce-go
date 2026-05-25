# Sistema de Gestión de E-commerce - Go + Programación Funcional

Proyecto desarrollado aplicando **Arrays, Slices y Maps de Go** (Unidad 1) combinados con principios de **Programación Funcional** (funciones puras, inmutabilidad, higher-order functions).

## 🎯 Objetivo General
Automatizar la gestión de una tienda online ecuatoriana permitiendo:
- CRUD de productos (artesanales, café, chocolates)
- Gestión de carrito con cálculo funcional de totales
- Reportes de ventas usando reduce funcional

## 📚 Conexión con Unidad 1: Estructuras de Datos en Go

| Concepto Unidad 1 | Aplicación en el Proyecto | Beneficio |
|------------------|---------------------------|-----------|
| **Arrays** `[n]T` | Definición de categorías fijas: `[3]string{"Artesanías", "Café", "Chocolate"}` | Tipo seguro, tamaño conocido en compile-time |
| **Slices** `[]T` | Lista dinámica de productos y carrito de compras | Flexibilidad para agregar/eliminar sin copiar arrays |
| **Maps** `map[K]V` | Catálogo: `map[string]Product` para búsqueda O(1) por ID | Búsqueda eficiente sin recorrer listas |

## 🔧 Tecnologías
- Go 1.21+
- Enfoque funcional: funciones puras, inmutabilidad, composición
- Sin frameworks externos (solo stdlib: `fmt`, `json`, `sort`)

## 📦 Módulos Funcionales
1. `products.go`: CRUD con slices y maps
2. `cart.go`: Cálculo de totales con `reduce` funcional
3. `reports.go`: Agregaciones con funciones de orden superior

## 🚀 Ejecutar
```bash
go run main.go
