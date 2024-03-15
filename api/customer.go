package api

import (
	"database/sql"
	"net/http"

	db "github.com/dohaelsawy/bookStore/db/sqlc"
	"github.com/dohaelsawy/bookStore/util"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


type customerData struct {
	FirstName   string      `json:"first_name" validate:"required,min=1,max=20"`
	LastName    string      `json:"last_name" validate:"required,min=1,max=20"`
	Email       string      `json:"email" validate:"required,email"`
	Password    string      `json:"password" validate:"required,min=8"`
	City        string      `json:"city" validate:"required,min=1,max=10"`
	PhoneNumber string      `json:"phone_number" validate:"required,min=11,max=11"`
}


type loginCustomerRequest struct {
	Email       string      `json:"email" validate:"required,email"`
	Password    string      `json:"password" validate:"required,min=8"`
}

type responseLoginCustomer struct {
	Access_token string     `json:"access_token"`
	Customer customerData   `json:"customer"`
}


func (c *customerData ) Validate() error {
	validate := validator.New()
	if err := validate.Struct(c) ; err != nil {
		return err	
	}
	return nil 
}

func (c *loginCustomerRequest ) Validate() error {
	validate := validator.New()
	if err := validate.Struct(c) ; err != nil {
		return err	
	}
	return nil 
}

func (server *Server) Login (ctx *gin.Context) {
	server.l.Println("i'm in login function yaaahhhh ... ")

	var req loginCustomerRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest , "can't bind input json with request")
		return
	}
	if err = req.Validate() ; err != nil {
		ctx.JSON(http.StatusBadRequest , errorResponce(err))
		return
	}

	customer , err := server.store.GetCustomerByEmail(ctx,req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound , errorResponce(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError , errorResponce(err))
		return
	}

	err = util.CheckPassword(req.Password , customer.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized , "the password isn't correct")
		return
	}

	accessToken , err := server.tokenMaker.CreateToken(req.Email , server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError , "can't create token")
		return
	}

	customerResponse := responseLoginCustomer {
		Access_token: accessToken,
		Customer: customerData{
			FirstName: customer.FirstName,
			LastName: customer.LastName,
			Email: customer.Email,
			Password: customer.Password,
			City: customer.City,
			PhoneNumber: customer.PhoneNumber,
		},
	}
	server.l.Println(customer)
	ctx.JSON(http.StatusOK , customerResponse)
}



func (server * Server) SignUp (ctx *gin.Context) {
	server.l.Println("i'm in signup function yaaahhhh ... ")

	var req customerData 
	err := ctx.ShouldBindJSON(&req) 
	if err != nil {
		ctx.JSON(http.StatusBadRequest , "can't bind input json with request")
		return 
	}

	if err = req.Validate() ; err != nil {
		ctx.JSON(http.StatusBadRequest , errorResponce(err))
		return 
	}

	_, err = server.store.GetCustomerByEmail(ctx , req.Email)
	if err == nil {
		ctx.JSON(http.StatusConflict , "the user already exists!!! ")
		return 
	}

	hashedPassword , err :=  util.HashPassword(req.Password) 
	if err != nil {
		ctx.JSON(http.StatusInternalServerError , err)
		return 
	}
	

	arg := db.CreatecustomerParams {
		FirstName: req.FirstName,
		LastName: req.LastName,
		Email: req.Email,
		Password: hashedPassword ,   
		City: req.City,      
		PhoneNumber: req.PhoneNumber,
	}

	customer , err := server.store.Createcustomer(ctx , arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,err)
		return
	}

	ctx.JSON(http.StatusOK , customer)
}