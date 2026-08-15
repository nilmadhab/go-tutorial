package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Stock       int     `json:"stock"`
	Image       string  `json:"image"`
}

type ProductsData struct {
	Products []Product `json:"products"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello World")
}

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data, err := os.ReadFile("data/products.json")
	if err != nil {
		http.Error(w, "Failed to read products data", http.StatusInternalServerError)
		log.Printf("Error reading file: %v", err)
		return
	}

	var products ProductsData
	if err := json.Unmarshal(data, &products); err != nil {
		http.Error(w, "Failed to parse products data", http.StatusInternalServerError)
		log.Printf("Error parsing JSON: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func main() {
	http.HandleFunc("/", enableCORS(helloHandler))
	http.HandleFunc("/products", enableCORS(productsHandler))

	fmt.Println("Server starting on port 8080...")
	fmt.Println("Endpoints:")
	fmt.Println("  GET /         - Hello World")
	fmt.Println("  GET /products - List all products")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
