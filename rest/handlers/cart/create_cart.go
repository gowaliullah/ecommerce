package cart

import (
	"fmt"
	"net/http"

	"github.com/gowalillah/ecommerce/domain"
	"github.com/gowalillah/ecommerce/util"
)

func (h *CartHandler) CreateCart(w http.ResponseWriter, r *http.Request) {
	var cart domain.Cart

	usr, _ := util.GetUserFromContext(r, h.cnf)

	if usr != nil {
		cart.UserID = usr.Uuid
	}

	createdCart, err := h.svc.Create(cart)

	if err != nil {
		fmt.Println("handlers:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusOK, createdCart)

}
