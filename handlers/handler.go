package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"fmt"
	"github.com/dohaelsawy/bookStore/data"
)

type Products struct {
	l *log.Logger 
}

func NewProduct (l *log.Logger) *Products {
	return &Products{l}
}

func (product *Products) ServeHTTP(rw http.ResponseWriter , r *http.Request){
	if r.Method == http.MethodGet {
		product.getProducts(rw ,r)
		return
	}
	if r.Method == http.MethodPost{
		product.addProducts(rw ,r)
		return
	}

	if r.Method == http.MethodPut{
		path := r.URL.Path
		reg := regexp.MustCompile(`/([0-9]+)`) // "/" here refer to the root 
		group := reg.FindAllStringSubmatch(path,-1)
		//product.l.Printf("%v",len(group[0][1]))
		if len(group) > 1 {
			http.Error(rw , "it receives more than a group .. panic !!!" , http.StatusBadRequest)
			return
		}
		if len(group[0]) != 2 {
			http.Error(rw , "it receives more than a group .. panic !!!" , http.StatusBadRequest)
			return
		}

		idString := group[0][1] 
		id , err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw , "can't convert id string to int .. panic !!!" , http.StatusBadRequest)
			return
		}
		product.updateProduct(rw , r , id)
		return

	}
}

func (product *Products) updateProduct(rw http.ResponseWriter , r * http.Request , id int){
	product.l.Println("watch out i am about to change product object")
	upP := &data.Product{}

	err := upP.FromJson(r.Body)
	//product.l.Printf("prod : %#v" , upP)

	data.UpdateProduct(id , upP)
	var ErrProductNotFound = fmt.Errorf("product not found")
	if err == ErrProductNotFound {
		http.Error(rw , "can't update the product .. panic" , http.StatusInternalServerError)
		return
	}
}

func (product *Products) addProducts(rw http.ResponseWriter ,r *http.Request){
	product.l.Println("handle adding new product")

	newP := &data.Product{}

	err := newP.FromJson(r.Body)

	if err != nil {
		http.Error(rw , "can't add the product .. panic" , http.StatusInternalServerError)
	}
	//product.l.Printf("profuct is : %#v" , newP)
	data.AddProduct(newP)
}

func (product *Products) getProducts (rw http.ResponseWriter , r *http.Request){
	//1 -> add transactions to logger file
	product.l.Println("fetching some data watch out !!!")
	//2 -> fetch data from file / database whatever 
	ourProduct := data.GetProducts()
	//3 -> covert this list to json format 
	err := ourProduct.ToJson(rw)
	if err != nil {
		http.Error(rw , "products are not found , our getting method is's working .. panic" , http.StatusInternalServerError)
	}
}
