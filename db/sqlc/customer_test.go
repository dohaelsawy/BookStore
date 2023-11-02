package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/dohaelsawy/bookStore/util"
)

func TestCustomerCreating(t *testing.T) {
	arg := CreatecustomerParams{
		FirstName: util.RandomString(6),
		LastName: util.RandomString(6),
		Email: util.RandomEmail(),
		Password: util.RandomPssword(),
		City: "cairo",
		PhoneNumber: util.RandomPhoneNumber(),
	}
	customer , err := testDB.Createcustomer(context.Background(), arg)
	require.NoError(t , err)
	require.NotEmpty(t, customer)
	require.Equal(t, arg.FirstName, customer.FirstName)
	require.Equal(t, arg.LastName, customer.LastName)
	require.Equal(t, arg.Email,customer.Email)
	require.Equal(t, arg.Password, customer.Password)
	require.Equal(t, arg.City, customer.City)
	require.Equal(t, arg.PhoneNumber, customer.PhoneNumber)
	require.NotZero(t , customer.CustomerID)
}

func TestCustomerUpdating(t *testing.T) {
	arg := UpdatecustomerParams{
		CustomerID: 3,
		FirstName: util.RandomString(6),
		LastName: util.RandomString(6),
		Email: util.RandomEmail(),
		Password: util.RandomPssword(),
		City: "cairo",
		PhoneNumber: util.RandomString(8),
	}
	customer , err := testDB.Updatecustomer(context.Background(), arg)
	require.NoError(t , err)
	require.NotEmpty(t, customer)
	require.Equal(t, arg.FirstName, customer.FirstName)
	require.Equal(t, arg.LastName, customer.LastName)
	require.Equal(t, arg.Email,customer.Email)
	require.Equal(t, arg.Password, customer.Password)
	require.Equal(t, arg.City, customer.City)
	require.Equal(t, arg.PhoneNumber, customer.PhoneNumber)
	require.NotZero(t , customer.CustomerID)
}

func TestGetCustomer(t *testing.T){
	customer , err := testDB.GetCustomer(context.Background() , 3)
	require.NoError(t , err)
	require.NotNil(t , customer)
}

func TestCustomerDeleting(t *testing.T) {
	err := testDB.Deletecustomer(context.Background(), 1)
	require.NoError(t , err)
}

func TestCustomerAllList(t *testing.T){
	customers , err := testDB.ListCustomers(context.Background())
	require.NoError(t , err)
	require.NotEmpty(t , customers)
	require.NotZero(t , len(customers))
}