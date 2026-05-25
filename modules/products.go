package modules

import (
	"gestion-ecommerce-go/utils"
)

// Product: Estructura de dato para productos ecuatorianos
type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Category    string  `json:"category"` // Array fijo: [3]string{"Artesanías", "Café", "Chocolate"}
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
}

// ProductCatalog: Map para búsqueda O(1) por ID (concepto Maps de Unidad 1)
type ProductCatalog map[string]Product

// === FUNCIONES PURAS ===

// GetAllProducts: Retorna copia del catálogo (inmutabilidad)
func GetAllProducts(catalog ProductCatalog) []Product {
	products := make([]Product, 0, len(catalog))
	for _, p := range catalog {
		products = append(products, p) // Copia por valor
	}
	return products
}

// FindByCategory: Filtra productos por categoría usando FilterBy (Higher-Order Function)
func FindByCategory(catalog ProductCatalog, category string) []Product {
	allProducts := GetAllProducts(catalog)
	return utils.FilterBy(allProducts, func(p Product) bool {
		return p.Category == category
	})
}

// CalculateInventoryValue: Reduce funcional para calcular valor total del inventario
func CalculateInventoryValue(catalog ProductCatalog) float64 {
	products := GetAllProducts(catalog)
	return utils.Reduce(products, 0.0, func(acc float64, p Product) float64 {
		return acc + (p.Price * float64(p.Stock))
	})
}

// AddProduct: Retorna nuevo catálogo sin modificar el original (inmutabilidad)
func AddProduct(catalog ProductCatalog, newProduct Product) ProductCatalog {
	// Crear nuevo map (no modificar el original)
	newCatalog := make(ProductCatalog, len(catalog)+1)
	for k, v := range catalog {
		newCatalog[k] = v // Copia por valor
	}
	newCatalog[newProduct.ID] = newProduct
	return newCatalog
}

// UpdateStock: Actualiza stock retornando nuevo catálogo (patrón funcional)
func UpdateStock(catalog ProductCatalog, productID string, newStock int) ProductCatalog {
	newCatalog := make(ProductCatalog, len(catalog))
	for k, v := range catalog {
		if k == productID {
			// Crear nueva instancia con stock actualizado (inmutabilidad)
			updated := v
			updated.Stock = newStock
			newCatalog[k] = updated
		} else {
			newCatalog[k] = v
		}
	}
	return newCatalog
}
