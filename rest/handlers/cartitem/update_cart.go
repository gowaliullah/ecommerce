package cartitem

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gowalillah/ecommerce/domain"
	"github.com/gowalillah/ecommerce/util"
)

func (h *CartItemHandler) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	catId := r.PathValue("id")
	id, err := strconv.Atoi(catId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req domain.CartItem

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedCart, err := h.svc.Update(domain.CartItem{
		ID:        id,
		Uuid:      req.Uuid,
		CartID:    req.Uuid,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusOK, updatedCart)
}
