package db

import (
	"github.com/MarkyMan4/gin-tutorial/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type BookDb struct {
	Db *gorm.DB
}

func New(dbName string) *BookDb {
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Book{})

	return &BookDb{db}
}
