package registerLogin

import (
	"log"
	"net/http"
	"fmt"
	"context"
	data "github.com/dohaelsawy/bookStore/data/person"
	//import "github.com/golang-jwt/jwt/v5"
)

type People struct {
	l *log.Logger 
}

func NewPerson(l *log.Logger) *People {
	return &People{l}
}


func (person *People)MiddlewarePersonValidation (next http.Handler)  http.Handler{
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			human := &data.Person{}
			err := human.FromJson(r.Body)
			if err != nil {
				http.Error(rw ,fmt.Sprintln("error in middleware regestration:" , err) , http.StatusInternalServerError)
				return
			}
			err = human.Validate() 
			if err != nil {
				http.Error(rw ,fmt.Sprintln("you have validate the porduct fields %v register" , err) , http.StatusInternalServerError)
				return
			}
			ctx := context.WithValue(r.Context(),keyPerson{},human)
			newReqest := r.WithContext(ctx)
			next.ServeHTTP(rw , newReqest)
		})

}