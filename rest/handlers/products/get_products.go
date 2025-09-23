package products

import (
	"encoding/json"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(productList)
}
