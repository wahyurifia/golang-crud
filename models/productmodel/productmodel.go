package productmodel

import (
	"CRUD-Golang/config"
	"CRUD-Golang/entities"
)

func GetAll() []entities.Product{
	rows, err := config.DB.Query(`
	SELECT 
	products.id,
	products.name,
	categories.name as category_name,
	products.stock,
	products.description,
	products.created_at,
	products.update_at
	FROM products
	 JOIN categories ON products.category_id = categories.id
	`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []entities.Product

	for rows.Next(){
		var product entities.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Category.Name ,&product.Stock, &product.Description, &product.CreatedAt, &product.UpdateAt); err != nil {
			panic(err)
		}

		products = append(products, product)
	}
	return products
}

func Detail(id int) entities.Product {
	row := config.DB.QueryRow(`
		SELECT 
		products.id,
		products.name,
		categories.name as category_name,
		products.stock,
		products.description,
		products.created_at,
		products.update_at
		FROM products
		JOIN categories ON products.category_id = categories.id
		WHERE products.id = $1
	`, id)

	var product entities.Product
	if err := row.Scan(&product.Id, &product.Name, &product.Category.Name ,&product.Stock, &product.Description, &product.CreatedAt, &product.UpdateAt); err != nil {
		panic(err)
	}

	return product
}


func Create(product entities.Product)bool {
	var lastInsertId int
	 err := config.DB.QueryRow(`
	INSERT INTO products (name, category_id, stock, description, created_at, update_at) 
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
	 product.Name,
	 product.Category.Id,
	 product.Stock,
	 product.Description,
	 product.CreatedAt,
	 product.UpdateAt,
	).Scan(&lastInsertId)

	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Update(id int, product entities.Product)bool {
	query, err := config.DB.Exec(`
	UPDATE products SET
	name = $1,
	category_id = $2,
	stock = $3,
	description = $4,
	update_at = $5
	WHERE id = $6
	`,
	product.Name,
	product.Category.Id,
	product.Stock,
	product.Description,
	product.UpdateAt,
	id,)

	if err != nil {
		panic(err)
	}
	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int)error {
	_, err := config.DB.Exec("DELETE FROM products WHERE id = $1", id)
	return err
}
