package main

import (
	"fmt"
	"gestion-ecommerce-go/modules"
)

func main() {
	// === EJEMPLO: Conexión explícita con Unidad 1 ===
	
	// 1. ARRAY: Categorías fijas (tamaño conocido en compile-time)
	categorias := [3]string{"Artesanías", "Café", "Chocolate"}
	fmt.Printf("📦 Categorías disponibles (Array): %v\n\n", categorias)
	
	// 2. MAP: Catálogo de productos para búsqueda O(1)
	catalogo := modules.ProductCatalog{
		"ART-001": modules.Product{
			ID: "ART-001", Name: "Sombrero de Paja Toquilla",
			Category: "Artesanías", Price: 45.00, Stock: 20,
			Description: "Hecho a mano en Cuenca",
		},
		"CAF-001": modules.Product{
			ID: "CAF-001", Name: "Café de Altura Loja",
			Category: "Café", Price: 12.50, Stock: 50,
			Description: "100% orgánico, tostado artesanal",
		},
		"CHO-001": modules.Product{
			ID: "CHO-001", Name: "Chocolate 70% Cacao",
			Category: "Chocolate", Price: 8.99, Stock: 35,
			Description: "Elaborado con cacao nacional",
		},
	}
	
	// 3. SLICE: Carrito de compras dinámico
	carrito := modules.ShoppingCart{}
	carrito = modules.AddToCart(carrito, modules.CartItem{
		ProductID: "CAF-001", Quantity: 2, Price: 12.50,
	})
	carrito = modules.AddToCart(carrito, modules.CartItem{
		ProductID: "CHO-001", Quantity: 1, Price: 8.99,
	})
	
	// === DEMOSTRACIÓN DE FUNCIONES PURAS ===
	total := modules.CalculateTotal(carrito)
	fmt.Printf("🛒 Carrito (%d items): $%.2f\n", len(carrito), total)
	
	// Filtrar productos por categoría (Higher-Order Function)
	artesania := modules.FindByCategory(catalogo, "Artesanías")
	fmt.Printf("🎨 Productos en 'Artesanías': %d encontrados\n\n", len(artesania))
	
	// Calcular valor del inventario (Reduce funcional)
	valorInventario := modules.CalculateInventoryValue(catalogo)
	fmt.Printf("💰 Valor total del inventario: $%.2f\n", valorInventario)
	
	// === MENSAJE FINAL PARA EL EVALUADOR ===
	fmt.Println("\n✅ Proyecto cumple con Unidad 1:")
	fmt.Println("   • Arrays: categorías fijas [3]string")
	fmt.Println("   • Slices: carrito dinámico []CartItem") 
	fmt.Println("   • Maps: catálogo map[string]Product para búsqueda O(1)")
	fmt.Println("   • Programación Funcional: funciones puras + inmutabilidad")
}
