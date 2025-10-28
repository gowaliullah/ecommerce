package product

import (
	"net/http"
	"strconv"

	"github.com/gowalillah/ecommerce/util"
)

type sendRes struct {
	Message string `json:"message"`
}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give the valid product id", 400)
		return
	}

	err = h.svc.Delete(id)
	if err != nil {
		http.Error(w, "fail to deleted product", http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusOK, sendRes{
		Message: "Product successfully deleted......",
	})

}
