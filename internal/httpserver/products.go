package httpserver

import (
	"encoding/json"
	"net/http"
)

func (s *Server) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := s.productService.GetProducts()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	jsonErr := json.NewEncoder(w).Encode(products)

	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
		return
	}
}
