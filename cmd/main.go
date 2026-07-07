package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"seindonesia/internal/domain"
)


func main() { 

  productJson, err := os.Open("../shared/data/product.json")
  if err != nil { 
    fmt.Println(err)
  }

  var product[] domain.Product
  byteValue, err := io.ReadAll(productJson)
  err = json.Unmarshal(byteValue, &product)

  for _, p := range product { 
    fmt.Println(p.ID)
    fmt.Println(p.Name)
    fmt.Println(p.Description)
    fmt.Println(p.Category.DisplayName)
    fmt.Println(p.Price)
  }
}