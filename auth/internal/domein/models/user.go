package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Login    string    `json:"login" gorm:"unique;not null"`
	Password string    `json:"password" gorm:"not null"`
	Role     string    `json:"role" gorm:"not null"`
	Nick     string    `json:"nick" gorm:"not null"`
}
