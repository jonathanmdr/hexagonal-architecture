package cli_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/jonathanmdr/go-hexagonal/adapters/cli"
	"github.com/jonathanmdr/go-hexagonal/application"
	mock_application "github.com/jonathanmdr/go-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	productId := "abc"
	productName := "Product"
	productPrice := 25.99
	productStatus := application.ENABLED

	productMock := mock_application.NewMockProductInterface(controller)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	serviceMock := mock_application.NewMockProductServiceInterface(controller)
	serviceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	expected := fmt.Sprintf(
		"Product ID %s with name %s has been created with the price %f and status %s.",
		productId,
		productName,
		productPrice,
		productStatus,
	)
	actual, err := cli.Run(serviceMock, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expected, actual)

	expected = fmt.Sprintf("Product %s has been enabled.", productName)
	actual, err = cli.Run(serviceMock, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expected, actual)

	expected = fmt.Sprintf("Product %s has been disabled.", productName)
	actual, err = cli.Run(serviceMock, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expected, actual)

	expected = fmt.Sprintf(
		"Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		productId,
		productName,
		productPrice,
		productStatus,
	)
	actual, err = cli.Run(serviceMock, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expected, actual)

	expected = fmt.Sprintf("The operation '%s' doesn't a valid command", "other")
	actual, err = cli.Run(serviceMock, "other", "", "", 0)
	require.Nil(t, err)
	require.Equal(t, expected, actual)
}
