package repositories

import (
	"sistema_pos_go/models"

	"gorm.io/gorm"
)

type CategoriaRepository struct {
	DB *gorm.DB
}

func NewCategoriaRepository(db *gorm.DB) *CategoriaRepository {
	return &CategoriaRepository{DB: db}
}

func (r *CategoriaRepository) FindAll() ([]models.Categoria, error) {
	var categorias []models.Categoria
	err := r.DB.Find(&categorias).Error
	return categorias, err
}
