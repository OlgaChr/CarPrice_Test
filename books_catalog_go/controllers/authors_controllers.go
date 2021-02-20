package controllers

import (
	"books_catalog_go/dto"
	"books_catalog_go/models"
	u "books_catalog_go/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var CreateAuthor = func(w http.ResponseWriter, r *http.Request) {
	author := &models.Author{}

	err := json.NewDecoder(r.Body).Decode(author)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := author.Create()
	u.Respond(w, resp)
}

var UpdateAuthor = func(w http.ResponseWriter, r *http.Request) {
	author := &models.Author{}

	err := json.NewDecoder(r.Body).Decode(author)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := author.Update()
	u.Respond(w, resp)
}

var DeleteAuthor = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//Переданный параметр пути не является целым числом
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.DeleteAuthor(uint(id))
	if data == nil {
		resp := u.Message(true, "success")
		u.Respond(w, resp)
	} else {
		resp := u.Message(false, "error")
		resp["data"] = data
		u.Respond(w, resp)
	}
}

var GetAuthor = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//Переданный параметр пути не является целым числом
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetAuthor(uint(id))
	if data != nil {
		resp := u.Message(true, "success")
		resp["data"] = data
		u.Respond(w, resp)
	} else {
		resp := u.Message(false, "error")
		u.Respond(w, resp)
	}
}

var GetAuthors = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetAuthors()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetAuthorsWithSearch = func(w http.ResponseWriter, r *http.Request) {
	authorCondition := &dto.AuthorCondition{}
	err := json.NewDecoder(r.Body).Decode(authorCondition)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	data := models.GetAuthorsWithSearch(authorCondition)
	if data != nil {
		resp := u.Message(true, "success")
		resp["data"] = data
		u.Respond(w, resp)
	} else {
		resp := u.Message(false, "error")
		u.Respond(w, resp)
	}
}
