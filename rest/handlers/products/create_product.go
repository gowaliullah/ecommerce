package products

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid JSON", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(productList)
}
