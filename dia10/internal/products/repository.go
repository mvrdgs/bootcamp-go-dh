package products

import (
	"database/sql"
	"github.com/mvrdgs/bootcamp-go-dh/dia10/internal/products/models"
	"log"
)

const (
	InsertProduct     = "INSERT INTO products(name, type, count, price) VALUES(?, ?, ?, ?)"
	GetProduct        = "SELECT id, name, type, count, price FROM products WHERE ID = ?"
	GetAllProducts    = "SELECT id, name, type, count, price FROM products"
	UpdateProduct     = "UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?"
	UpdateProductName = "UPDATE products SET name = ? WHERE id = ?"
	DeleteProduct     = "DELETE FROM products WHERE id = ?"
)

type Repository interface {
	Store(product models.Product) (models.Product, error)
	GetOne(id int) models.Product
	Update(product models.Product) (models.Product, error)
	UpdateName(product models.Product) (models.Product, error)
	GetAll() ([]models.Product, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(product models.Product) (models.Product, error) {
	stmt, err := r.db.Prepare(InsertProduct)
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
	var product models.Product
	rows, err := r.db.Query(GetProduct, id)
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
	stmt, err := r.db.Prepare(UpdateProduct)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price, product.ID)
	if err != nil {
		log.Println(err.Error())
		return models.Product{}, err
	}
	return product, nil
}

func (r *repository) UpdateName(product models.Product) (models.Product, error) {
	stmt, err := r.db.Prepare(UpdateProductName)
	if err != nil {
		log.Fatalln(err.Error())
		return models.Product{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.ID)
	if err != nil {
		log.Println(err.Error())
		return models.Product{}, err
	}

	return product, nil
}

func (r *repository) GetAll() ([]models.Product, error) {
	var products []models.Product

	rows, err := r.db.Query(GetAllProducts)
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
	stmt, err := r.db.Prepare(DeleteProduct)
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
