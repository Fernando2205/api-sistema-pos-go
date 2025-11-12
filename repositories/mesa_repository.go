package repositories

import (
	"sistema_pos_go/models"

	"gorm.io/gorm"
)

type MesaRepository struct {
	DB *gorm.DB
}

func NewMesaRepository(db *gorm.DB) *MesaRepository {
	return &MesaRepository{DB: db}
}

func (r *MesaRepository) FindAll() ([]models.Mesa, error) {
	var mesas []models.Mesa
	err := r.DB.Find(&mesas).Error
	return mesas, err
}

func (r *MesaRepository) FindById(id int) (*models.Mesa, error) {
	var mesa models.Mesa
	err := r.DB.First(&mesa, id).Error
	return &mesa, err
}

func (r *MesaRepository) FindByNumero(numero int) (*models.Mesa, error) {
	var mesa models.Mesa
	err := r.DB.Where("numero = ?", numero).First(&mesa).Error
	return &mesa, err
}

func (r *MesaRepository) ExistsById(id int) (bool, error) {
	var count int64
	err := r.DB.Model(&models.Mesa{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}

func (r *MesaRepository) ExistsByNumero(numero int) (bool, error) {
	var count int64
	err := r.DB.Model(&models.Mesa{}).Where("numero = ?", numero).Count(&count).Error
	return count > 0, err
}

func (r *MesaRepository) Create(mesa *models.Mesa) error {
	return r.DB.Create(mesa).Error
}

func (r *MesaRepository) Delete(mesa *models.Mesa) error {
	return r.DB.Delete(mesa).Error
}

func (r *MesaRepository) Update(mesa *models.Mesa) error {
	return r.DB.Save(mesa).Error
}

func (r *MesaRepository) PartialUpdate(id int, mesa *models.MesaPatch) error {
	updates := make(map[string]interface{})

	if mesa.Numero != nil {
		updates["numero"] = *mesa.Numero
	}
	if mesa.Capacidad != nil {
		updates["capacidad"] = *mesa.Capacidad
	}

	return r.DB.Model(&models.Mesa{}).Where("id = ?", id).Updates(updates).Error
}
