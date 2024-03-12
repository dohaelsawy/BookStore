package main
import (
	"context"
	"log"
	"os"

	"github.com/dohaelsawy/bookStore/api"
	db "github.com/dohaelsawy/bookStore/db/sqlc"
	"github.com/dohaelsawy/bookStore/util"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func main() {
	logger := log.New(os.Stdout, "BookStore", log.LstdFlags)
	config , err := util.LoadAppEnv(".")
	if err != nil {
		log.Fatal("can't load env variables ... panic " , err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("the error is in db connection ... panic ", err)
	}
	store := db.NewStore(connPool)
	server := api.NewServer(store, logger)

	err = server.Start(config.ServerAdress)
	if err != nil {
		log.Fatal("the error is cannot start server... panic ", err)
	}

}
