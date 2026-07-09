package cli

import (
    "fmt"

    "seindonesia/internal/service"
    "seindonesia/shared/utils"
)

func CategorySection() {
    utils.Clear()
    fmt.Print("Category Sei Indonesia\n\n")
    
    categories := service.GetAllCategories()
    
    for idx, cat := range categories {
        fmt.Printf("%d. %s\n", idx+1, cat.DisplayName)
    }

    fmt.Print("\nChoose Menu: ")
    var question int
    fmt.Scan(&question)

    if question > 0 && question <= len(categories) {
        ProductSection(categories[question-1].ID)
    }
}