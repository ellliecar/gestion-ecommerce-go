# Sistema de Gestión de E-commerce - Go + Programación Funcional

> Proyecto desarrollado para **Aprendizaje Autónomo 1 - Unidad 1**  
> Aplicando **Arrays, Slices y Maps de Go** + **Programación Funcional**

## 👤 Autor
**Elizabeth Cardona**  
Ingeniería en Ciberseguridad | Universidad Internacional del Ecuador (UIDE)  
Fecha: Mayo 2026

---

## 🎯 Objetivo General
Automatizar la gestión de una tienda online ecuatoriana permitiendo:
- CRUD de productos artesanales, café y chocolate
- Gestión de carrito con cálculo funcional de totales
- Reportes de ventas usando reduce funcional

---

## 📚 Conexión Explícita con Unidad 1: Arrays, Slices, Maps

| Concepto Unidad 1 | Sintaxis en Go | Aplicación en Proyecto | Beneficio Empresarial |
|------------------|----------------|----------------------|---------------------|
| **Array** `[n]T` | `[3]string{"A","B","C"}` | Categorías fijas de productos | Tipo seguro, tamaño conocido en compile-time |
| **Slice** `[]T` | `[]CartItem` | Carrito de compras dinámico | Flexibilidad para agregar/eliminar sin copiar arrays |
| **Map** `map[K]V` | `map[string]Product` | Catálogo para búsqueda por ID | Búsqueda O(1), escalable para miles de productos |

---

## 🔧 Tecnologías
- **Go 1.21+**
- **Enfoque funcional**: funciones puras, inmutabilidad, higher-order functions
- **Sin dependencias externas**: solo stdlib (`fmt`, `encoding/json`, `os`)

---

## 📦 Estructura del Proyecto
