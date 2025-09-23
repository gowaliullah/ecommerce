package handlers

import (
	"net/http"

	"github.com/gowalillah/ecommerce/util"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, http.StatusOK, "Welcome the this era.......!")
}
