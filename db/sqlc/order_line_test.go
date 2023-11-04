package db

import (
	"context"
	"testing"

	"github.com/dohaelsawy/bookStore/util"
	"github.com/stretchr/testify/require"
)

func CreateOrderLine(t *testing.T) OrderLine{
	book := CreateBook(t)
	order := CreateCustomerOrder(t)

	arg := CreateOrderLineParams {
		OrderID: order.CustomerOrderID,
		BookID: book.BookID,
		Price: int32(util.RandomNumber(1,100)),
	}
	line , err := testDB.CreateOrderLine(context.Background() , arg)
	require.NoError(t ,err)
	require.NotNil(t, order)
	require.Equal(t , line.OrderID , arg.OrderID)
	require.Equal(t , line.BookID , arg.BookID)
	
	require.NotZero(t , line.LineID)

	return line
}


func UpdateOrderLine(t *testing.T) OrderLine {
	object := CreateOrderLine(t)
	arg := UpdateOrderLineParams {
		LineID: 1,
		OrderID: object.OrderID,
		BookID: object.BookID,
		Price: int32(util.RandomNumber(1,100)),
	}
	line , err := testDB.UpdateOrderLine(context.Background() , arg)
	require.NoError(t ,err)
	require.NotNil(t, line)
	require.Equal(t , line.OrderID , arg.OrderID)
	require.Equal(t , line.BookID , arg.BookID)
	
	require.NotZero(t , line.LineID)

	return line
}

func TestOrderLineCresting(t *testing.T) {
	CreateOrderLine(t)
}

func TestOrderLineUpdating(t *testing.T){
	UpdateOrderLine(t)
}


func TestGetOrderLine(t *testing.T) {
	line , err := testDB.GetOrderLine(context.Background() , 3)
	require.NoError(t , err)
	require.NotNil(t , line)
}

func TestOrderLineDeleting(t *testing.T) {
	err := testDB.DeleteOrderLine(context.Background(), 4)
	require.NoError(t , err)
}

func TestOrderLineList(t *testing.T){
	lines , err := testDB.ListOrderLines(context.Background())
	require.NoError(t , err)
	require.NotEmpty(t , lines)
	require.NotZero(t , len(lines))
}
