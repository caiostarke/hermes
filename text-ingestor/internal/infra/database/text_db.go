package database

import (
	"github.com/caiostarke/hermes/text-ingestor/internal/entity"
	"gorm.io/gorm"
)

type Text struct {
	DB *gorm.DB
}

func NewText(db *gorm.DB) *Text {
	return &Text{DB: db}
}

func (t *Text) Create(text *entity.Text) error {
	return t.DB.Create(text).Error
}

func (t *Text) FindAll() ([]entity.Text, error) {
	var texts []entity.Text
	r := t.DB.Find(&texts)
	if r.Error != nil {
		return nil, r.Error
	}

	return texts, nil
}
