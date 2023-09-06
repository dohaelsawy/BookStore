package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/dohaelsawy/bookStore/data"
	"github.com/gorilla/mux"
)


func (product *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, e := strconv.Atoi(vars["id"])
	if e != nil {
		http.Error(rw, "can't update the product .. panic", http.StatusInternalServerError)
		return
	}
	product.l.Println("watch out i am about to change product object")
	upP := r.Context().Value(keyProduct{}).(*data.Product)
	err := data.UpdateProduct(id,upP)
	var ErrProductNotFound = fmt.Errorf("product not found")
	if err == ErrProductNotFound {
		http.Error(rw, "can't update the product .. panic", http.StatusInternalServerError)
		return
	}
}
