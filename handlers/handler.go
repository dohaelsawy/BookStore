package handlers


import (
	"log"
	"net/http"
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
