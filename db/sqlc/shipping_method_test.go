package db

import (
	"context"
	"testing"

	"github.com/dohaelsawy/bookStore/util"
	"github.com/stretchr/testify/require"
)


func CreateShippingMethod(t *testing.T) ShippingMethod{
	arg := CreateShippingMehodParams{
		Name: util.RandomShippingMethod(),
		Cost: int32(util.RandomNumber(100 , 500)),
	}

	method , err := testDB.CreateShippingMehod(context.Background() , arg)

	require.NoError(t , err)
	require.NotNil(t , method)
	require.Equal(t , method.Name , arg.Name)
	require.Equal(t , method.Cost , arg.Cost)

	require.NotZero(t , method.ShippingMethodID)
	return method
}

func TestShippingMethodCreating(t *testing.T) {
	CreateShippingMethod(t)
}

func TestShippingMethodUpdating(t *testing.T) {
	arg := UpdateShippingMehodParams{
		ShippingMethodID: 1,
		Name: util.RandomShippingMethod(),
		Cost: int32(util.RandomNumber(100 , 500)),
	}

	method , err := testDB.UpdateShippingMehod(context.Background() , arg)

	require.NoError(t , err)
	require.NotNil(t , method)
	require.Equal(t , method.Name , arg.Name)
	require.Equal(t , method.Cost , arg.Cost)

	require.NotZero(t , method.ShippingMethodID)
}

func TestGetShippingMethod(t *testing.T) {
	method , err := testDB.GetShippingMehod(context.Background(),2)
	require.NoError(t , err)
	require.NotNil(t , method)
}

func TestGetListShippingMethod(t *testing.T) {
	methods , err := testDB.ListShippingMthods(context.Background())
	require.NoError(t , err)
	require.NotNil(t , methods)
	require.NotZero(t , len(methods))
}


func TestShippingMethodDeleting(t *testing.T){
	err := testDB.DeleteShippingMethod(context.Background() , 2)
	require.NoError(t , err)
}
