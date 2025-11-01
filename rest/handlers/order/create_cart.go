package order

import (
	"encoding/json"
	"net/http"

	"github.com/gowalillah/ecommerce/domain"
	"github.com/gowalillah/ecommerce/util"
)

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req domain.Order

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	usr, ok := util.GetUserFromContext(r, h.cnf)

	if !ok {
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	createdOrder, err := h.svc.Create(domain.Order{
		UserID:    usr.ID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Total:     req.Total,
		Status:    req.Status,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusCreated, createdOrder)
}
