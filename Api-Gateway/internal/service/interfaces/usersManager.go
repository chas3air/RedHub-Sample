package interfaces

import "api-gateway/internal/domain"

type IUsersManager interface {
	ListUsers() ([]domain.User, error)
	GetUser(id int) (domain.User, error)
	Insert(user domain.User) error
	Update(id int, user domain.User) error
	Delete(id int) (domain.User, error)
}
