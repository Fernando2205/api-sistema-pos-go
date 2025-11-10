package services

import (
	"errors"
	"fmt"
	"sistema_pos_go/models"
	"sistema_pos_go/repositories"
	"sistema_pos_go/utils"

	"gorm.io/gorm"
)

type MesaService struct {
	repo *repositories.MesaRepository
}

func NewMesaService(repo *repositories.MesaRepository) *MesaService {
	return &MesaService{repo: repo}
}
func (s *MesaService) ValidateNumeroUniqueness(numero int, excludeID int) error {
	mesa, err := s.repo.FindByNumero(numero)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No existe mesa con ese número, entonces está disponible
			return nil
		}
		// Error diferente, lo retornamos
		return err
	}
	// Si encontramos una mesa, verificamos si es la misma que estamos actualizando
	if mesa.ID != excludeID {
		return fmt.Errorf("%w: ya existe una mesa con ese número", utils.ErrAlreadyExists)
	}
	return nil
}
func (s *MesaService) GetAll() ([]models.Mesa, error) {
	return s.repo.FindAll()
}

func (s *MesaService) GetById(id int) (*models.Mesa, error) {
	if err := utils.ValidateID(id); err != nil {
		return nil, err
	}

	mesa, err := s.repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w: categoría con ID %d no encontrada", utils.ErrNotFound, id)
		}
		return nil, err
	}
	return mesa, nil
}

func (s *MesaService) Create(mesa *models.Mesa) error {
	//validar numero
	if mesa.Numero <= 0 {
		return fmt.Errorf("%w: el número de mesa debe ser mayor que cero", utils.ErrInvalidData)
	}
	if mesa.Capacidad <= 0 {
		return fmt.Errorf("%w: la capacidad de la mesa debe ser mayor que cero", utils.ErrInvalidData)
	}
	//Validar que no exista otra mesa con el mismo numero
	existe, err := s.repo.ExistByNumero(mesa.Numero)
	if err != nil {
		return err
	}
	if existe {
		return fmt.Errorf("%w: ya existe una mesa con el número %d", utils.ErrAlreadyExists, mesa.Numero)
	}
	if err := s.repo.Create(mesa); err != nil {
		return err
	}
	return nil
}

func (s *MesaService) Delete(id int) error {
	if err := utils.ValidateID(id); err != nil {
		return err
	}

	mesa, err := s.repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("%w: categoría con ID %d no encontrada", utils.ErrNotFound, id)
		}
		return err
	}
	return s.repo.Delete(mesa)
}

func (s *MesaService) Update(id int, mesa *models.Mesa) error {
	if err := utils.ValidateID(id); err != nil {
		return err
	}

	if mesa.Numero <= 0 {
		return fmt.Errorf("%w: el número de mesa debe ser mayor que cero", utils.ErrInvalidData)
	}
	if mesa.Capacidad <= 0 {
		return fmt.Errorf("%w: la capacidad de la mesa debe ser mayor que cero", utils.ErrInvalidData)
	}

	existe, err := s.repo.ExistsById(id)
	if err != nil {
		return err
	}
	if !existe {
		return fmt.Errorf("%w: categoría con ID %d no encontrada", utils.ErrNotFound, id)
	}
	if err := s.ValidateNumeroUniqueness(mesa.Numero, id); err != nil {
		return err
	}
	mesa.ID = id
	return s.repo.Update(mesa)
}

func (s *MesaService) PartialUpdate(id int, request *models.MesaPatch) (*models.Mesa, error) {
	if err := utils.ValidateID(id); err != nil {
		return nil, err
	}

	exist, err := s.repo.ExistsById(id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("%w: mesa con ID %d no encontrada", utils.ErrNotFound, id)
	}
	if request.Numero == nil && request.Capacidad == nil {
		return nil, fmt.Errorf("%w: %s", utils.ErrInvalidData, utils.MsgNoFieldsToUpdate)
	}

	//Validar número si en envió
	if request.Numero != nil {
		numero := *request.Numero
		if numero < utils.MinNumeroMesa {
			return nil, fmt.Errorf("%w: el número de mesa debe ser mayor a 0", utils.ErrInvalidData)
		}
		if err := s.ValidateNumeroUniqueness(numero, id); err != nil {
			return nil, err
		}

	}
	//Validar capacidad si en envió
	if request.Capacidad != nil {
		capacidad := *request.Capacidad
		if capacidad < utils.MinCapacidadMesa {
			return nil, fmt.Errorf("%w: la capacidad de la mesa debe ser mayor a 0", utils.ErrInvalidData)
		}
	}

	if err := s.repo.PartialUpdate(id, request); err != nil {
		return nil, err
	}

	// Retornar la mesa actualizada
	return s.repo.FindById(id)

}
