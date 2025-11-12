package models

type Mesa struct {
	ID        int `gorm:"primaryKey;autoIncrement" json:"id"`
	Numero    int `gorm:"not null;unique" json:"numero" binding:"required,min=1"`
	Capacidad int `gorm:"not null" json:"capacidad" binding:"required,min=1"`
}

func (Mesa) TableName() string {
	return "mesas"
}

type MesaPatch struct {
	Numero    *int `json:"numero" binding:"omitempty,min=1"`
	Capacidad *int `json:"capacidad" binding:"omitempty,min=1"`
}
