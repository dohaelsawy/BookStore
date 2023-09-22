package handlers

import (
	"fmt"
	"log"
	"net/http"
	"context"
	"github.com/dohaelsawy/bookStore/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}


func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod:= &data.Product{}
		err := prod.FromJson(r.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		err = prod.Validate()
		if err != nil {
			http.Error(rw, fmt.Sprint("you have validate the porduct fields %v" , err), http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(),keyProduct{},prod)
		newReq := r.WithContext(ctx)
		next.ServeHTTP(rw , newReq)
	})


}
