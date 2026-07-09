package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"seindonesia/internal/domain"
	"seindonesia/shared/utils"

)

func getProductAll() []domain.Product { 
  productJson, err := os.Open("../shared/data/product.json")
  if err != nil { 
    fmt.Println(err)
  }

  var product[] domain.Product
  byteValue, err := io.ReadAll(productJson)
  err = json.Unmarshal(byteValue, &product)
  return product 
}

func getProductByCategory(id int) []domain.Product { 
  productJson, err := os.Open("../shared/data/product.json")
  if err != nil { 
    fmt.Println(err)
  }
  var product[] domain.Product
  var result[] domain.Product
  byteValue, err := io.ReadAll(productJson)
  err = json.Unmarshal(byteValue, &product)
  for _ ,prd := range product { 
    if prd.Category.ID == id { 
      result = append(result, prd)
    }
  }
  return result
}


func CategorySection() { 
  utils.Clear()
  fmt.Print("Category Sei Indonesia\n\n")
  categoryJson, _ := os.Open("../shared/data/category.json")
  var category [] *domain.Category
  byteValue, err := io.ReadAll(categoryJson)
  err = json.Unmarshal(byteValue, &category)
  if err != nil {
    fmt.Println("Failed")
  }
  for idx, cat := range category { 
    fmt.Printf("%d. %s\n", idx + 1, cat.DisplayName)
  }

  fmt.Print("\nChoose Menu: ")
  var question int 
  fmt.Scan(&question)
  product := getProductByCategory(category[question - 1].ID)
  for i ,prod := range product { 
    fmt.Printf("%d. %s %d\n", i + 1, prod.Name, prod.Price)
  }
}

func MainMenu() { 
  fmt.Print("Sei Indonesia\n\n")
  fmt.Print("1.Check Menu")

  fmt.Print("\n\nChoose Menu: ")  
  var question string 
  fmt.Scan(&question)

  switch question { 
    case "0": 
    os.Exit(0)
    case "1": 
    CategorySection()
  }
}

func main() { 
  MainMenu()

 
  // for _, p := range product { 
  //   fmt.Println(p.ID)
  //   fmt.Println(p.Name)
  //   fmt.Println(p.Description)
  //   fmt.Println(p.Category.DisplayName)
  //   fmt.Println(p.Price)
  // }
}