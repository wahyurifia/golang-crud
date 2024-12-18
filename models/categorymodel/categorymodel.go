package categorymodel

import (
	"CRUD-Golang/config"
	"CRUD-Golang/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query("SELECT * FROM categories")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next(){
		var category entities.Category
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdateAt)
		if err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}
	return categories
}

func Create(category entities.Category) bool {
	var lastInsertId int
	err := config.DB.QueryRow(`
	INSERT INTO categories (name, created_at, update_at)
	VALUES ($1, $2, $3) RETURNING id`,
	category.Name, category.CreatedAt, category.UpdateAt).Scan(&lastInsertId)

	if err != nil {
		panic(err)
	}

	return lastInsertId > 0 
}

func Detail(id int) entities.Category{
	row := config.DB.QueryRow(`SELECT id, name FROM categories WHERE id = $1`, id)

	var category entities.Category
	if err := row.Scan(&category.Id, &category.Name); err != nil {
		panic(err.Error())
	}
	return category
}

func Update(id int, category entities.Category) bool {
	query, err := config.DB.Exec(`
	UPDATE categories SET name = $1, update_at = $2 WHERE id = $3
	`, category.Name, category.UpdateAt, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM categories WHERE id = $1`, id)
	
	return err
}