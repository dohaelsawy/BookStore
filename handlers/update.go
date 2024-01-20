package handlers

import (
	"net/http"
	"strconv"
	"github.com/dohaelsawy/bookStore/data"
	"github.com/gorilla/mux"
	
)


func (product *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, e := strconv.Atoi(vars["id"])
	if e != nil {
		http.Error(rw, "can't update the product .. panic !!!", http.StatusInternalServerError)
		return
	}
	db , err := data.DbCreateObject()
	if err != nil {
		http.Error(rw , "can't open db connection .. panic !!!" , http.StatusInternalServerError)
	}
	defer db.Close()
	product.l.Println("watch out i am about to change product object")
	upP := r.Context().Value(keyProduct{}).(*data.Product)
	err = data.UpdateProduct(db,id,upP)
	if err != nil{
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	product.l.Println("product updated :}")
}
