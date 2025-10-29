package models

type empleado struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Nombre     string `gorm:"type:varchar(100);not null" json:"nombre"`
	Rol        string `gorm:"type:varchar(100);not null" json:"rol"`
	Email      string `gorm:"type:varchar(100);not null;unique" json:"email"`
	Contraseña string `gorm:"type:varchar(100);not null" json:"contraseña"`
}
