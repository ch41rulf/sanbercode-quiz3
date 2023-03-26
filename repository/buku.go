package repository

import (
	"database/sql"
	"errors"
	"quiz-3/structs"
	"regexp"
)

func GetBooks(db *sql.DB) (err error, results []structs.Book) {
	sql := "SELECT * FROM books"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var books = structs.Book{}

		err = rows.Scan(&books.ID, &books.Title, &books.Description, &books.ReleaseYear)
		if err != nil {
			panic(err)
		}

		results = append(results, books)
	}
	return
}

func InsertsBooks(db *sql.DB, books structs.Book) (err error) {

	urlRegexp := regexp.MustCompile(`^https?://`)
	if !urlRegexp.MatchString(books.ImageURL) && (books.ReleaseYear < 1980 || books.ReleaseYear > 2021) {
		return errors.New("salah URL format dan tahun harus 1980 dan 2021")
	} else if !urlRegexp.MatchString(books.ImageURL) {
		return errors.New("salah")
	} else if books.ReleaseYear < 1980 || books.ReleaseYear > 2021 {
		return errors.New("tahun harus 1980 dan 2021")
	}

	if books.TotalPage <= 100 {
		books.Thickness = "tipis"
	} else if books.TotalPage <= 200 {
		books.Thickness = "sedang"
	} else {
		books.Thickness = "tebal"
	}
	sql := "INSERT INTO books (title, description, image_url, release_year, price, total_page, category_id, thickness, created_at )VALUES ($1,$2,$3,$4,$5,$6,$7,$8,NOW())"
	errs := db.QueryRow(sql, books.Title, books.Description, books.ImageURL, books.ReleaseYear, books.Price, books.TotalPage, books.CategoryID, books.Thickness, books.CreatedAt)
	return errs.Err()

}

func UpdateBooks(db *sql.DB, books structs.Book) (err error) {
	sql := "UPDATE books SET title = $1, descriptione_at = $2, image_url=$3, release_year=$4,price=$5,total_page=$6,category_id=$7, update_at=NOW() WHERE id = $9"

	errs := db.QueryRow(sql, books.Title, books.Description, books.ImageURL, books.ReleaseYear, books.Price, books.TotalPage, books.CategoryID, books.UpdatedAt, books.ID)

	return errs.Err()
}

func DeleteBooks(db *sql.DB, books structs.Book) (err error) {
	sql := "DELETE FROM books WHERE id = $1"

	errs := db.QueryRow(sql, books)

	return errs.Err()
}
