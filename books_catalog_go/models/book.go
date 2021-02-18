package models

import (
	"books_catalog_go/types"
	u "books_catalog_go/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name            string
	PublicationYear uint
	Summary         string
	AuthorID        *uint //[]*Author `gorm:"many2many:authors_books;"`
}

func (book *Book) Create() map[string]interface{} {
	GetDB().Create(book)

	resp := u.Message(true, "success")
	resp["book"] = book
	return resp
}

func (book *Book) Update() map[string]interface{} {
	GetDB().Model(&book).Update(book)

	resp := u.Message(true, "success")
	resp["author"] = book
	return resp
}

func DeleteBook(id uint) error {
	book := GetBook(id)
	return GetDB().Delete(&book).Error
}

func GetBook(id uint) *Book {
	book := &Book{}
	err := GetDB().Table("books").Where("id = ?", id).First(book).Error
	if err != nil {
		return nil
	}
	return book
}

func GetBooks() []*Book {
	books := make([]*Book, 0)
	err := GetDB().Table("books").Find(&books).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return books
}

func GetBooksWithSearch(bookCondition *types.BookCondition) []*Book {
	books := make([]*Book, 0)
	db := GetDB().Table("books")

	if bookCondition.AuthorID > 0 {
		db = db.Where("author_id = ?", bookCondition.AuthorID)
	}
	if bookCondition.Page > 0 && bookCondition.Size > 0 {
		offset := (bookCondition.Page - 1) * bookCondition.Size
		db = db.Order("id").Limit(bookCondition.Size).Offset(offset)
	}

	err := db.Find(&books).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return books
}
