package user

import (
	"net/http"
	"strconv"

	"github.com/gowalillah/ecommerce/util"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	page, _ := strconv.Atoi(params.Get("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(params.Get("limit"))
	if limit == 0 {
		limit = 10
	}

	users, err := h.svc.List(limit, page)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.SendData(w, http.StatusOK, users)
}
