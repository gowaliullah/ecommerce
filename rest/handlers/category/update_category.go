package category

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gowalillah/ecommerce/domain"
	"github.com/gowalillah/ecommerce/util"
)

func (h *Handler) UpdateCategory(w http.ResponseWriter, r *http.Request) {

	catId := r.PathValue("id")
	id, err := strconv.Atoi(catId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req ReqCreatedCategory

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedCategory, err := h.svc.Update(domain.Category{
		ID:   id,
		Name: req.Name,
		Slug: req.Slug,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusOK, updatedCategory)

}
