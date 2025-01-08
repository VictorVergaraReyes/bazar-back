package usecases

import (
	"bazar-back/internal/domain/entities"
)

type UserUseCase interface {
	GetAll() ([]*entities.Client, error)
	GetByID(id string) (*entities.Client, error)
	CreateNew(user *entities.Client) error
	DeleteByID(id string) (*entities.Client, error)
}
