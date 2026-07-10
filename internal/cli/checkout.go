package cli

import (
    "fmt"

    "seindonesia/internal/service"
    "seindonesia/shared/utils"
)

func CheckoutSection() {
    utils.Clear()
    fmt.Print("Checkout\n\n")
    
    orders := service.GetOrders()
    
    if len(orders) == 0 {
        fmt.Print("Keranjang Checkout Kosong\n")
        fmt.Print("Enter untuk kembali ke menu utama... ")
        fmt.Scanln()
        return
    }

    var grandTotal int
    for i, order := range orders {
        fmt.Printf("%d. %s\n", i+1, order.Product.Name)
        fmt.Printf("   Harga    : %s\n", utils.FormatIDR(order.Product.Price))
        fmt.Printf("   Quantity : %d\n", order.Quantity)
        fmt.Printf("   Total    : %s\n", utils.FormatIDR(order.Total))
        grandTotal += order.Total
        fmt.Println()
    }

    fmt.Printf("Grand Total: %s\n", utils.FormatIDR(grandTotal))
    fmt.Println("\n1. Back to main menu")
    fmt.Println("2. Continue Payment")
    
    var choice string
    fmt.Print("\nChoose option: ")
    fmt.Scan(&choice)
    
    switch choice {
    case "2":
        service.ClearOrders()
        utils.Clear()
        fmt.Print("Pembayaran Berhasil\n")
        fmt.Print("Enter untuk kembali ke menu utama... ")
        fmt.Scanln()
    }
}