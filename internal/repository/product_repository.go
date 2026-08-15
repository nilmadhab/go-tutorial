package repository

import (
	"database/sql"

	"hello-api/internal/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	rows, err := r.db.Query("SELECT id, name, description, price, category, stock, image FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
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

func (r *ProductRepository) GetByID(id int) (*models.Product, error) {
	var p models.Product
	query := "SELECT id, name, description, price, category, stock, image FROM products WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Category, &p.Stock, &p.Image)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) Create(p *models.Product) error {
	query := `INSERT INTO products (name, description, price, category, stock, image)
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
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
