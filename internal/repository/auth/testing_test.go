package auth

import (
	"database/sql"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type testFixtures struct {
	repo   Repo
	db     *sql.DB
	dbMock sqlmock.Sqlmock
	redis  *redis.Client
	rMock  redismock.ClientMock
}

func setUp(t *testing.T) testFixtures {
	db, dbMock, err := sqlmock.New()
	redis, rMock := redismock.NewClientMock()
	require.NoError(t, err)
	repo := NewRepoLayer(db, redis)
	return testFixtures{
		repo:   repo,
		db:     db,
		dbMock: dbMock,
		redis:  redis,
		rMock:  rMock,
	}
}
