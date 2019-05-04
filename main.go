package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"golang-book-store/controllers"
	"golang-book-store/driver"
	"log"
	"net/http"
	"os"
)

var db *sql.DB

func init() {
	_ = gotenv.Load()
	db = driver.ConnectDB()
}

func main() {

	router := mux.NewRouter()
	controller := controllers.Controller{}

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	port := os.Getenv("PORT")
	fmt.Println("Server is running at port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
