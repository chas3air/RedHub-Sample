package usermanager

import (
	"context"
	"fmt"
	"log/slog"
	"usersManageService/internal/domain/interfaces/storage"
	"usersManageService/internal/domain/models"

	"github.com/google/uuid"
)

// TODO: разделить интерфейсы Storage и UserManager
type UserManager struct {
	log     *slog.Logger
	storage storage.Storage
}

func New(log *slog.Logger, storage storage.Storage) *UserManager {
	return &UserManager{
		log:     log,
		storage: storage,
	}
}

func (um *UserManager) GetUsers(ctx context.Context) ([]models.User, error) {
	const op = "services.userManager.ListUsers"
	users, err := um.storage.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return users, nil
}

func (um *UserManager) GetUserById(ctx context.Context, uid uuid.UUID) (models.User, error) {
	const op = "services.userManager.GetUserById"

	user, err := um.storage.GetUserById(ctx, uid)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (um *UserManager) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	const op = "services.userManager.GetUserByEmail"

	user, err := um.storage.GetUserByEmail(ctx, email)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (um *UserManager) Insert(ctx context.Context, user models.User) error {
	const op = "services.userManager.Insert"

	err := um.storage.Insert(ctx, user)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (um *UserManager) Update(ctx context.Context, uid uuid.UUID, user models.User) error {
	const op = "services.userManager.Update"

	err := um.storage.Update(ctx, uid, user)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (um *UserManager) Delete(ctx context.Context, uid uuid.UUID) (models.User, error) {
	const op = "services.userManager.Delete"

	user, err := um.storage.Delete(ctx, uid)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
