package product

import (
	"encoding/json"
	"net/http"

	"github.com/gowalillah/ecommerce/database"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.ProductList)
}
