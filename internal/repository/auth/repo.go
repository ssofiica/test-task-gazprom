package auth

import (
	"context"
	"database/sql"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ssofiica/test-task-gazprom/internal/entity"
	"github.com/ssofiica/test-task-gazprom/internal/entity/dto"
	"github.com/ssofiica/test-task-gazprom/pkg/myerrors"
)

type Repo interface {
	CreateUser(ctx context.Context, user *dto.SignUp) error
	SetSessionValue(ctx context.Context, session *entity.Session) error
	GetSessionValue(ctx context.Context, sessionId string) (string, error)
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

func (repo *RepoLayer) CreateUser(ctx context.Context, user *dto.SignUp) error {
	res, err := repo.db.ExecContext(ctx, `INSERT INTO "user" (name, surname, email, password, birthday) VALUES ($1, $2, $3, $4, $5)`, user.Name, user.Surname, user.Email, user.Password, user.Birthday)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return myerrors.NoCreatingUser
	}
	return nil
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

//err := repo.redis.Del(ctx, key).Err()
