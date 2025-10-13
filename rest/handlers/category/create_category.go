package category

import (
	"encoding/json"
	"net/http"

	"github.com/gowalillah/ecommerce/domain"
	"github.com/gowalillah/ecommerce/util"
)

func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var data domain.Category

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Provide valid JSON data", http.StatusBadRequest)
		return
	}

	createdCategory, err := h.svc.Create(domain.Category{
		Name:     data.Name,
		ImageUrl: data.ImageUrl,
	})
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusOK, createdCategory)
}
