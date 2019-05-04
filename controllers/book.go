package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"golang-book-store/models"
	"golang-book-store/repository/relational"
	"golang-book-store/utils"
	"net/http"
	"strconv"
)

type Controller struct{}

var books []models.Book
var repository relational.BookRepository

func init() {
	repository = relational.BookRepository{}
}

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		books = []models.Book{}
		data, err := repository.GetBooks(db, book, books)
		utils.SendResult(w, data, err)
	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		data, err := repository.GetBook(db, book, id)
		utils.SendResult(w, data, err)
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book

		_ = json.NewDecoder(r.Body).Decode(&book)

		if book.Author == "" || book.Title == "" || book.Year == "" {
			utils.SendBadRequest(w, errors.New("enter missing fields"))
			return
		}

		data, err := repository.AddBook(db, book)
		utils.SendResult(w, data, err)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		_ = json.NewDecoder(r.Body).Decode(&book)

		if book.ID == 0 || book.Author == "" || book.Title == "" || book.Year == "" {
			utils.SendBadRequest(w, errors.New("all fields age required"))
			return
		}

		data, err := repository.UpdateBook(db, book)
		utils.SendResult(w, data, err)
	}
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		data, err := repository.RemoveBook(db, id)
		utils.SendResult(w, data, err)
	}
}
