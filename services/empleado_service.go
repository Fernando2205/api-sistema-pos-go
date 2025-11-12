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

type EmpleadoService struct {
	repo *repositories.EmpleadoRepository
}

func NewEmpleadoService(repo *repositories.EmpleadoRepository) *EmpleadoService {
	return &EmpleadoService{repo: repo}
}

// validateEmailUniqueness verifica que no exista otro empleado con el mismo email
func (s *EmpleadoService) validateEmailUniqueness(email string, excludeID int) error {
	empleadoExistente, err := s.repo.FindByEmail(email)

	// Si no se encontró, no hay conflicto
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	// Si se encontró y no es el mismo registro que estamos actualizando
	if empleadoExistente.ID != excludeID {
		return fmt.Errorf("%w: ya existe un empleado con ese email", utils.ErrAlreadyExists)
	}

	return nil
}

// normalizeData normaliza los datos del empleado
func (s *EmpleadoService) normalizeData(empleado *models.Empleado) {
	empleado.Nombre = strings.TrimSpace(empleado.Nombre)
	empleado.Rol = strings.TrimSpace(empleado.Rol)
	empleado.Email = strings.TrimSpace(strings.ToLower(empleado.Email))
}

func (s *EmpleadoService) GetAll() ([]models.Empleado, error) {
	return s.repo.FindAll()
}

func (s *EmpleadoService) GetById(id int) (*models.Empleado, error) {
	if err := utils.ValidateID(id); err != nil {
		return nil, err
	}

	empleado, err := s.repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w: empleado con ID %d no encontrado", utils.ErrNotFound, id)
		}
		return nil, err
	}

	return empleado, nil
}

func (s *EmpleadoService) Create(empleado *models.Empleado) error {
	// Normalizar datos
	s.normalizeData(empleado)

	// Validar el email
	if err := s.validateEmailUniqueness(empleado.Email, 0); err != nil {
		return err
	}

	// Hashear contraseña antes de guardar
	hashedPassword, err := utils.HashPassword(empleado.Contraseña)
	if err != nil {
		return fmt.Errorf("error al procesar la contraseña: %w", err)
	}
	empleado.Contraseña = hashedPassword

	return s.repo.Create(empleado)
}

func (s *EmpleadoService) Delete(id int) error {
	if err := utils.ValidateID(id); err != nil {
		return err
	}

	empleado, err := s.repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("%w: empleado con ID %d no encontrado", utils.ErrNotFound, id)
		}
		return err
	}

	return s.repo.Delete(empleado)
}

func (s *EmpleadoService) Update(id int, empleado *models.Empleado) error {
	if err := utils.ValidateID(id); err != nil {
		return err
	}

	// Validar que existe
	existe, err := s.repo.ExistsById(id)
	if err != nil {
		return err
	}
	if !existe {
		return fmt.Errorf("%w: empleado con ID %d no encontrado", utils.ErrNotFound, id)
	}

	// Normalizar datos
	s.normalizeData(empleado)

	// Validar unicidad del email
	if err := s.validateEmailUniqueness(empleado.Email, id); err != nil {
		return err
	}

	// Hashear contraseña antes de guardar
	hashedPassword, err := utils.HashPassword(empleado.Contraseña)
	if err != nil {
		return fmt.Errorf("error al procesar la contraseña: %w", err)
	}
	empleado.Contraseña = hashedPassword

	// Mantener el ID original
	empleado.ID = id
	return s.repo.Update(empleado)
}

func (s *EmpleadoService) PartialUpdate(id int, request *models.EmpleadoPatch) (*models.Empleado, error) {
	if err := utils.ValidateID(id); err != nil {
		return nil, err
	}

	existe, err := s.repo.ExistsById(id)
	if err != nil {
		return nil, err
	}
	if !existe {
		return nil, fmt.Errorf("%w: empleado con ID %d no encontrado", utils.ErrNotFound, id)
	}

	// Validar que al menos un campo fue enviado
	if request.Nombre == nil && request.Rol == nil && request.Email == nil && request.Contraseña == nil {
		return nil, fmt.Errorf("%w: %s", utils.ErrInvalidData, utils.MsgNoFieldsToUpdate)
	}

	// Normalizar y validar email si fue enviado
	if request.Email != nil {
		email := strings.TrimSpace(strings.ToLower(*request.Email))
		request.Email = &email

		if err := s.validateEmailUniqueness(email, id); err != nil {
			return nil, err
		}
	}

	// Normalizar nombre si fue enviado
	if request.Nombre != nil {
		nombre := strings.TrimSpace(*request.Nombre)
		request.Nombre = &nombre
	}

	// Normalizar rol si fue enviado
	if request.Rol != nil {
		rol := strings.TrimSpace(*request.Rol)
		request.Rol = &rol
	}

	// Hashear contraseña si fue enviada
	if request.Contraseña != nil {
		hashedPassword, err := utils.HashPassword(*request.Contraseña)
		if err != nil {
			return nil, fmt.Errorf("error al procesar la contraseña: %w", err)
		}
		request.Contraseña = &hashedPassword
	}

	if err := s.repo.PartialUpdate(id, request); err != nil {
		return nil, err
	}

	// Retornar el empleado actualizado
	return s.repo.FindById(id)
}
