package repositories

import (
	"sistema_pos_go/models"

	"gorm.io/gorm"
)

// CategoriaRepository maneja las operaciones de base de datos para categorías
type CategoriaRepository struct {
	DB *gorm.DB
}

// NewCategoriaRepository crea una nueva instancia del repositorio
func NewCategoriaRepository(db *gorm.DB) *CategoriaRepository {
	return &CategoriaRepository{DB: db}
}

// FindByNombre busca una categoría por su nombre (case-insensitive)
func (r *CategoriaRepository) FindByNombre(nombre string) (*models.Categoria, error) {
	var categoria models.Categoria
	err := r.DB.Where("LOWER(nombre) = LOWER(?)", nombre).First(&categoria).Error
	return &categoria, err
}

/*Un slice internamente es una referencia a un array
por lo que no es necesario retornar un puntero*/

// FindAll retorna todas las categorías
func (r *CategoriaRepository) FindAll() ([]models.Categoria, error) {
	var categorias []models.Categoria
	err := r.DB.Find(&categorias).Error
	return categorias, err
}

/*Retornar un puntero permite que este pueda ser un valor nulo,
si retornamos un valor se retornaría un struct vacío en caso de no encontrar el registro*/

// FindById busca una categoría por su ID
func (r *CategoriaRepository) FindById(id int) (*models.Categoria, error) {
	var categoria models.Categoria
	err := r.DB.First(&categoria, id).Error
	return &categoria, err
}

// ExistsById verifica si existe una categoría con el ID dado
func (r *CategoriaRepository) ExistsById(id int) (bool, error) {
	var count int64
	err := r.DB.Model(&models.Categoria{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}

// Create inserta una nueva categoría en la base de datos
func (r *CategoriaRepository) Create(categoria *models.Categoria) error {
	return r.DB.Create(categoria).Error
}

// Delete elimina una categoría de la base de datos
func (r *CategoriaRepository) Delete(categoria *models.Categoria) error {
	return r.DB.Delete(categoria).Error
}

// Update actualiza completamente una categoría
func (r *CategoriaRepository) Update(categoria *models.Categoria) error {
	return r.DB.Save(categoria).Error // Save reemplaza el registro completo
}

// PartialUpdate actualiza campos específicos de una categoría
/*func (r *CategoriaRepository) PartialUpdate(id int, campos map[string]interface{}) error {
	return r.DB.Model(&models.Categoria{}).Where("id = ?", id).Updates(campos).Error // Updates solo actualiza campos no-cero o especificados
}*/
