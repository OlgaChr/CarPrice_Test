package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"books_catalog_go/controllers"
	_ "books_catalog_go/models"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")            		// список книг
	router.HandleFunc("/api/books", controllers.GetBooksWithSearch).Methods("POST") 		// список книг с поиском по автору и пагинацией
	router.HandleFunc("/api/book/{id:[0-9]+}", controllers.GetBook).Methods("GET")  		// получение книги
	router.HandleFunc("/api/book", controllers.CreateBook).Methods("POST") 				// добавление книги
	router.HandleFunc("/api/book", controllers.UpdateBook).Methods("PUT")           		// обновление книги
	router.HandleFunc("/api/book/{id:[0-9]+}", controllers.DeleteBook).Methods("DELETE") // удаление книги

	router.HandleFunc("/api/authors", controllers.GetAuthors).Methods("GET")            		// список авторов
	router.HandleFunc("/api/authors", controllers.GetAuthorsWithSearch).Methods("POST") 		// список авторов с поиском по автору и пагинацией
	router.HandleFunc("/api/author/{id:[0-9]+}", controllers.GetAuthor).Methods("GET")  		// получение автора
	router.HandleFunc("/api/author", controllers.CreateAuthor).Methods("POST") 				// добавление автора
	router.HandleFunc("/api/author", controllers.UpdateAuthor).Methods("PUT")           		// обновление автора
	router.HandleFunc("/api/author/{id:[0-9]+}", controllers.DeleteAuthor).Methods("DELETE") // удаление автора

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Print(err)
	}
}
