package domain

type Category struct { 
  ID int `json:"id"`
  Name string `json:"name"`
  DisplayName string `json:"display_name"`
}

type Product struct { 
  ID string `json:"id"`
  Name string `json:"name"`
  Description string `json:"description"`
  Category Category `json:"category"`
  Price int `json:"price"`
}