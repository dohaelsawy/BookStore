package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"github.com/dohaelsawy/bookStore/handlers"
)

func main (){

	logger := log.New(os.Stdout, "BookStore", log.LstdFlags)
	// 1-> init handler of product 
	productHandler := handlers.NewProduct(logger)
	
	// 2-> init serve mux
	serveMux := http.NewServeMux()
	serveMux.Handle("/" , productHandler)

	// 3-> init the server with specified settings 
	s := http.Server {
		Addr: ":9090",
		Handler: serveMux,
		ErrorLog: logger,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 120 * time.Second,
	}
	// 4 -> start the server 
	go func ()  {
		logger.Println("the server is running on port 9090")
		err := s.ListenAndServe()
		if err != nil {
			logger.Println("panic there is a problem !!!!")
			os.Exit(1)
		}
	}()



}