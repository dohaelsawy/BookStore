package handlers

import (
	"net/http"
	"strconv"

	"github.com/dohaelsawy/bookStore/data"
	"github.com/gorilla/mux"
)

func (product *Products) DeleteProduct (rw http.ResponseWriter , r * http.Request){
	vars := mux.Vars(r)
	id , err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw , "con't convert id to integer ... panic !!!" , http.StatusBadRequest)
		return
	}
	err = data.DeleteProduct(id)
	if err != nil {
		http.Error(rw,"can't delete the product ... panic !!!", http.StatusBadRequest)
		return
	}

}