package auth

import (
	"github.com/ssofiica/test-task-gazprom/internal/entity"
	auth "github.com/ssofiica/test-task-gazprom/internal/repository"
)

type UseCase interface {
	SignUp(user *entity.User, session *entity.Session) error
	SignIn(session *entity.Session) error
}

type UseCaseLayer struct {
	repo auth.Repo
}

func NewUseCaseLayer(repoProps auth.Repo) UseCase {
	return &UseCaseLayer{
		repo: repoProps,
	}
}

func (repo *UseCaseLayer) SignUp(user *entity.User, session *entity.Session) error {
	return nil
}

func (repo *UseCaseLayer) SignIn(session *entity.Session) error {
	return nil
}
