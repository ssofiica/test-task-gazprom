package auth

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
	"github.com/ssofiica/test-task-gazprom/internal/entity"
)

type Repo interface {
	CreateUser(user *entity.User) error
	SetSessionValue(session *entity.Session) error
	GetSessionValue(sessionId string) error
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

func (repo *RepoLayer) CreateUser(user *entity.User) error {
	return nil
}

func (repo *RepoLayer) SetSessionValue(session *entity.Session) error {
	return nil
}

func (repo *RepoLayer) GetSessionValue(sessionId string) error {
	return nil
}
