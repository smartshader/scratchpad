package handlers

import (
	"log"
	"micro/models"
	"net/http"
)

type Products struct {
	log *log.Logger
}

func NewProducts(log *log.Logger) *Products {
	return &Products{
		log: log,
	}
}

func (h *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getProducts(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (h *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := models.GetProducts()

	err := lp.ToJSON(w)
	if err != nil {
		h.log.Println(err)
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
		return
	}
}