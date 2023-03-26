package repository

import (
	"database/sql"
	"quiz-3/structs"
)

func GetAllCategories(db *sql.DB) (err error, results []structs.Category) {
	sql := "SELECT * FROM category"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var category = structs.Category{}

		err = rows.Scan(&category.ID, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			panic(err)
		}

		results = append(results, category)
	}
	return
}

func InsertCategories(db *sql.DB, category structs.Category) (err error) {
	sql := "INSERT INTO category (id, nama,  created_at,update_at) VALUES ($1,$2,$3,$4)"

	errs := db.QueryRow(sql, category.ID, category.Name, category.CreateAt, category.UpdateAt)

	return errs.Err()
}

func UpdateCategories(db *sql.DB, category structs.Category) (err error) {
	sql := "UPDATE category SET nama = $1, create_at = $2, update_at=$3 WHERE id = $4"

	errs := db.QueryRow(sql, category.Name, category.CreateAt, category.UpdateAt, category.ID)

	return errs.Err()
}

func DeleteCategories(db *sql.DB, category structs.Category) (err error) {
	sql := "DELETE FROM category WHERE id = $1"

	errs := db.QueryRow(sql, category.ID)

	return errs.Err()
}

func GetBooksByCategoryIDFromDatabase(db *sql.DB) (err error, results []structs.Book) {
	sql := "SELECT * FROM books WHERE category_id = $1"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var books = structs.Book{}

		err = rows.Scan(&books.ID, &books.Title, &books.Description, &books.ImageURL, &books.ReleaseYear, &books.Price,
			&books.TotalPage, &books.Thickness, &books.CreatedAt, &books.UpdatedAt, &books.CategoryID)
		if err != nil {
			panic(err)
		}
		results = append(results, books)
	}
	return

}
