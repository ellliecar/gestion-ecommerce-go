package modules

import "gestion-ecommerce-go/utils"

// Product: Estructura de dato para productos ecuatorianos
type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
}

// ProductCatalog: MAP para búsqueda O(1) por ID ← Concepto Maps Unidad 1
type ProductCatalog map[string]Product

// === FUNCIONES PURAS - Sin efectos secundarios ===

// GetAllProducts: Retorna copia del catálogo (INMUTABILIDAD)
func GetAllProducts(catalog ProductCatalog) []Product {
	products := make([]Product, 0, len(catalog))
	for _, p := range catalog {
		products = append(products, p) // Copia por valor
	}
	return products
}

// FindByCategory: Filtra por categoría usando FilterBy (HIGHER-ORDER FUNCTION)
func FindByCategory(catalog ProductCatalog, category string) []Product {
	allProducts := GetAllProducts(catalog)
	return utils.FilterBy(allProducts, func(p Product) bool {
		return p.Category == category
	})
}

// CalculateInventoryValue: REDUCE funcional para valor total del inventario
func CalculateInventoryValue(catalog ProductCatalog) float64 {
	products := GetAllProducts(catalog)
	return utils.Reduce(products, 0.0, func(acc float64, p Product) float64 {
		return acc + (p.Price * float64(p.Stock))
	})
}

// AddProduct: Retorna NUEVO catálogo sin modificar original (INMUTABILIDAD)
func AddProduct(catalog ProductCatalog, newProduct Product) ProductCatalog {
	newCatalog := make(ProductCatalog, len(catalog)+1)
	for k, v := range catalog {
		newCatalog[k] = v
	}
	newCatalog[newProduct.ID] = newProduct
	return newCatalog
}

// UpdateStock: Actualiza stock retornando nuevo catálogo (patrón funcional)
func UpdateStock(catalog ProductCatalog, productID string, newStock int) ProductCatalog {
	newCatalog := make(ProductCatalog, len(catalog))
	for k, v := range catalog {
		if k == productID {
			updated := v
			updated.Stock = newStock
			newCatalog[k] = updated
		} else {
			newCatalog[k] = v
		}
	}
	return newCatalog
}
