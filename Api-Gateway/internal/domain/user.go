package domain

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	Nick     string    `json:"nick"`
}
