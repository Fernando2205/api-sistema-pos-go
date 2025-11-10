package models

type Mesa struct {
	ID        int `gorm:"primaryKey;autoIncrement" json:"id"`
	Numero    int `gorm:"not null;unique" json:"numero"`
	Capacidad int `gorm:"not null" json:"capacidad"`
}

func (Mesa) TableName() string {
	return "mesas"
}

type MesaPatch struct {
	Numero    *int `json:"numero"`
	Capacidad *int `json:"capacidad"`
}
