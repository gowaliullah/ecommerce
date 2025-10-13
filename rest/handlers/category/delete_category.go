package category

import (
	"net/http"
	"strconv"

	"github.com/gowalillah/ecommerce/util"
)

func (h *Handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := r.PathValue("id")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		http.Error(w, "Please give the valid category id", 400)
		return
	}

	err = h.svc.Delete(id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	util.SendData(w, http.StatusOK, "Successfully deleted product")
}
