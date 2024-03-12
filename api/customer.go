package api

import (
	// "net/http"
	// db "github.com/dohaelsawy/bookStore/db/sqlc"
	//"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
)


type customerRequest struct {
	FirstName   string      `json:"first_name" validate:"required,min=1,max=10"`
	LastName    string      `json:"last_name" validate:"required , min=1,max= 10"`
	Email       string      `json:"email" validate:"required,email"`
	Password    string      `json:"password" validate:"required , min=8"`
	City        string      `json:"city" validate:"required min=1 , max= 10"`
	PhoneNumber string      `json:"phone_number" validate:"required min=10 , max= 10"`
	Token       string `json:"token"`
}

func (server *Server) login (ctx *gin.Context) {
	server.l.Println("i'm in login function yaaahhhh ... ")


}



func (server * Server) signUp (ctx *gin.Context){
	server.l.Println("i'm in signup function yaaahhhh ... ")




}