package cli

import (
    "fmt"
    "strconv"

    "seindonesia/internal/domain"
    "seindonesia/internal/service"
    "seindonesia/shared/utils"
)

func ProductSection(categoryID int) {
    products := service.GetProductByCategory(categoryID)
    
    if len(products) == 0 {
        fmt.Println("No products found in this category")
        return
    }

    for {
        utils.Clear()
        fmt.Printf("Sei Indonesia Menu %s\n\n", products[0].Category.DisplayName)
        
        orders := service.GetOrders()
        if len(orders) > 0 {
            fmt.Printf("Berhasil Menambahkan %s\n", orders[len(orders)-1].Product.Name)
        }
        
        for i, prod := range products {
            fmt.Printf("%d. %s %s\n", i+1, prod.Name, utils.FormatIDR(prod.Price))
        }
        fmt.Println("0. Back to main menu")

        fmt.Print("\nChoose Menu: ")
        var question string
        fmt.Scan(&question)

        num, err := strconv.Atoi(question)
        if err != nil {
            fmt.Println("Input tidak valid")
            continue
        }

        if num == 0 {
            return
        }

        if num < 1 || num > len(products) {
            continue
        }

        selected := products[num-1]
        fmt.Printf("\nQuantity %s: ", selected.Name)
        var quantity int
        fmt.Scan(&quantity)

        if quantity <= 0 {
            fmt.Println("Quantity harus lebih dari 0")
            continue
        }

        order := domain.OrderItem{
            Product:  selected,
            Quantity: quantity,
            Total:    selected.Price * quantity,
        }
        
        service.AddOrder(order)
        utils.Clear()
    }
}