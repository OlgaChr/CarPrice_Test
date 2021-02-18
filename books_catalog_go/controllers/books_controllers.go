package controllers

import (
	"books_catalog_go/types"
	"encoding/json"
	"net/http"
	"strconv"

	"books_catalog_go/models"
	u "books_catalog_go/utils"

	"github.com/gorilla/mux"
)

var CreateBook = func(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}

	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := book.Create()
	u.Respond(w, resp)
}

var UpdateBook = func(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}

	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := book.Update()
	u.Respond(w, resp)
}

var DeleteBook = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//Переданный параметр пути не является целым числом
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.DeleteBook(uint(id))
	if data == nil {
		resp := u.Message(true, "success")
		u.Respond(w, resp)
	} else {
		resp := u.Message(false, "error")
		resp["data"] = data
		u.Respond(w, resp)
	}
}

var GetBook = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//Переданный параметр пути не является целым числом
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetBook(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetBooks = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetBooks()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetBooksWithSearch = func(w http.ResponseWriter, r *http.Request) {
	bookCondition := &types.BookCondition{}
	err := json.NewDecoder(r.Body).Decode(bookCondition)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	data := models.GetBooksWithSearch(bookCondition)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
