package handlers

import (
	"net/http"
	"strconv"

	"github.com/dohaelsawy/bookStore/data"
	"github.com/gorilla/mux"
)


func (product *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	product.l.Println("fetching some data watch out !!!")
	ourProducts := data.GetProducts()
	err := ourProducts.ToJson(rw)
	if err != nil {
		http.Error(rw, "products are not found , our getting method is's working .. panic", http.StatusInternalServerError)
		return
	}
}

func (product *Products) GetProduct (rw http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	id , err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "can't convert product id to an string panic !!!!!!", http.StatusInternalServerError)
		return
	}

	product.l.Println("fetching ONE product !!!!!")
	p , err := data.GetOneProduct(id) 
	if err != nil {
		http.Error(rw, "products are not found , our getting method isn't working .. panic", http.StatusInternalServerError)
		return
	}
	err = p.ToJson(rw)
	if err != nil {
		http.Error(rw, "can't convert product to json .. panic", http.StatusInternalServerError)
		return
	}
}

