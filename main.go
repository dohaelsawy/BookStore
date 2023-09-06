package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/dohaelsawy/bookStore/handlers"
	"github.com/gorilla/mux"
)

func main() {

	logger := log.New(os.Stdout, "BookStore", log.LstdFlags)
	// 1-> init handler of product
	productHandler := handlers.NewProduct(logger)

	// 2-> init goriall mux
	serveMux := mux.NewRouter()

	// get Router
	getRouter := serveMux.Methods("GET").Subrouter()
	getRouter.HandleFunc("/product", productHandler.GetProducts)
	getRouter.HandleFunc("/product/{id:[0-9]+}",productHandler.GetProduct)

	// POST Router
	postRouter := serveMux.Methods("POST").Subrouter()
	postRouter.HandleFunc("/product", productHandler.AddProducts)
	postRouter.Use(productHandler.MiddlewareProductValidation)

	//PUT Router
	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/product/{id:[0-9]+}", productHandler.UpdateProduct)
	putRouter.Use(productHandler.MiddlewareProductValidation)

	//DELETE Router
	deleteRouter := serveMux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/product/{id:[0-9]+}", productHandler.DeleteProduct)

	// 3-> init the server with specified settings
	s := http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	// 4 -> start the server
	go func() {
		logger.Println("the server is running on port 9090")
		err := s.ListenAndServe()
		if err != nil {
			logger.Println("panic there is a problem !!!!")
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}
