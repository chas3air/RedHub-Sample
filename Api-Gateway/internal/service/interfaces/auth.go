package interfaces

import "github.com/google/uuid"

type IAuth interface {
	Login(email, password string) (token string, err error)
	Register(email, password string) (uid uuid.UUID)
	IsAdmin(uid uuid.UUID) (isAdmin bool, err error)
}
