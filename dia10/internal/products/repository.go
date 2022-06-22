package products

import (
	"database/sql"
	"github.com/mvrdgs/bootcamp-go-dh/dia10/internal/products/models"
	"github.com/mvrdgs/bootcamp-go-dh/dia10/pkg/mysqlStore"
	"log"
)

const (
	InsertProduct  = "INSERT INTO products(name, type, count, price) VALUES(?, ?, ?, ?)"
	GetProduct     = "SELECT id, name, type, count, price FROM products WHERE ID = ?"
	GetAllProducts = "SELECT id, name, type, count, price FROM products"
	UpdateProduct  = "UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?"
	DeleteProduct  = "DELETE FROM products WHERE id = ?"
)

type Repository interface {
	Store(product models.Product) (models.Product, error)
	GetOne(id int) models.Product
	Update(product models.Product) (models.Product, error)
	GetAll() ([]models.Product, error)
	Delete(id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Store(product models.Product) (models.Product, error) {
	db := mysqlStore.StorageDB
	stmt, err := db.Prepare(InsertProduct)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price)
	if err != nil {
		return models.Product{}, err
	}
	insertedId, _ := result.LastInsertId()
	product.ID = int(insertedId)

	return product, nil
}

func (r *repository) GetOne(id int) models.Product {
	db := mysqlStore.StorageDB

	var product models.Product
	rows, err := db.Query(GetProduct, id)
	if err != nil {
		log.Println(err.Error())
		return product
	}

	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Println(err.Error())
			return product
		}
	}
	return product
}

func (r *repository) Update(product models.Product) (models.Product, error) {
	db := mysqlStore.StorageDB
	stmt, err := db.Prepare(UpdateProduct)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price, product.ID)
	if err != nil {
		log.Println(err.Error())
		return models.Product{}, err
	}
	return product, nil
}

func (r *repository) GetAll() ([]models.Product, error) {
	db := mysqlStore.StorageDB

	var products []models.Product

	rows, err := db.Query(GetAllProducts)
	if err != nil {
		log.Println(err.Error())
		return products, err
	}

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Println(err.Error())
			return products, err
		}
		products = append(products, product)
	}
	return products, err
}

func (r *repository) Delete(id int) error {
	db := mysqlStore.StorageDB

	stmt, err := db.Prepare(DeleteProduct)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
