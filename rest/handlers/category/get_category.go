package category

import (
	"net/http"
	"strconv"

	"github.com/gowalillah/ecommerce/util"
)

func (h *Handler) GetCategory(w http.ResponseWriter, r *http.Request) {
	catId := r.PathValue("id")
	id, err := strconv.Atoi(catId)
	if err != nil {
		http.Error(w, "Please give the valid product id", http.StatusBadRequest)
		return
	}

	category, _ := h.svc.Get(id)
	if category == nil {
		util.SendError(w, http.StatusNotFound, "Category not found")
	}

	util.SendData(w, http.StatusOK, category)
}
