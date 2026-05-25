package modules

import "gestion-ecommerce-go/utils"

// CartItem: Elemento individual del carrito
type CartItem struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"` // Precio congelado al agregar (inmutabilidad histórica)
}

// ShoppingCart: SLICE dinámico de items ← Concepto Slices Unidad 1
type ShoppingCart []CartItem

// === FUNCIONES PURAS ===

// AddToCart: Retorna NUEVO carrito con item agregado (sin mutar original)
func AddToCart(cart ShoppingCart, item CartItem) ShoppingCart {
	newCart := make(ShoppingCart, len(cart)+1)
	copy(newCart, cart) // Copia segura del slice original
	newCart[len(cart)] = item
	return newCart
}

// CalculateTotal: REDUCE funcional para calcular total del carrito
func CalculateTotal(cart ShoppingCart) float64 {
	return utils.Reduce(cart, 0.0, func(total float64, item CartItem) float64 {
		return total + (item.Price * float64(item.Quantity))
	})
}

// GetCount: Retorna número de items en el carrito
func GetCount(cart ShoppingCart) int {
	return len(cart)
}

// GetItemsByCategory: Filtra items del carrito por categoría (COMPOSICIÓN)
func GetItemsByCategory(cart ShoppingCart, products ProductCatalog, category string) []CartItem {
	// Transformar cart items a productos completos
	productItems := utils.MapTransform(cart, func(item CartItem) Product {
		return products[item.ProductID]
	})
	
	// Filtrar por categoría
	filtered := utils.FilterBy(productItems, func(p Product) bool {
		return p.Category == category
	})
	
	// Convertir de vuelta a CartItem
	result := make([]CartItem, len(filtered))
	for i, p := range filtered {
		result[i] = CartItem{
			ProductID: p.ID,
			Quantity:  1,
			Price:     p.Price,
		}
	}
	return result
}
