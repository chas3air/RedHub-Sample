package implementation

import "api-gateway/internal/domain"

type UsersManager struct{}

func (um *UsersManager) ListUsers() ([]domain.User, error)
func (um *UsersManager) GetUser(id int) (domain.User, error)
func (um *UsersManager) Insert(user domain.User) error
func (um *UsersManager) Update(id int, user domain.User) error
func (um *UsersManager) Delete(id int) (domain.User, error)
