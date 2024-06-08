package auth

import (
	"context"

	"github.com/ssofiica/test-task-gazprom/internal/entity"
	"github.com/ssofiica/test-task-gazprom/internal/entity/dto"
	"github.com/ssofiica/test-task-gazprom/internal/repository/auth"
	"golang.org/x/crypto/bcrypt"
)

type UseCase interface {
	SignUp(ctx context.Context, user *dto.SignUp, session *entity.Session) error
	SignIn(ctx context.Context, session *entity.Session) error
	SetSessionValue(ctx context.Context, session *entity.Session) error
	GetSessionValue(ctx context.Context, sessionId string) (string, error)
}

type UseCaseLayer struct {
	repo auth.Repo
}

func NewUseCaseLayer(repoProps auth.Repo) UseCase {
	return &UseCaseLayer{
		repo: repoProps,
	}
}

func (uc *UseCaseLayer) SignUp(ctx context.Context, user *dto.SignUp, session *entity.Session) error {
	res, err := GetHash(user.Password)
	if err != nil {
		return err
	}
	user.Password = res
	err = uc.repo.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return uc.repo.SetSessionValue(ctx, session)
}

func (uc *UseCaseLayer) SignIn(ctx context.Context, session *entity.Session) error {
	return nil
}

func (uc *UseCaseLayer) SetSessionValue(ctx context.Context, session *entity.Session) error {
	return uc.repo.SetSessionValue(ctx, session)
}

func (uc *UseCaseLayer) GetSessionValue(ctx context.Context, sessionId string) (string, error) {
	email, err := uc.repo.GetSessionValue(ctx, sessionId)
	if err != nil {
		return "", err
	}
	return email, nil
}

func GetHash(input string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
