package services

import (
	"sistema_pos_go/models"
	"sistema_pos_go/repositories"
)

type CategoriaService struct {
	repo *repositories.CategoriaRepository
}

func NewCategoriaService(repo *repositories.CategoriaRepository) *CategoriaService {
	return &CategoriaService{repo: repo}
}

func (s *CategoriaService) GetAll() ([]models.Categoria, error) {
	return s.repo.FindAll()
}
