package models

type Mesa struct {
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Numero    int  `gorm:"not null;unique" json:"numero"`
	Capacidad int  `gorm:"not null" json:"capacidad"`
}
