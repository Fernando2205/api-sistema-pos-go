package models

type Categoria struct {
	ID     int    `gorm:"column:id;primaryKey" json:"id"`
	Nombre string `gorm:"column:nombre;type:varchar(100);not null" json:"nombre"`
}

func (Categoria) TableName() string {
	return "categorias"
}
