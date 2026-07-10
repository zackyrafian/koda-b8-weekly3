package service

import (
    "encoding/json"
    "io"
    "os"
    "path/filepath"

    "seindonesia/internal/domain"
)

func getProductFilePath() string {
    return filepath.Join("..", "shared", "data", "product.json")
}

func getCategoryFilePath() string {
    return filepath.Join("..", "shared", "data", "category.json")
}

func GetAllCategories() []domain.Category {
    categoryFile, err := os.Open(getCategoryFilePath())
    if err != nil {
        return []domain.Category{}
    }
    defer categoryFile.Close()

    var categories []domain.Category
    byteValue, err := io.ReadAll(categoryFile)
    if err != nil {
        return []domain.Category{}
    }

    err = json.Unmarshal(byteValue, &categories)
    if err != nil {
        return []domain.Category{}
    }

    return categories
}

func GetAllProducts() []domain.Product {
    productFile, err := os.Open(getProductFilePath())
    if err != nil {
        return []domain.Product{}
    }
    defer productFile.Close()

    var products []domain.Product
    byteValue, err := io.ReadAll(productFile)
    if err != nil {
        return []domain.Product{}
    }

    err = json.Unmarshal(byteValue, &products)
    if err != nil {
        return []domain.Product{}
    }

    return products
}

func GetProductByCategory(categoryID int) []domain.Product {
    allProducts := GetAllProducts()
    var result []domain.Product

    for _, product := range allProducts {
        if product.Category.ID == categoryID {
            result = append(result, product)
        }
    }

    return result
}