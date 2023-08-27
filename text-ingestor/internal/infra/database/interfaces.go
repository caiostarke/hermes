package database

import "github.com/caiostarke/hermes/text-ingestor/internal/entity"

type TextInterface interface {
	Create(text *entity.Text) error
	FindAll() ([]entity.Text, error)
}


