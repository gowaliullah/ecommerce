package products

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gowalillah/ecommerce/database"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid JSON", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(database.ProductList)
}
