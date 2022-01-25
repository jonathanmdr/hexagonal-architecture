package db_test

import (
	"database/sql"
	"github.com/jonathanmdr/go-hexagonal/adapters/db"
	"github.com/jonathanmdr/go-hexagonal/application"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products(
			"id" string,
			"name" string,
			"price" float,
			"status" string
		);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES("abc", "First Product", 0, "disabled");`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "First Product", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product := application.NewProduct()
	product.Name = "Product Saved"
	product.Price = 10
	result, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, "Product Saved", result.GetName())
	require.Equal(t, 10.0, result.GetPrice())
	require.Equal(t, "disabled", result.GetStatus())

	product.Name = "Product Updated"
	product.Price = 20
	product.Status = application.ENABLED

	result, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, "Product Updated", result.GetName())
	require.Equal(t, 20.0, result.GetPrice())
	require.Equal(t, "enabled", result.GetStatus())
}
