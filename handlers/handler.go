package handlers


import (
	"log"
	"net/http"

)

type Products struct {
	l *log.Logger 
}

func NewProduct (l *log.Logger) *Products {
	return &Products{l}
}

func (product *Products) ServeHTTP(rw http.ResponseWriter , r *http.Request){
	


}