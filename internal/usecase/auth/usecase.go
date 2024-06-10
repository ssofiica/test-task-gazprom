package auth

import (
	"context"

	"github.com/ssofiica/test-task-gazprom/internal/entity"
	"github.com/ssofiica/test-task-gazprom/internal/entity/dto"
	"github.com/ssofiica/test-task-gazprom/internal/repository/auth"
	"github.com/ssofiica/test-task-gazprom/pkg/myerrors"
	"golang.org/x/crypto/bcrypt"
)

type UseCase interface {
	SignUp(ctx context.Context, user *dto.SignUp, session *entity.Session) (*entity.User, error)
	SignIn(ctx context.Context, user *entity.User, signInInfo *dto.SignIn, session *entity.Session) error
	SetSessionValue(ctx context.Context, session *entity.Session) error
	GetSessionValue(ctx context.Context, sessionId string) (string, error)
	DeleteSession(ctx context.Context, sessionId string) error
}

type UseCaseLayer struct {
	repo auth.Repo
}

func NewUseCaseLayer(repoProps auth.Repo) UseCase {
	return &UseCaseLayer{
		repo: repoProps,
	}
}

func (uc *UseCaseLayer) SignUp(ctx context.Context, user *dto.SignUp, session *entity.Session) (*entity.User, error) {
	res, err := GetHash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = res
	u, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	err = uc.repo.SetSessionValue(ctx, session)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (uc *UseCaseLayer) SignIn(ctx context.Context, user *entity.User, sigInInfo *dto.SignIn, session *entity.Session) error {
	// чекнуть почту и пароль с тем, что в базе
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(sigInInfo.Password)); err == nil {
		// если норм, то установить в редис сессию
		err = uc.SetSessionValue(ctx, session)
		if err != nil {
			return err
		}
		return nil
	}
	return myerrors.WrongPassword
}

func (uc *UseCaseLayer) DeleteSession(ctx context.Context, sessionId string) error {
	return uc.repo.DeleteSessionValue(ctx, sessionId)
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
