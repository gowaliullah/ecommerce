package cart

import (
	"net/http"
	"strconv"

	"github.com/gowalillah/ecommerce/types"
	"github.com/gowalillah/ecommerce/util"
)

func (h *CartHandler) DeleteCart(w http.ResponseWriter, r *http.Request) {
	categoryId := r.PathValue("id")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.svc.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusOK, types.DeleteRes{
		Message: "Successfully deleted cart",
	})
}
