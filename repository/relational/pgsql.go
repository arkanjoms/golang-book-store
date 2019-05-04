package relational

import (
	"database/sql"
	"golang-book-store/models"
	"golang-book-store/utils"
)

type BookRepository struct{}

func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) (interface{}, error) {
	rows, err := db.Query("SELECT * FROM books ORDER BY id")
	if err != nil {
		return []models.Book{}, err
	}

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		books = append(books, book)
	}

	return utils.ResultData(books, []models.Book{}, err)
}

func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) (interface{}, error) {
	row := db.QueryRow("SELECT * FROM books where id=$1", id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)

	return utils.ResultData(book, models.Book{}, err)
}

func (b BookRepository) AddBook(db *sql.DB, book models.Book) (interface{}, error) {
	err := db.QueryRow("INSERT INTO books (title, author, year) VALUES ($1,$2, $3) RETURNING id;",
		book.Title,
		book.Author,
		book.Year,
	).Scan(&book.ID)

	return utils.ResultData(book.ID, 0, err)
}

func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) (interface{}, error) {
	result, err := db.Exec("UPDATE books set title = $1, author = $2, year = $3 WHERE id=$4 RETURNING id",
		book.Title,
		book.Author,
		book.Year,
		book.ID,
	)

	r, err := result.RowsAffected()
	if r == 0 && err == nil {
		err = sql.ErrNoRows
	}
	return utils.ResultData(r, 0, err)
}

func (b BookRepository) RemoveBook(db *sql.DB, id int) (interface{}, error) {
	result, err := db.Exec("DELETE FROM books where id=$1", id)

	r, err := result.RowsAffected()
	if r == 0 && err == nil {
		err = sql.ErrNoRows
	}
	return utils.ResultData(r, 0, err)
}
