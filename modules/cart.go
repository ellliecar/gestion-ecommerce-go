package modules

import "gestion-ecommerce-go/utils"

// CartItem: Elemento del carrito
type CartItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     float64 `json:"price"` // Precio al momento de agregar (inmutabilidad histórica)
}

// ShoppingCart: Slice dinámico de items (concepto Slices de Unidad 1)
type ShoppingCart []CartItem

// === FUNCIONES PURAS ===

// AddToCart: Retorna nuevo carrito con item agregado (sin mutar original)
func AddToCart(cart ShoppingCart, item CartItem) ShoppingCart {
	newCart := make(ShoppingCart, len(cart)+1)
	copy(newCart, cart) // Copia segura del slice original
	newCart[len(cart)] = item
	return newCart
}

// CalculateTotal: Reduce funcional para calcular total del carrito
func CalculateTotal(cart ShoppingCart) float64 {
	return utils.Reduce(cart, 0.0, func(total float64, item CartItem) float64 {
		return total + (item.Price * float64(item.Quantity))
	})
}

// GetItemsByCategory: Filtra items del carrito por categoría de producto
// Usa composición: MapTransform + FilterBy
func GetItemsByCategory(cart ShoppingCart, products ProductCatalog, category string) []CartItem {
	// Transformar cart items a productos completos
	productItems := utils.MapTransform(cart, func(item CartItem) Product {
		return products[item.ProductID]
	})
	
	// Filtrar por categoría
	filtered := utils.FilterBy(productItems, func(p Product) bool {
		return p.Category == category
	})
	
	// Convertir de vuelta a CartItem (simplificado para el ejemplo)
	result := make([]CartItem, len(filtered))
	for i, p := range filtered {
		result[i] = CartItem{
			ProductID: p.ID,
			Quantity:  1, // Simplificación
			Price:     p.Price,
		}
	}
	return result
}
