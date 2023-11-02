package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/dohaelsawy/bookStore/util"
)

func TestBookCreating(t *testing.T) {
	arg := CreateBookParams{
		Name:        util.RandomString(10),
		PublishDate: "2021-01-01",
		Price:       int32(util.RandomNumber(100 , 400)),
		Sku:         "abc-def-ghi",
		Description: util.RandomString(50),
	}
	book, err := testDB.CreateBook(context.Background(), arg)
	require.NoError(t , err)
	require.NotEmpty(t, book)
	require.Equal(t, arg.Name, book.Name)
	require.Equal(t, arg.PublishDate, book.PublishDate)
	require.Equal(t, arg.Price, book.Price)
	require.Equal(t, arg.Sku, book.Sku)
	require.Equal(t, arg.Description, book.Description)
	require.NotZero(t , book.BookID)
	require.NotZero(t , book.CreatedAt)
	require.NotZero(t , book.UpdatedAt)
}

func TestBookUpdating(t *testing.T) {
	arg := UpdateBookParams{
		BookID: 	2,
		Name:       util.RandomString(10),
		PublishDate: "2021-01-01",
		Price:       int32(util.RandomNumber(100 , 400)),
		Sku:         "abc-def-ghi",
		Description: "test description",
	}
	book, err := testDB.UpdateBook(context.Background(), arg)
	require.NoError(t , err)
	require.NotEmpty(t, book)
	require.Equal(t, arg.Name, book.Name)
	require.Equal(t, arg.PublishDate, book.PublishDate)
	require.Equal(t, arg.Price, book.Price)
	require.Equal(t, arg.Sku, book.Sku)
	require.Equal(t, arg.Description, book.Description)
	require.NotZero(t , book.BookID)
	require.NotZero(t , book.CreatedAt)
	require.NotZero(t , book.UpdatedAt)
}

func TestBookDeleting(t *testing.T) {
	err := testDB.DeleteBook(context.Background(), 1)
	require.NoError(t , err)
}

func TestBookAllList(t *testing.T){
	books , err := testDB.ListBooks(context.Background())
	require.NoError(t , err)
	require.NotEmpty(t , books)
	require.NotZero(t , len(books))
}