package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gowalillah/ecommerce/domain"
	"github.com/gowalillah/ecommerce/util"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	prdId := r.PathValue("id")
	id, err := strconv.Atoi(prdId)
	if err != nil {
		http.Error(w, "Please give the valid product id", http.StatusBadRequest)
		return
	}

	var req ReqCreatedProduct

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Provide valid JSON data", http.StatusBadRequest)
		return
	}

	updatedProduct, err := h.svc.Update(domain.Product{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		ImgUrl:      req.ImgUrl,
	})

	if err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusOK, updatedProduct)

}
