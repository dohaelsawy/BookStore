package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dohaelsawy/bookStore/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (product *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, e := strconv.Atoi(vars["id"])
	if e != nil {
		http.Error(rw, "can't update the product .. panic", http.StatusInternalServerError)
		return
	}
	product.l.Println("watch out i am about to change product object")
	upP := &data.Product{}

	err := upP.FromJson(r.Body)
	//product.l.Printf("prod : %#v" , upP)

	data.UpdateProduct(id, upP)
	var ErrProductNotFound = fmt.Errorf("product not found")
	if err == ErrProductNotFound {
		http.Error(rw, "can't update the product .. panic", http.StatusInternalServerError)
		return
	}
}

func (product *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	product.l.Println("handle adding new product")

	newP := &data.Product{}

	err := newP.FromJson(r.Body)

	if err != nil {
		http.Error(rw, "can't add the product .. panic", http.StatusInternalServerError)
	}
	//product.l.Printf("profuct is : %#v" , newP)
	data.AddProduct(newP)
}

func (product *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	//1 -> add transactions to logger file
	product.l.Println("fetching some data watch out !!!")
	//2 -> fetch data from file / database whatever
	ourProduct := data.GetProducts()
	//3 -> covert this list to json format
	err := ourProduct.ToJson(rw)
	if err != nil {
		http.Error(rw, "products are not found , our getting method is's working .. panic", http.StatusInternalServerError)
	}
}
