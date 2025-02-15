package mock

import (
	"context"
	"fmt"
	"usersManageService/internal/domain/models"

	"github.com/google/uuid"
)

type MockStorage struct {
	users []models.User
}

func New() *MockStorage {
	return &MockStorage{
		users: make([]models.User, 0),
	}
}

// GetUsers implements storage.Storage.
func (m *MockStorage) GetUsers(ctx context.Context) ([]models.User, error) {
	return m.users, nil
}

// GetUserById implements storage.Storage.
func (m *MockStorage) GetUserById(ctx context.Context, id uuid.UUID) (models.User, error) {
	for _, v := range m.users {
		if v.Id == id {
			return v, nil
		}
	}

	return models.User{}, fmt.Errorf("%s: %s", "not found", "not found")
}

// GetUserByEmail implements storage.Storage.
func (m *MockStorage) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	for _, v := range m.users {
		if v.Email == email {
			return v, nil
		}
	}

	return models.User{}, fmt.Errorf("%s: %s", "not found", "not found")
}

// Insert implements storage.Storage.
func (m *MockStorage) Insert(ctx context.Context, user models.User) error {
	m.users = append(m.users, user)
	return nil
}

// Update implements storage.Storage.
func (m *MockStorage) Update(ctx context.Context, id uuid.UUID, user models.User) error {
	for i, v := range m.users {
		if v.Id == id {
			m.users[i] = user
			return nil
		}
	}
	return fmt.Errorf("not found")
}

// Delete implements storage.Storage.
func (m *MockStorage) Delete(context.Context, uuid.UUID) (models.User, error) {
	panic("unimplemented")
}
