package handler

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jonathanmdr/go-hexagonal/application"
	"net/http"
)

func MakeProductHandlers(router *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	router.Handle("/products/{id}",
		n.With(
			negroni.Wrap(getProduct(service),
		),
	)).Methods("OPTIONS", "GET")
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
