package cli

import (
	"fmt"
	"github.com/jonathanmdr/go-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		resp, err := create(service, productName, price)
		if err != nil {
			return result, err
		}
		result = resp
	case "enable":
		resp, err := enable(service, productId)
		if err != nil {
			return result, err
		}
		result = resp
	case "disable":
		resp, err := disable(service, productId)
		if err != nil {
			return result, err
		}
		result = resp
	case "get":
		resp, err := get(service, productId)
		if err != nil {
			return result, err
		}
		result = resp
	default:
		result = fmt.Sprintf("The operation '%s' doesn't a valid command", action)
	}
	return result, nil
}

func create(service application.ProductServiceInterface, productName string, productPrice float64) (string, error) {
	result := ""
	product, err := service.Create(productName, productPrice)
	if err != nil {
		return result, err
	}
	result = fmt.Sprintf(
		"Product ID %s with name %s has been created with the price %f and status %s.",
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)
	return result, nil
}

func enable(service application.ProductServiceInterface, productId string) (string, error) {
	result := ""
	product, err := service.Get(productId)
	if err != nil {
		return result, err
	}
	res, err := service.Enable(product)
	if err != nil {
		return result, err
	}
	result = fmt.Sprintf("Product %s has been enabled.", res.GetName())
	return result, nil
}

func disable(service application.ProductServiceInterface, productId string) (string, error) {
	result := ""
	product, err := service.Get(productId)
	if err != nil {
		return result, err
	}
	res, err := service.Disable(product)
	if err != nil {
		return result, err
	}
	result = fmt.Sprintf("Product %s has been disabled.", res.GetName())
	return result, nil
}

func get(service application.ProductServiceInterface, productId string) (string, error) {
	result := ""
	product, err := service.Get(productId)
	if err != nil {
		return result, err
	}
	result = fmt.Sprintf(
		"Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)
	return result, nil
}
