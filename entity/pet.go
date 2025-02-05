package entity

import "time"

type Pet struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Species   string    `json:"species"`
	Age       int       `json:"age"`
	UserId    uint      `json:"user_id"` // Foreign key from User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
