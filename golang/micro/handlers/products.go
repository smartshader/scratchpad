package handlers

import (
	"log"
	"micro/models"
	"net/http"
	"regexp"
	"strconv"
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
	case http.MethodPost:
		h.addProduct(w, r)
		return
	case http.MethodPut:
		h.updateProduct(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	h.log.Println("Handle GET Products")

	lp := models.GetProducts()

	err := lp.ToJSON(w)
	if err != nil {
		h.log.Println(err)
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (h *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	h.log.Println("Handle POST Product")

	product := &models.Product{}

	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	h.log.Printf("Product: %#v", product)

	models.AddProduct(product)
}

func (h *Products) updateProduct(w http.ResponseWriter, r *http.Request) {
	h.log.Println("Handle PUT Product")

	x := regexp.MustCompile("/([0-9]+)")
	g := x.FindAllStringSubmatch(r.URL.Path, -1)

	if len(g) != 1 {
		h.log.Println("Invalid URI more than one id")
		http.Error(w, "Invalid URI", http.StatusBadRequest)
		return
	}

	if len(g[0]) != 2 {
		h.log.Println("Invalid URI more than two capture groups")
		http.Error(w, "Invalid URI", http.StatusBadRequest)
		return
	}

	idString := g[0][1]
	id, err := strconv.Atoi(idString)
	if err != nil {
		h.log.Println("Unable to convert to number")
		http.Error(w, "Invalid URI", http.StatusBadRequest)
		return
	}

	product := &models.Product{}

	err = product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	err = models.UpdateProduct(id, product)
	if err == models.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}
}