package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jonathanmdr/go-hexagonal/adapters/web/dto"
	"github.com/jonathanmdr/go-hexagonal/application"
)

func MakeProductHandlers(router *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	router.Handle("/products/{id}",
		n.With(
			negroni.Wrap(getProduct(service),
		),
	)).Methods("OPTIONS", "GET")
	router.Handle("/products",
		n.With(
			negroni.Wrap(createProduct(service),
		),
	)).Methods("OPTIONS", "POST")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			response.Header().Set("Content-Type", "application/json")
			vars := mux.Vars(request)
			id := vars["id"]
			product, err := service.Get(id)
			if err != nil {
				response.WriteHeader(http.StatusNotFound)
				return
			}
			err = json.NewEncoder(response).Encode(product)
			if err != nil {
				response.WriteHeader(http.StatusInternalServerError)
				return
			}
		},
	)
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			response.Header().Set("Content-Type", "application/json")
			var productDto dto.Product
			err := json.NewDecoder(request.Body).Decode(&productDto)
			if err != nil {
				response.WriteHeader(http.StatusInternalServerError)
				response.Write(jsonError(err.Error()))
				return
			}
			product, err := service.Create(productDto.Name, productDto.Price)
			if err != nil {
				response.WriteHeader(http.StatusInternalServerError)
				response.Write(jsonError(err.Error()))
				return
			}
			err = json.NewEncoder(response).Encode(product)
			if err != nil {
				response.WriteHeader(http.StatusInternalServerError)
				response.Write(jsonError(err.Error()))
				return
			}
		},
	)
}
