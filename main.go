//Package classification Product API
//
//Documenting for product api
//
//	Schemes: http
//	Host: localhost
//	BasePath: /product
//
//	Consumes:
//	- application/json
//
//	Produces:
//	-application/json
//swagger:meta

package main

import (
	"context"
	"log"
	"os"

	"github.com/dohaelsawy/bookStore/api"
	db "github.com/dohaelsawy/bookStore/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

const (
	dbDriver     = "postgres"
	DBSource   = "postgresql://root:root@localhost:5432/postgres?sslmode=disable"
	serverAdress = "0.0.0.0:8080"
)

func main() {
	logger := log.New(os.Stdout, "BookStore", log.LstdFlags)
	connPool, err := pgxpool.New(context.Background(), DBSource)
	if err != nil {
		log.Fatal("the error is in db connection ... panic ", err)
	}
	store := db.NewStore(connPool)
	server := api.NewServer(store, logger)

	err = server.Start(serverAdress)
	if err != nil {
		log.Fatal("the error is cannot start server... panic ", err)
	}
}
