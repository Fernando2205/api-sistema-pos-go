package models

type Categoria struct {
	ID     int    `gorm:"column:id;primaryKey" json:"id"`
	Nombre string `gorm:"column:nombre;type:varchar(100);not null" json:"nombre" binding:"required,max=100"`
}

func (Categoria) TableName() string {
	return "categorias"
}
