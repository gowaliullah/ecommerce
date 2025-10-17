package cart

import (
	"encoding/json"
	"net/http"

	"github.com/gowalillah/ecommerce/domain"
	"github.com/gowalillah/ecommerce/util"
)

func (h *Handler) CreateCart(w http.ResponseWriter, r *http.Request) {
	var cart domain.Cart

	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		http.Error(w, "Provide valid JSON data", http.StatusBadRequest)
		return
	}

	createdCart, err := h.svc.Create(domain.Cart{
		UserID: cart.ID,
	})

	if err != nil {
		http.Error(w, "Failed to added cart", http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusOK, createdCart)

}
