package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"seindonesia/internal/domain"
	"seindonesia/shared/utils"
	"strconv"
)

type OrderItem struct { 
  Product domain.Product
  Quantity int
  Total int
}

var orderList []OrderItem

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
  productList(category[question - 1].ID)
}

func productList(Id int) { 
  product := getProductByCategory(Id)
  utils.Clear()
  for { 
    fmt.Printf("Sei Indonesia Menu %s\n\n", product[0].Category.DisplayName)
    if len(orderList) > 0 { 
      fmt.Printf("Berhasil Menambahkan %s\n", orderList[len(orderList) - 1].Product.Name)
    }
    for i ,prod := range product { 
      fmt.Printf("%d. %s %d\n", i + 1, prod.Name, prod.Price)
    }

    fmt.Print("\nChoose Menu: ")
    var question string 
    var quantity int

    fmt.Scan(&question)
    num, err := strconv.Atoi(question)
    if err != nil {
        fmt.Println("Input tidak valid")
        continue
    }

    if num == 0 {
        main()
        return
    }
    
    if num > len(product) || num < 1 {
        continue
    }
    
    selected := product[num-1]
    fmt.Printf("\nQuantity %s: ", selected.Name)
    fmt.Scan(&quantity)
    
    fmt.Print(product[num - 1])
    orderList = append(orderList, OrderItem{
        Product:  selected,
        Quantity: quantity,
        Total:    selected.Price * quantity,
    })
    fmt.Print(orderList)
    utils.Clear()
  }
}

func CheckoutSection() { 
  fmt.Print("Checkout\n")
  for i, order := range orderList { 
    fmt.Printf("%d. %s %d\n", i + 1, order.Product.Name, order.Product.Price)
  }
}

func MainMenu() { 
  utils.Clear()
  fmt.Print("Sei Indonesia\n\n")
  fmt.Print("1. Check Menu\n2. Checkout")

  fmt.Print("\n\nChoose Menu: ")  
  var question string 
  fmt.Scan(&question)

  switch question { 
    case "0": 
    os.Exit(0)
    case "1": 
    CategorySection()
    case "2":
    CheckoutSection() 
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