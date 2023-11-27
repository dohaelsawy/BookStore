package db

import (
	"context"
	"log"
	"os"
	"testing"
	_"github.com/lib/pq"
	"github.com/jackc/pgx/v5"
)
const (
	dataSource = "postgresql://root:root@localhost:5432/postgres?sslmode=disable"
)
var testDB *Queries

func TestMain(m *testing.M){
	conn , err := pgx.Connect(context.Background() , dataSource)
	if err != nil {
		log.Fatal("the error is in db connection ... panic " , err)
	} 
	defer conn.Close(context.Background())
	testDB = New(conn)
	os.Exit(m.Run())
}