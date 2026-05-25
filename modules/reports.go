package modules

import "gestion-ecommerce-go/utils"

// ReportSummary: Estructura para reportes básicos (usa slices y maps)
type ReportSummary struct {
    TotalVentas   float64
    TotalUnidades int
    Categorias    []string
}

// GenerateBasicReport: Función pura que genera reporte de ventas
// Aplica Reduce + Maps + Slices (conceptos Unidad 1)
func GenerateBasicReport(cart ShoppingCart, catalog ProductCatalog) ReportSummary {
    // Calcular total de ventas con Reduce (función de orden superior)
    totalVentas := utils.Reduce(cart, 0.0, func(acc float64, item CartItem) float64 {
        return acc + (item.Price * float64(item.Quantity))
    })

    // Contar unidades totales
    totalUnidades := 0
    for _, item := range cart {
        totalUnidades += item.Quantity
    }

    // Extraer categorías únicas presentes en el carrito (Map + Slice)
    categoriaMap := make(map[string]bool)
    for _, item := range cart {
        if prod, ok := catalog[item.ProductID]; ok {
            categoriaMap[prod.Category] = true
        }
    }
    
    var categorias []string
    for cat := range categoriaMap {
        categorias = append(categorias, cat)
    }

    return ReportSummary{
        TotalVentas:   totalVentas,
        TotalUnidades: totalUnidades,
        Categorias:    categorias,
    }
}
