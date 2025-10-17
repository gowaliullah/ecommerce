package cart

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gowalillah/ecommerce/domain"
	"github.com/gowalillah/ecommerce/util"
)

func (h *Handler) UpdateCart(w http.ResponseWriter, r *http.Request) {
	catId := r.PathValue("id")
	id, err := strconv.Atoi(catId)
	if err != nil {
		http.Error(w, "Please give the valid cart id", http.StatusBadRequest)
		return
	}

	var req domain.Category

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Provide valid JSON data", http.StatusBadRequest)
		return
	}

	updatedCart, err := h.svc.Update(domain.Cart{
		ID:     id,
		UserID: req.ID,
	})

	if err != nil {
		http.Error(w, "Failed to update cart", http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusOK, updatedCart)

}
