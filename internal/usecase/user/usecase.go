package user

import (
	"context"

	"github.com/ssofiica/test-task-gazprom/internal/entity"
	"github.com/ssofiica/test-task-gazprom/internal/repository/user"
)

type UseCase interface {
	GetAll(ctx context.Context) ([]*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Search(ctx context.Context, name string, surname string) ([]*entity.User, error)
	Subscribe(ctx context.Context, birthdayUserId uint64, subscribingUserId uint64) error
	UnSubscribe(ctx context.Context, birthdayUserId uint64, subscribingUserId uint64) error
	GetTodayBirthdayUsers(ctx context.Context, userId uint64) ([]*entity.User, error)
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
	return uc.repo.GetByEmail(ctx, email)
}

func (uc *UseCaseLayer) Search(ctx context.Context, name string, surname string) ([]*entity.User, error) {
	return uc.repo.Search(ctx, name, surname)
}

func (uc *UseCaseLayer) Subscribe(ctx context.Context, birthdayUserId uint64, subscribingUserId uint64) error {
	_, err := uc.repo.GetById(ctx, birthdayUserId)
	if err != nil {
		return err
	}
	return uc.repo.Subscribe(ctx, birthdayUserId, subscribingUserId)
}

func (uc *UseCaseLayer) UnSubscribe(ctx context.Context, birthdayUserId uint64, subscribingUserId uint64) error {
	_, err := uc.repo.GetById(ctx, birthdayUserId)
	if err != nil {
		return err
	}
	return uc.repo.UnSubscribe(ctx, birthdayUserId, subscribingUserId)
}

func (uc *UseCaseLayer) GetTodayBirthdayUsers(ctx context.Context, userId uint64) ([]*entity.User, error) {
	users, err := uc.repo.GetTodayBirthdayUsers(ctx, userId)
	if err != nil {
		return nil, err
	}
	return users, nil
}
