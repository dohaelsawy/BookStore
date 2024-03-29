// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	CreateBook(ctx context.Context, arg CreateBookParams) (Book, error)
	CreateCustomerOrder(ctx context.Context, arg CreateCustomerOrderParams) (CustomerOrder, error)
	CreateOrderLine(ctx context.Context, arg CreateOrderLineParams) (OrderLine, error)
	CreatePaymentTranscation(ctx context.Context, arg CreatePaymentTranscationParams) (TransactionPayment, error)
	CreateShippingMehod(ctx context.Context, arg CreateShippingMehodParams) (ShippingMethod, error)
	Createcustomer(ctx context.Context, arg CreatecustomerParams) (Customer, error)
	DeleteBook(ctx context.Context, bookID int32) error
	DeleteCustomerOrder(ctx context.Context, customerOrderID int32) error
	DeleteOrderLine(ctx context.Context, lineID int32) error
	DeletePaymentTranscation(ctx context.Context, transactionID int32) error
	DeleteShippingMethod(ctx context.Context, shippingMethodID int32) error
	Deletecustomer(ctx context.Context, customerID int32) error
	GetBook(ctx context.Context, bookID int32) (Book, error)
	GetCustomer(ctx context.Context, customerID int32) (Customer, error)
	GetOrder(ctx context.Context, customerOrderID int32) (CustomerOrder, error)
	GetOrderLine(ctx context.Context, lineID int32) (OrderLine, error)
	GetPaymentTransaction(ctx context.Context, transactionID int32) (TransactionPayment, error)
	GetShippingMehod(ctx context.Context, shippingMethodID int32) (ShippingMethod, error)
	ListBooks(ctx context.Context) ([]Book, error)
	ListCustomerOrder(ctx context.Context) ([]CustomerOrder, error)
	ListCustomers(ctx context.Context) ([]Customer, error)
	ListOrderLines(ctx context.Context) ([]OrderLine, error)
	ListPaymentTranscation(ctx context.Context) ([]TransactionPayment, error)
	ListShippingMthods(ctx context.Context) ([]ShippingMethod, error)
	UpdateBook(ctx context.Context, arg UpdateBookParams) (Book, error)
	UpdateCustomerOrder(ctx context.Context, arg UpdateCustomerOrderParams) (CustomerOrder, error)
	UpdateOrderLine(ctx context.Context, arg UpdateOrderLineParams) (OrderLine, error)
	UpdateShippingMehod(ctx context.Context, arg UpdateShippingMehodParams) (ShippingMethod, error)
	Updatecustomer(ctx context.Context, arg UpdatecustomerParams) (Customer, error)
}

var _ Querier = (*Queries)(nil)
