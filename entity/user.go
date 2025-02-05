package entity

import "time"

type User struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Puedes agregar una relación: Un usuario puede tener muchas mascotas
	Pet []Pet `json:"pets,omitempty"`
}
