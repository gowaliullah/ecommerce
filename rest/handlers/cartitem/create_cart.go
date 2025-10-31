package cartitem

import (
	"encoding/json"
	"net/http"

	"github.com/gowalillah/ecommerce/domain"
	"github.com/gowalillah/ecommerce/util"
)

func (h *CartItemHandler) CreateCartItem(w http.ResponseWriter, r *http.Request) {
	var cartItem domain.CartItem

	err := json.NewDecoder(r.Body).Decode(&cartItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdCart, err := h.svc.Create(domain.CartItem{
		CartID:    cartItem.CartID,
		ProductID: cartItem.ProductID,
		Quantity:  cartItem.Quantity,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusOK, createdCart)
}
