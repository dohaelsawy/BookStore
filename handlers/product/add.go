package handlers

import (
	"net/http"
	"github.com/dohaelsawy/bookStore/data/product"
)

type keyProduct struct {}

func (product *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	product.l.Println("handle adding new product")
	newP := r.Context().Value(keyProduct{}).(*data.Product)
	//product.l.Printf("profuct is : %#v" , newP)
	db , err := data.DbCreateObject()
	if err != nil {
		http.Error(rw, "can't open db object ... panic !!!", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	err = data.InsertProduct(db , newP)
	if err != nil {
		http.Error(rw , err.Error() , http.StatusBadRequest)
		return
	}
	product.l.Println("product created :}")
}

