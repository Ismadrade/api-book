package migrations

import (
	"github.com/Ismadrade/api-book/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
