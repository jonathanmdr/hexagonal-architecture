package main

import (
	"database/sql"
	adapters "github.com/jonathanmdr/go-hexagonal/adapters/db"
	"github.com/jonathanmdr/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := adapters.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product Example", 30)
	
	productService.Enable(product)
}