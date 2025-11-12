package models

type Empleado struct {
	ID         int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Nombre     string `gorm:"type:varchar(100);not null" json:"nombre" binding:"required,max=100"`
	Rol        string `gorm:"type:varchar(100);not null" json:"rol" binding:"required,max=100"`
	Email      string `gorm:"type:varchar(100);not null;unique" json:"email" binding:"required,email,max=100"`
	Contraseña string `gorm:"type:varchar(100);not null" json:"contraseña" binding:"required,min=6,max=100"`
}

func (Empleado) TableName() string {
	return "empleados"
}

type EmpleadoPatch struct {
	Nombre     *string `json:"nombre" binding:"omitempty,max=100"`
	Rol        *string `json:"rol" binding:"omitempty,max=100"`
	Email      *string `json:"email" binding:"omitempty,email,max=100"`
	Contraseña *string `json:"contraseña" binding:"omitempty,min=6,max=100"`
}
