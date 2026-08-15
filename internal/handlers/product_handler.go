package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"hello-api/internal/models"
	"hello-api/internal/repository"
)

type ProductHandler struct {
	repo *repository.ProductRepository
}

func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

// GetAll handles GET /products
func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	products, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		log.Printf("Error querying products: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ProductsResponse{Products: products})
}

// Create handles POST /products/create
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var p models.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error decoding request: %v", err)
		return
	}

	if err := h.repo.Create(&p); err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		log.Printf("Error inserting product: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

// GetByID handles GET /products/:id
func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := h.extractID(r)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, "Failed to fetch product", http.StatusInternalServerError)
		log.Printf("Error querying product: %v", err)
		return
	}
	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// Update handles PUT /products/:id
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := h.extractID(r)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var p models.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.repo.Update(id, &p); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		log.Printf("Error updating product: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

// Delete handles DELETE /products/:id
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := h.extractID(r)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		log.Printf("Error deleting product: %v", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// HandleSingle routes requests for /products/:id based on method
func (h *ProductHandler) HandleSingle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r)
	case http.MethodPut:
		h.Update(w, r)
	case http.MethodDelete:
		h.Delete(w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) extractID(r *http.Request) (int, error) {
	idStr := strings.TrimPrefix(r.URL.Path, "/products/")
	return strconv.Atoi(idStr)
}
