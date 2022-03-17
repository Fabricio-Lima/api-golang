package migrations

import (
	"github.com/Fabricio-Lima/api-golang/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Game{})
}
