package handlers

import (
	"FirstService/service/data"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}

	//handle an Update

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {
	p.l.Printf("Handle %s Product\n", http.MethodGet)
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to encode json", http.StatusInternalServerError)
	}

}

func (p *Products) AddProducts(rw http.ResponseWriter, h http.Request) {
	p.l.Printf("Handle %s Product\n", http.MethodPost)

	prod := &data.Product{}

	err := prod.FromJSON(h.Body)
	if err != nil {
		http.Error(rw, "Unable to decode JSON", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p Products) UpdateProducts(rw http.ResponseWriter, h *http.Request) {

	vars := mux.Vars(h)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert Id", http.StatusBadRequest)
		return
	}

	p.l.Printf("Handle %s Product\n", http.MethodPut, id)
	prod := h.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}

type KeyProduct struct{}

func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
