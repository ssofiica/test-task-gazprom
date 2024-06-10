package auth

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ssofiica/test-task-gazprom/internal/entity"
	"github.com/ssofiica/test-task-gazprom/internal/entity/dto"
	"github.com/ssofiica/test-task-gazprom/pkg/myerrors"
)

type Repo interface {
	CreateUser(ctx context.Context, user *dto.SignUp) (*entity.User, error)
	SetSessionValue(ctx context.Context, session *entity.Session) error
	GetSessionValue(ctx context.Context, sessionId string) (string, error)
	DeleteSessionValue(ctx context.Context, sessionId string) error
}

type RepoLayer struct {
	db    *sql.DB
	redis *redis.Client
}

func NewRepoLayer(dbProps *sql.DB, redisProps *redis.Client) Repo {
	return &RepoLayer{
		db:    dbProps,
		redis: redisProps,
	}
}

func (repo *RepoLayer) CreateUser(ctx context.Context, user *dto.SignUp) (*entity.User, error) {
	row := repo.db.QueryRowContext(ctx, `INSERT INTO "user" (name, surname, email, password, birthday) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, surname, email`, user.Name, user.Surname, user.Email, user.Password, user.Birthday)
	u := entity.User{}
	err := row.Scan(&u.Id, &u.Name, &u.Surname, &u.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, myerrors.NoCreatingUser
		}
		return nil, err
	}
	return &u, nil
}

func (repo *RepoLayer) SetSessionValue(ctx context.Context, session *entity.Session) error {
	err := repo.redis.Set(ctx, session.Id, session.Email, 14*24*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func (repo *RepoLayer) GetSessionValue(ctx context.Context, sessionId string) (string, error) {
	value, err := repo.redis.Get(ctx, sessionId).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (repo *RepoLayer) DeleteSessionValue(ctx context.Context, sessionId string) error {
	return repo.redis.Del(ctx, sessionId).Err()
}
