package repositories

import (
	"sistema_pos_go/models"

	"gorm.io/gorm"
)

type EmpleadoRepository struct {
	DB *gorm.DB
}

func NewEmpleadoRepository(db *gorm.DB) *EmpleadoRepository {
	return &EmpleadoRepository{DB: db}
}

func (r *EmpleadoRepository) FindAll() ([]models.Empleado, error) {
	var empleados []models.Empleado
	err := r.DB.Find(&empleados).Error
	return empleados, err
}

func (r *EmpleadoRepository) FindById(id int) (*models.Empleado, error) {
	var empleado models.Empleado
	err := r.DB.First(&empleado, id).Error
	return &empleado, err
}

func (r *EmpleadoRepository) FindByEmail(email string) (*models.Empleado, error) {
	var empleado models.Empleado
	err := r.DB.Where("email = ?", email).First(&empleado).Error
	return &empleado, err
}

func (r *EmpleadoRepository) ExistsById(id int) (bool, error) {
	var count int64
	err := r.DB.Model(&models.Empleado{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}

func (r *EmpleadoRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.DB.Model(&models.Empleado{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

func (r *EmpleadoRepository) Create(empleado *models.Empleado) error {
	return r.DB.Create(empleado).Error
}

func (r *EmpleadoRepository) Delete(empleado *models.Empleado) error {
	return r.DB.Delete(empleado).Error
}

func (r *EmpleadoRepository) Update(empleado *models.Empleado) error {
	return r.DB.Save(empleado).Error
}

func (r *EmpleadoRepository) PartialUpdate(id int, empleado *models.EmpleadoPatch) error {
	updates := make(map[string]interface{})

	if empleado.Nombre != nil {
		updates["nombre"] = *empleado.Nombre
	}
	if empleado.Rol != nil {
		updates["rol"] = *empleado.Rol
	}
	if empleado.Email != nil {
		updates["email"] = *empleado.Email
	}
	if empleado.Contraseña != nil {
		updates["contraseña"] = *empleado.Contraseña
	}

	return r.DB.Model(&models.Empleado{}).Where("id = ?", id).Updates(updates).Error
}
