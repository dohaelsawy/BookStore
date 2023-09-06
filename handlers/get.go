package handlers

import(
	"net/http"
	"github.com/dohaelsawy/bookStore/data"
)


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

