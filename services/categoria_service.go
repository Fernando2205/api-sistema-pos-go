package services

import (
	"errors"
	"fmt"
	"sistema_pos_go/models"
	"sistema_pos_go/repositories"
	"sistema_pos_go/utils"
	"strings"

	"gorm.io/gorm"
)

// CategoriaService contiene la lógica de negocio para categorías
type CategoriaService struct {
	repo *repositories.CategoriaRepository
}

// NewCategoriaService crea una nueva instancia del servicio
func NewCategoriaService(repo *repositories.CategoriaRepository) *CategoriaService {
	return &CategoriaService{repo: repo}
}

// validateNombre valida y normaliza el nombre de una categoría
func (s *CategoriaService) validateNombre(nombre string) (string, error) {
	// Normalizar nombre
	nombre = strings.TrimSpace(nombre)

	// Validar que no esté vacío
	if nombre == "" {
		return "", fmt.Errorf("%w: el nombre de la categoría es obligatorio", utils.ErrEmptyField)
	}

	// Validar longitud máxima
	if len(nombre) > utils.MaxNombreCategoriaLength {
		return "", fmt.Errorf("%w: el nombre no puede exceder %d caracteres", utils.ErrMaxLengthExceeded, utils.MaxNombreCategoriaLength)
	}

	return nombre, nil
}

// validateNombreUniqueness verifica que no exista otra categoría con el mismo nombre
func (s *CategoriaService) validateNombreUniqueness(nombre string, excludeID int) error {
	categoriaExistente, err := s.repo.FindByNombre(nombre)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if categoriaExistente != nil && categoriaExistente.ID != excludeID {
		return fmt.Errorf("%w: ya existe una categoría con ese nombre", utils.ErrDuplicateEntry)
	}

	return nil
}

// GetAll retorna todas las categorías
func (s *CategoriaService) GetAll() ([]models.Categoria, error) {
	return s.repo.FindAll()
}

// GetById busca una categoría por su ID
func (s *CategoriaService) GetById(id int) (*models.Categoria, error) {
	if err := utils.ValidateID(id); err != nil {
		return nil, err
	}

	categoria, err := s.repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w: categoría con ID %d no encontrada", utils.ErrNotFound, id)
		}
		return nil, err
	}

	return categoria, nil
}

// Create crea una nueva categoría
func (s *CategoriaService) Create(categoria *models.Categoria) error {
	// Validar y normalizar nombre
	nombreValidado, err := s.validateNombre(categoria.Nombre)
	if err != nil {
		return err
	}
	categoria.Nombre = nombreValidado

	// Validar que no exista una categoría con el mismo nombre
	if err := s.validateNombreUniqueness(nombreValidado, 0); err != nil {
		return err
	}

	return s.repo.Create(categoria)
}

// Delete elimina una categoría por su ID con validaciones
func (s *CategoriaService) Delete(id int) error {
	if err := utils.ValidateID(id); err != nil {
		return err
	}

	categoria, err := s.repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("%w: categoría con ID %d no encontrada", utils.ErrNotFound, id)
		}
		return err
	}

	return s.repo.Delete(categoria)
}

// Update actualiza completamente una categoría
func (s *CategoriaService) Update(id int, categoria *models.Categoria) error {
	if err := utils.ValidateID(id); err != nil {
		return err
	}

	// Validar que existe
	existe, err := s.repo.ExistsById(id)
	if err != nil {
		return err
	}
	if !existe {
		return fmt.Errorf("%w: categoría con ID %d no encontrada", utils.ErrNotFound, id)
	}

	// Validar y normalizar nombre
	nombreValidado, err := s.validateNombre(categoria.Nombre)
	if err != nil {
		return err
	}
	categoria.Nombre = nombreValidado

	// Validar que no exista otra categoría con el mismo nombre
	if err := s.validateNombreUniqueness(nombreValidado, id); err != nil {
		return err
	}

	// Mantener el ID original
	categoria.ID = id
	return s.repo.Update(categoria)
}

// PartialUpdate actualiza campos específicos de una categoría
/*func (s *CategoriaService) PartialUpdate(id int, campos map[string]interface{}) (*models.Categoria, error) {
	if id <= 0 {
		return nil, utils.ErrInvalidID
	}

	if len(campos) == 0 {
		return nil, fmt.Errorf("%w: no se enviaron campos para actualizar", utils.ErrInvalidData)
	}

	// Validar que existe
	existe, err := s.repo.ExistsById(id)
	if err != nil {
		return nil, err
	}
	if !existe {
		return nil, fmt.Errorf("%w: categoría con ID %d no encontrada", utils.ErrNotFound, id)
	}

	// Si se está actualizando el nombre, validarlo
	if nombre, existe := campos["nombre"]; existe {
		nombreStr, ok := nombre.(string)
		if !ok {
			return nil, fmt.Errorf("%w: el nombre debe ser una cadena de texto", utils.ErrInvalidData)
		}

		// Validar y normalizar nombre
		nombreValidado, err := s.validateNombre(nombreStr)
		if err != nil {
			return nil, err
		}

		// Validar que no exista otra categoría con ese nombre
		if err := s.validateNombreUniqueness(nombreValidado, id); err != nil {
			return nil, err
		}

		campos["nombre"] = nombreValidado
	}

	if err := s.repo.PartialUpdate(id, campos); err != nil {
		return nil, err
	}

	// Obtener y retornar la categoría actualizada
	return s.repo.FindById(id)
}*/
