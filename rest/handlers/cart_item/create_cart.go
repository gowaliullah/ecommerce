package cartitem

import (
	"encoding/json"
	"net/http"

	"github.com/gowalillah/ecommerce/domain"
	"github.com/gowalillah/ecommerce/util"
)

func (h *CartItemHandler) CreateCartItem(w http.ResponseWriter, r *http.Request) {
	var cart domain.Cart

	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usr, _ := util.GetUserFromContext(r, h.cnf)

	createdCart, err := h.svc.Create(domain.Cart{
		UserID: usr.Uuid,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusOK, createdCart)

}
