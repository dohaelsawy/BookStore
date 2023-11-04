package db

import (
	"context"
	"testing"

	"github.com/dohaelsawy/bookStore/util"
	"github.com/stretchr/testify/require"
)

func CreateCustomerOrder(t *testing.T) CustomerOrder{

	customer := CreateCustomer(t)
	shippingMrthod := CreateShippingMethod(t)

	arg := CreateCustomerOrderParams {
		CustomerID: customer.CustomerID,
		ShippingMethodID: shippingMrthod.ShippingMethodID,
	}
	customerOrder , err := testDB.CreateCustomerOrder(context.Background() , arg)
	require.NoError(t ,err)
	require.NotNil(t, customerOrder)
	require.Equal(t , customerOrder.CustomerID , arg.CustomerID)
	require.Equal(t , customerOrder.ShippingMethodID , arg.ShippingMethodID)
	
	require.NotZero(t , customerOrder.CustomerOrderID)

	return customerOrder
}


func UpdateCustomerOrder(t *testing.T) CustomerOrder {
	object := CreateCustomerOrder(t)
	arg := UpdateCustomerOrderParams{
		CustomerOrderID: object.CustomerOrderID,
		CustomerID: object.CustomerID,
		ShippingMethodID: int32(util.RandomNumber(1,4)),
	}
	customerOrder , err := testDB.UpdateCustomerOrder(context.Background() , arg)
	require.NoError(t ,err)
	require.NotNil(t, customerOrder)
	require.Equal(t , customerOrder.CustomerID , arg.CustomerID)
	require.Equal(t , customerOrder.ShippingMethodID , arg.ShippingMethodID)
	
	require.NotZero(t , customerOrder.CustomerOrderID)

	return customerOrder
}

func TestCustomerOrderCresting(t *testing.T) {
	CreateCustomerOrder(t)
}

func TestCustomerOrderUpdating(t *testing.T){
	UpdateCustomerOrder(t)
}


func TestGetCustomerOrder(t *testing.T) {
	order , err := testDB.GetOrder(context.Background() , 3)
	require.NoError(t , err)
	require.NotNil(t , order)
}

func TestCustomerOrderDeleting(t *testing.T) {
	err := testDB.DeleteCustomerOrder(context.Background(), 4)
	require.NoError(t , err)
}

func TestCustomerOrderList(t *testing.T){
	orders , err := testDB.ListCustomerOrder(context.Background())
	require.NoError(t , err)
	require.NotEmpty(t , orders)
	require.NotZero(t , len(orders))
}
