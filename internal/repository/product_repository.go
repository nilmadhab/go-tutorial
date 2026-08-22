package repository

import (
	"database/sql"

	"hello-api/internal/models"
)

// ProductRepository handles all database operations for products
//
// POINTER EXPLANATION:
// * (asterisk) - means "pointer to". A pointer holds the memory address of a value.
//   - *sql.DB means "pointer to sql.DB" - we store the address, not a copy
//   - This allows us to share the same DB connection across the app
//
// & (ampersand) - means "address of". Gets the memory address of a variable.
//   - &ProductRepository{} creates a struct and returns its address
//   - This avoids copying the entire struct, just passes the address
//
// WHY USE POINTERS?
// 1. Efficiency: Pass address (8 bytes) instead of copying entire struct
// 2. Mutation: Changes to pointed value affect the original
// 3. Shared state: Multiple parts of code can access the same data
type ProductRepository struct {
	db *sql.DB // *sql.DB = pointer to database connection (shared, not copied)
}

// NewProductRepository creates a new repository instance
// Parameters:
//   - db *sql.DB: pointer to database connection (we receive an address)
//
// Returns:
//   - *ProductRepository: pointer to new repository (we return an address)
func NewProductRepository(db *sql.DB) *ProductRepository {
	// & creates a new struct and returns its memory address
	// Without &: would return a copy (ProductRepository)
	// With &:    returns a pointer (*ProductRepository)
	return &ProductRepository{db: db}
}

// GetAll retrieves all products from the database
// (r *ProductRepository) = method receiver
//   - r is the variable name (like "this" or "self" in other languages)
//   - *ProductRepository means r is a pointer to ProductRepository
//   - This lets the method access r.db (the database connection)
func (r *ProductRepository) GetAll() ([]models.Product, error) {
	rows, err := r.db.Query("SELECT id, name, description, price, category, stock, image FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		// Scan needs pointers (&) so it can write values directly into p's fields
		// &p.ID = "address of p.ID" - Scan writes the database value to this address
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Category, &p.Stock, &p.Image); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// GetByID retrieves a single product by ID
// Returns *models.Product (pointer) so we can return nil if not found
// If we returned models.Product (value), we couldn't return "nothing"
func (r *ProductRepository) GetByID(id int) (*models.Product, error) {
	var p models.Product
	query := "SELECT id, name, description, price, category, stock, image FROM products WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Category, &p.Stock, &p.Image)
	if err == sql.ErrNoRows {
		return nil, nil // nil = no product found (only possible with pointer return type)
	}
	if err != nil {
		return nil, err
	}
	return &p, nil // &p = return address of p, not a copy
}

// Create inserts a new product into the database
// p *models.Product = receives a pointer so we can modify the original
// After insert, we set p.ID - this changes the caller's product because we have its address
func (r *ProductRepository) Create(p *models.Product) error {
	query := `INSERT INTO products (name, description, price, category, stock, image)
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	// Scan(&p.ID) writes the new ID directly into the caller's product
	return r.db.QueryRow(query, p.Name, p.Description, p.Price, p.Category, p.Stock, p.Image).Scan(&p.ID)
}

func (r *ProductRepository) Update(id int, p *models.Product) error {
	query := `UPDATE products SET name=$1, description=$2, price=$3, category=$4, stock=$5, image=$6 WHERE id=$7`
	result, err := r.db.Exec(query, p.Name, p.Description, p.Price, p.Category, p.Stock, p.Image, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	p.ID = id
	return nil
}

func (r *ProductRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
