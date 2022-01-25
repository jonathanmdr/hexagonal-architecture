package application_test

import (
	"github.com/jonathanmdr/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	error := product.Enable()
	require.Nil(t, error)

	product.Price = 0

	error = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", error.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	error := product.Disable()
	require.Nil(t, error)

	product.Price = 10

	error = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disabled", error.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, error := product.IsValid()
	require.Nil(t, error)

	product.Status = "INVALID"
	_, error = product.IsValid()
	require.Equal(t, "the status must be enable or disabled", error.Error())

	product.Status = application.ENABLED
	_, error = product.IsValid()
	require.Nil(t, error)

	product.Price = -10
	_, error = product.IsValid()
	require.Equal(t, "the proce must be greater or equal to zero", error.Error())

	product.ID = "asdfg"
	_, error = product.IsValid()
	require.Error(t, error)

	product.Name = ""
	_, error = product.IsValid()
	require.Error(t, error)
}
