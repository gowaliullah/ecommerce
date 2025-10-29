package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gowalillah/ecommerce/domain"
	"github.com/gowalillah/ecommerce/util"
)

type ReqCreatedProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	ImgUrl      string  `json:"img_url"`
	CreatedBy   string  `json:"created_by"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var req ReqCreatedProduct

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usr, ok := util.GetUserFromContext(r, h.cnf)
	fmt.Println(usr, ok)

	if !ok {
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if usr.Role != "seller" {
		util.SendError(w, http.StatusForbidden, "Forbidden: seller role required")
		return
	}

	fmt.Println("All its okay")

	createdProduct, err := h.svc.Create(domain.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		ImgUrl:      req.ImgUrl,
		CreatedBy:   usr.Uuid,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusCreated, createdProduct)
}
