package main

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// This is a placeholder for GORM initialization and usage.
	// In a real application, you would initialize the database connection here
	// and perform operations such as creating, reading, updating, and deleting records.

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	ctx := context.Background()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	err = gorm.G[Product](db).Create(ctx, &Product{Code: "D42", Price: 100})
	if err != nil {
		panic("failed to create product")
	}

}
