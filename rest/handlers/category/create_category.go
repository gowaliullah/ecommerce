package category

import (
	"encoding/json"
	"net/http"

	"github.com/gowalillah/ecommerce/domain"
	"github.com/gowalillah/ecommerce/util"
)

type ReqCreatedCategory struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var data ReqCreatedCategory

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// usr, ok := util.GetUserFromContext(r, h.cnf)

	// fmt.Println(usr)

	// if !ok {
	// 	util.SendError(w, http.StatusUnauthorized, "Unauthorized")
	// 	return
	// }

	// if usr.Role != "seller" && usr.Role != "admin" {
	// 	util.SendError(w, http.StatusForbidden, "Forbidden: seller role required")
	// 	return
	// }

	createdCategory, err := h.svc.Create(domain.Category{
		Name: data.Name,
		Slug: data.Slug,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusOK, createdCategory)
}
