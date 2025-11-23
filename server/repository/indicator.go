package repository

import (
	"github.com/developerasun/SignalDash/server/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewIndicator(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindOne(name string) (*models.Indicator, error) {
	var i models.Indicator
	tx := r.db.First(&i, "name = ?", name)

	return &i, tx.Error
}
