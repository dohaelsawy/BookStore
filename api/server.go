package api

import (
	"log"

	db "github.com/dohaelsawy/bookStore/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	store  *db.Store
	l      *log.Logger
}

func NewServer(store *db.Store , l *log.Logger) *Server{
	server := Server{store: store , l: l}
	router := gin.Default()

	router.POST("/addBook", server.addBook)
	router.PUT("/udateBook/:id", server.updateBook)
	router.DELETE("/deleteBook/:id", server.deleteBook)
	router.GET("/getBook/:id", server.getBook)
	router.GET("/listBooks",server.listBooks)
	server.router = router
	return &server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponce(err error) gin.H {
	return gin.H{"error": err.Error()}
}
