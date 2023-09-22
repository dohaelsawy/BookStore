package handlers

import (
	"net/http"
	"strconv"
	"github.com/dohaelsawy/bookStore/data"
	"github.com/gorilla/mux"
)


func (product *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	product.l.Println("fetching some data watch out !!!")
	db , err := data.DbCreateObject()
	if err != nil {
		http.Error(rw, "products are not found , our getting method is's working .. panic", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	prodlist , err := data.GetProducts(db)
	if err != nil {
		http.Error(rw, "can't get list of productss .. panic !!!", http.StatusInternalServerError)
		return
	}
	err = prodlist.ToJson(rw)
	if err != nil {
		http.Error(rw, "can't write it in rw .. panic !!!", http.StatusInternalServerError)
		return
	}
	product.l.Println("productsss fetched :}")
}

func (product *Products) GetProduct (rw http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	id , err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "can't convert product id to an string panic !!!!!!", http.StatusInternalServerError)
		return
	}
	db , err := data.DbCreateObject()
	if err != nil {
		http.Error(rw, "can't open db object ... panic !!!", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	product.l.Println("fetching ONE product !!!!!")
	p , err := data.GetOneProduct(db,id) 
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	err = p.ToJson(rw)
	if err != nil {
		http.Error(rw, "can't write it in rw .. panic !!!", http.StatusInternalServerError)
		return
	}
	product.l.Println("product fetched :}")
}

