package handlers

import (
	"context"
	"fmt"
	"log"
	"micro/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	log *log.Logger
}

func NewProducts(log *log.Logger) *Products {
	return &Products{
		log: log,
	}
}

func (h *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	h.log.Println("Handle GET Products")

	lp := models.GetProducts()

	err := lp.ToJSON(w)
	if err != nil {
		h.log.Println(err)
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (h *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	h.log.Println("Handle POST Product")

	product := r.Context().Value(KeyProduct{}).(models.Product)

	h.log.Printf("Product: %#v", product)

	models.AddProduct(&product)
}

func (h *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	h.log.Println("Handle PUT Product")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	product := r.Context().Value(KeyProduct{}).(models.Product)

	err = models.UpdateProduct(id, &product)
	if err == models.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct {}

func (h *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		product := models.Product{}

		err := product.FromJSON(r.Body)
		if err != nil {
			h.log.Println("[Error] deserializing product", err)
			http.Error(w, "Error reading product", http.StatusBadRequest)
			return
		}

		err = product.Validate()
		if err != nil {
			h.log.Println("[Error] validating product", err)
			http.Error(
				w,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}