package cli

import (
    "fmt"
    "os"

    "seindonesia/shared/utils"
)

func MainMenu() {
    for {
        utils.Clear()

        fmt.Println("Sei Indonesia")
        fmt.Println()
        fmt.Println("1. Check Menu")
        fmt.Println("2. Checkout")
        fmt.Println("0. Exit")

        var input string
        fmt.Print("\nChoose Menu: ")
        fmt.Scan(&input)

        switch input {
        case "1":
            CategorySection()
        case "2":
            CheckoutSection()
        case "0":
            os.Exit(0)
        }
    }
}