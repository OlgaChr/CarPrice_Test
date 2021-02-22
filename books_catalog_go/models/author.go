package models

import (
	"books_catalog_go/dto"
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
	Books      []*Book `json:"-" gorm:"many2many:authors_books;foreignKey:ID;joinForeignKey:AuthorID;References:ID;JoinReferences:BookID"`
	BooksCount uint
}

func (author *Author) Validate() (map[string]interface{}, bool) {

	if author.Name == "" {
		return u.Message(false, "Author name should be on the payload"), false
	}

	if author.Surname == "" {
		return u.Message(false, "Author surname should be on the payload"), false
	}

	if author.BirthYear == 0 {
		return u.Message(false, "Birth Year should be on the payload"), false
	}

	//Все обязательные параметры присутствуют
	return u.Message(true, "success"), true
}

func (author *Author) Create() map[string]interface{} {
	if resp, ok := author.Validate(); !ok {
		return resp
	}

	GetDB().Create(author)

	resp := u.Message(true, "success")
	resp["author"] = author
	return resp
}

func (author *Author) Update() map[string]interface{} {
	if resp, ok := author.Validate(); !ok {
		return resp
	}

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
	return author
}

func GetAuthors() []*Author {
	authors := make([]*Author, 0)
	err := GetDB().Preload("Books").Find(&authors).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return authors
}

func GetAuthorsWithSearch(authorCondition *dto.AuthorCondition) []*Author {
	authors := make([]*Author, 0)
	db := GetDB().
		Joins("INNER JOIN authors_books ON authors_books.author_id = authors.id").
		Joins("LEFT JOIN books ON authors_books.book_id = books.id").
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

	return authors
}

func (author *Author) setBooksCount() {
	author.BooksCount = uint(len(author.Books))
}

func (author *Author) AfterFind(tx *gorm.DB) (err error) {
	author.setBooksCount()
	return
}
