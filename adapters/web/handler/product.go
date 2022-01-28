package handler

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jonathanmdr/go-hexagonal/adapters/web/dto"
	"github.com/jonathanmdr/go-hexagonal/application"
	"net/http"
)

func MakeProductHandlers(router *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	router.Handle("/health",
		n.With(
			negroni.Wrap(health()),
		)).Methods("OPTIONS", "GET")
	router.Handle("/products/{id}",
		n.With(
			negroni.Wrap(getProduct(service)),
		)).Methods("OPTIONS", "GET")
	router.Handle("/products",
		n.With(
			negroni.Wrap(createProduct(service)),
		)).Methods("OPTIONS", "POST")
	router.Handle("/products/{id}/enable",
		n.With(
			negroni.Wrap(enableProduct(service)),
		)).Methods("OPTIONS", "PUT")
	router.Handle("/products/{id}/disable",
		n.With(
			negroni.Wrap(disableProduct(service)),
		)).Methods("OPTIONS", "PUT")
}

func health() http.Handler {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			response.WriteHeader(http.StatusOK)
		},
	)
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			response.Header().Set("Content-Type", "application/json")
			vars := mux.Vars(request)
			id := vars["id"]
			product, err := service.Get(id)
			if err != nil {
				hidrateError(response, err, http.StatusNotFound)
				return
			}
			err = json.NewEncoder(response).Encode(product)
			if err != nil {
				hidrateError(response, err, http.StatusInternalServerError)
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
				hidrateError(response, err, http.StatusInternalServerError)
				return
			}
			product, err := service.Create(productDto.Name, productDto.Price)
			if err != nil {
				hidrateError(response, err, http.StatusInternalServerError)
				return
			}
			err = json.NewEncoder(response).Encode(product)
			if err != nil {
				hidrateError(response, err, http.StatusInternalServerError)
				return
			}
			response.WriteHeader(http.StatusCreated)
		},
	)
}

func enableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			response.Header().Set("Content-Type", "application/json")
			vars := mux.Vars(request)
			id := vars["id"]
			product, err := service.Get(id)
			if err != nil {
				hidrateError(response, err, http.StatusNotFound)
				return
			}
			result, err := service.Enable(product)
			if err != nil {
				hidrateError(response, err, http.StatusInternalServerError)
				return
			}
			err = json.NewEncoder(response).Encode(result)
			if err != nil {
				hidrateError(response, err, http.StatusInternalServerError)
				return
			}
		},
	)
}

func disableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			response.Header().Set("Content-Type", "application/json")
			vars := mux.Vars(request)
			id := vars["id"]
			product, err := service.Get(id)
			if err != nil {
				hidrateError(response, err, http.StatusNotFound)
				return
			}
			result, err := service.Disable(product)
			if err != nil {
				hidrateError(response, err, http.StatusInternalServerError)
				return
			}
			err = json.NewEncoder(response).Encode(result)
			if err != nil {
				hidrateError(response, err, http.StatusInternalServerError)
				return
			}
		},
	)
}

func hidrateError(response http.ResponseWriter, err error, httpStatus int) {
	response.WriteHeader(httpStatus)
	response.Write(jsonError(err.Error()))
}
