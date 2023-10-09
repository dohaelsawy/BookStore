package handlers

import (
	"net/http"
	"strconv"

	"github.com/dohaelsawy/bookStore/data/product"
	"github.com/gorilla/mux"
)

func (product *Products) DeleteProduct (rw http.ResponseWriter , r * http.Request){
	vars := mux.Vars(r)
	id , err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw , "con't convert id to integer ... panic !!!" , http.StatusBadRequest)
		return
	}
	db , err := data.DbCreateObject()
	if err != nil {
		http.Error(rw, "can't open db object ... panic !!!", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	product.l.Println("watch out deleting one product !!!!")
	err = data.DeleteProduct(db,id)
	if err != nil {
		http.Error(rw,"can't delete the product ... panic !!!", http.StatusBadRequest)
		return
	}
	product.l.Println("product deleted :}")

}