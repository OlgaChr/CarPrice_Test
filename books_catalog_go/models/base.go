package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB //база данных

func init() {

	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.AutoMigrate(&Book{}, &Author{}) //Миграция базы данных

	db.Table("authors_books").AddForeignKey("author_id", "authors(id)", "RESTRICT", "RESTRICT")
	db.Table("authors_books").AddForeignKey("book_id", "books(id)", "RESTRICT", "RESTRICT")
}

// возвращает дескриптор объекта DB
func GetDB() *gorm.DB {
	return db
}
