package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/dohaelsawy/bookStore/util"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

var testDB *Queries

func TestMain(m *testing.M){
	config ,err := util.LoadAppEnv("../..")
	if err != nil {
		log.Fatal("can't load env variables ... panic " , err)
	}
	conn , err := pgx.Connect(context.Background() , config.DBSource)
	if err != nil {
		log.Fatal("the error is in db connection ... panic " , err)
	} 
	defer conn.Close(context.Background())
	testDB = New(conn)
	os.Exit(m.Run())
}