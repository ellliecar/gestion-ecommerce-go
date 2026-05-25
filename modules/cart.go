package modules

import "gestion-ecommerce-go/utils"

type CartItem struct {
    ProductID string
    Quantity  int
    Price     float64
}

type ShoppingCart []CartItem

func AddToCart(cart ShoppingCart, item CartItem) ShoppingCart {
    newCart := make(ShoppingCart, len(cart)+1)
    copy(newCart, cart)
    newCart[len(cart)] = item
    return newCart
}

func CalculateTotal(cart ShoppingCart) float64 {
    return utils.Reduce(cart, 0.0, func(total float64, item CartItem) float64 {
        return total + (item.Price * float64(item.Quantity))
    })
}
