package modules

import "gestion-ecommerce-go/utils"

type Product struct {
    ID, Name, Category, Description string
    Price                           float64
    Stock                           int
}

type ProductCatalog map[string]Product

func GetAllProducts(catalog ProductCatalog) []Product {
    products := make([]Product, 0, len(catalog))
    for _, p := range catalog {
        products = append(products, p)
    }
    return products
}

func FindByCategory(catalog ProductCatalog, category string) []Product {
    allProducts := GetAllProducts(catalog)
    return utils.FilterBy(allProducts, func(p Product) bool {
        return p.Category == category
    })
}

func CalculateInventoryValue(catalog ProductCatalog) float64 {
    products := GetAllProducts(catalog)
    return utils.Reduce(products, 0.0, func(acc float64, p Product) float64 {
        return acc + (p.Price * float64(p.Stock))
    })
}
