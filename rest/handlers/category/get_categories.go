package category

import (
	"net/http"

	"github.com/gowalillah/ecommerce/util"
)

func (h *Handler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.svc.List()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.SendData(w, http.StatusOK, categories)
}
