package db

import (
	"context"
	"testing"

	"github.com/dohaelsawy/bookStore/util"
	"github.com/stretchr/testify/require"
)

func TestShippingMethodCreating(t *testing.T) ShippingMethod{
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

func TestShippingMethodUpdating(t *testing.T) ShippingMethod{
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


	return method
}

func TestGetShippingMethod(t *testing.T) ShippingMethod{
	method , err := testDB.GetShippingMehod(context.Background(),2)
	require.NoError(t , err)
	require.NotNil(t , method)

	return method
}

func TestGetListShippingMethod(t *testing.T) *[]ShippingMethod{
	methods , err := testDB.ListShippingMthods(context.Background())
	require.NoError(t , err)
	require.NotNil(t , methods)
	require.NotZero(t , len(methods))

	return &methods
}


func TestShippingMethodDeleting(t *testing.T){
	err := testDB.DeleteShippingMethod(context.Background() , 2)
	require.NoError(t , err)
}
