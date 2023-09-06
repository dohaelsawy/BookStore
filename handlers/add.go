package handlers

import (
	"net/http"
	"github.com/dohaelsawy/bookStore/data"
)

type keyProduct struct {}

func (product *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	product.l.Println("handle adding new product")
	newP := r.Context().Value(keyProduct{}).(*data.Product)
	//product.l.Printf("profuct is : %#v" , newP)
	data.AddProduct(newP)
}

