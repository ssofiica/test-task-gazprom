package user

import (
	"context"

	"github.com/ssofiica/test-task-gazprom/internal/entity"
	"github.com/ssofiica/test-task-gazprom/internal/repository/user"
)

type UseCase interface {
	GetAll(ctx context.Context) ([]*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Search(ctx context.Context, name string, surname string) (*entity.User, error)
}

type UseCaseLayer struct {
	repo user.Repo
}

func NewUseCaseLayer(repoProps user.Repo) UseCase {
	return &UseCaseLayer{
		repo: repoProps,
	}
}

func (uc *UseCaseLayer) GetAll(ctx context.Context) ([]*entity.User, error) {
	users, err := uc.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uc *UseCaseLayer) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	return nil, nil
}

func (uc *UseCaseLayer) Search(ctx context.Context, name string, surname string) (*entity.User, error) {
	return nil, nil
}
