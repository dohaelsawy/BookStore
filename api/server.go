package api

import (
	"fmt"
	"log"

	"github.com/dohaelsawy/bookStore/token"
	"github.com/dohaelsawy/bookStore/util"

	db "github.com/dohaelsawy/bookStore/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	store  *db.Store
	l      *log.Logger
	tokenMaker token.Maker
	config util.Config
}

func NewServer(store *db.Store , l *log.Logger , config util.Config) (*Server , error){
	tokenmaker , err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil , fmt.Errorf("can't create token maker %v" , err)
	}
	server := Server{store: store , l: l , tokenMaker : tokenmaker , config: config}
	router := gin.Default()

	router.POST("/user/SignUp",server.SignUp)
	router.POST("/user/Login",server.Login)

	authRouters := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRouters.POST("/addBook", server.addBook)
	authRouters.PUT("/udateBook/:id", server.updateBook)
	authRouters.DELETE("/deleteBook/:id", server.deleteBook)
	authRouters.GET("/getBook/:id", server.getBook)
	authRouters.GET("/listBooks",server.listBooks)
	
	server.router = router
	return &server , nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponce(err error) gin.H {
	return gin.H{"error": err.Error()}
}
