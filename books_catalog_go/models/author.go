package models

import (
	"books_catalog_go/types"
	u "books_catalog_go/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Author struct {
	gorm.Model
	Surname    string
	Name       string
	Middlename string
	BirthYear  uint
	DeathYear  uint
	Books      []*Book `json:"-"`//`gorm:"many2many:authors_books;"`
	BooksCount uint
}

func (author *Author) Create() map[string]interface{} {
	GetDB().Create(author)

	resp := u.Message(true, "success")
	resp["author"] = author
	return resp
}

func (author *Author) Update() map[string]interface{} {
	GetDB().Model(&author).Update(author)

	resp := u.Message(true, "success")
	resp["author"] = author
	return resp
}

func DeleteAuthor(id uint) error {
	author := GetAuthor(id)
	return GetDB().Select("Books").Delete(&author).Error
}

func GetAuthor(id uint) *Author {
	author := &Author{}
	err := GetDB().Table("authors").Preload("Books").Where("id = ?", id).First(author).Error
	if err != nil {
		return nil
	}
	author.setBooksCount()
	return author
}

func GetAuthors() []*Author {
	authors := make([]*Author, 0)
	err := GetDB().Table("authors").Preload("Books").Find(&authors).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, author := range authors {
		author.setBooksCount()
	}

	return authors
}

func GetAuthorsWithSearch(authorCondition *types.AuthorCondition) []*Author {
	authors := make([]*Author, 0)
	db := GetDB().Table("authors").
		//Select("authors.*, count(distinct books.id) as books_count").
		Joins("left join books on authors.id = books.author_id").
		Preload("Books").
		Group("authors.id")

	if authorCondition.BookID > 0 {
		db = db.Where("books.id = ?", authorCondition.BookID)
	}
	if authorCondition.Page > 0 && authorCondition.Size > 0 {
		offset := (authorCondition.Page - 1) * authorCondition.Size
		db = db.Order("authors.id").Limit(authorCondition.Size).Offset(offset)
	}

	err := db.Find(&authors).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, author := range authors {
		author.setBooksCount()
	}

	return authors
}

func (author *Author) setBooksCount() {
	author.BooksCount = uint(len(author.Books))
}
