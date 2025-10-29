package product

import (
	"net/http"
	"strconv"

	"github.com/gowalillah/ecommerce/types"
	"github.com/gowalillah/ecommerce/util"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give the valid product id", 400)
		return
	}

	usr, ok := util.GetUserFromContext(r, h.cnf)

	if !ok {
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if usr.Role != "seller" {
		util.SendError(w, http.StatusForbidden, "Forbidden: seller role required")
		return
	}

	err = h.svc.Delete(id)
	if err != nil {
		http.Error(w, "fail to deleted product", http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusOK, types.DeleteRes{
		Message: "Product successfully deleted......",
	})

}
