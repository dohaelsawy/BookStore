package routes

import (
	"github.com/gorilla/mux"
	"github.com/dohaelsawy/BookStore/packages/controllers"
)


var RegisterBookRoutes = func (router *mux.Router){
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/",controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookID}",controllers.getBook).Methods("GET")
	router.HandleFunc("/book/{bookID}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookID}",controllers.DeleteBook).Methods("DELETE")
}