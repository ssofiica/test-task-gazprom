package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ssofiica/test-task-gazprom/internal/entity"
	"github.com/ssofiica/test-task-gazprom/internal/entity/dto"
	"github.com/ssofiica/test-task-gazprom/pkg/myerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("db error", func(t *testing.T) {
		s := setUp(t)
		defer s.db.Close()
		user := &dto.SignUp{
			Name:     "София",
			Surname:  "Валова",
			Email:    "s@mail.ru",
			Password: "$2a$14$kf53kHgdL4tvfiNWklMSvuICX6qFOwMmdruCoreR9uqhEtA4pnk5a",
			Birthday: "2003-04-26",
		}

		s.dbMock.ExpectQuery(`INSERT INTO "user" \(name, surname, email, password, birthday\) VALUES \(\$1, \$2, \$3, \$4, \$5\) RETURNING id, name, surname, email`).
			WithArgs(user.Name, user.Surname, user.Email, user.Password, user.Birthday).
			WillReturnError(fmt.Errorf("db_error"))

		_, err := s.repo.CreateUser(ctx, user)
		require.NoError(t, s.dbMock.ExpectationsWereMet())
		assert.Error(t, err)
	})

	t.Run("noRows", func(t *testing.T) {
		s := setUp(t)
		defer s.db.Close()
		user := &dto.SignUp{
			Name:     "София",
			Surname:  "Валова",
			Email:    "s@mail.ru",
			Password: "$2a$14$kf53kHgdL4tvfiNWklMSvuICX6qFOwMmdruCoreR9uqhEtA4pnk5a",
			Birthday: "2003-04-26",
		}

		s.dbMock.ExpectQuery(`INSERT INTO "user" \(name, surname, email, password, birthday\) VALUES \(\$1, \$2, \$3, \$4, \$5\) RETURNING id, name, surname, email`).
			WithArgs(user.Name, user.Surname, user.Email, user.Password, user.Birthday).
			WillReturnError(sql.ErrNoRows)

		_, err := s.repo.CreateUser(ctx, user)
		require.NoError(t, s.dbMock.ExpectationsWereMet())
		assert.Equal(t, myerrors.NoCreatingUser, err)
	})

	t.Run("ok", func(t *testing.T) {
		s := setUp(t)
		defer s.db.Close()
		user := &dto.SignUp{
			Name:     "София",
			Surname:  "Валова",
			Email:    "s@mail.ru",
			Password: "$2a$14$kf53kHgdL4tvfiNWklMSvuICX6qFOwMmdruCoreR9uqhEtA4pnk5a",
			Birthday: "2003-04-26",
		}

		rows := sqlmock.NewRows([]string{"id", "name", "surname", "email"}).AddRow(1, "София", "Валова", "s@mail.ru")

		s.dbMock.ExpectQuery(`INSERT INTO "user" \(name, surname, email, password, birthday\) VALUES \(\$1, \$2, \$3, \$4, \$5\) RETURNING id, name, surname, email`).
			WithArgs(user.Name, user.Surname, user.Email, user.Password, user.Birthday).
			WillReturnRows(rows)

		u, err := s.repo.CreateUser(ctx, user)
		require.NoError(t, s.dbMock.ExpectationsWereMet())
		assert.NoError(t, err)
		assert.Equal(t, &entity.User{
			Id:      1,
			Name:    "София",
			Surname: "Валова",
			Email:   "s@mail.ru",
		}, u)
	})
}

func TestSetSessionValue(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("redis error", func(t *testing.T) {
		s := setUp(t)
		s.db.Close()
		session := &entity.Session{
			Id:    "session_id",
			Email: "test@example.com",
		}

		myerr := errors.New("error")
		s.rMock.ExpectSet("session_id", "test@example.com", 14*24*time.Hour).SetErr(myerr)
		err := s.repo.SetSessionValue(context.Background(), session)
		require.NoError(t, s.rMock.ExpectationsWereMet())
		assert.Equal(t, myerr, err)
	})

	t.Run("ok", func(t *testing.T) {
		s := setUp(t)
		s.db.Close()
		session := &entity.Session{
			Id:    "session_id",
			Email: "test@example.com",
		}

		s.rMock.ExpectSet("session_id", "test@example.com", 14*24*time.Hour).SetVal("OK")
		err := s.repo.SetSessionValue(ctx, session)
		require.NoError(t, s.rMock.ExpectationsWereMet())
		assert.NoError(t, err)
	})
}

func TestGetSessionValue(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("redis error", func(t *testing.T) {
		s := setUp(t)
		s.db.Close()
		key := "session_id"

		myerr := errors.New("error")
		s.rMock.ExpectGet(key).SetErr(myerr)
		_, err := s.repo.GetSessionValue(context.Background(), key)
		require.NoError(t, s.rMock.ExpectationsWereMet())
		assert.Equal(t, myerr, err)
	})

	t.Run("ok", func(t *testing.T) {
		s := setUp(t)
		s.db.Close()

		expect := "s@mail.ru"
		key := "session_id"

		s.rMock.ExpectGet(key).SetVal(expect)
		session, err := s.repo.GetSessionValue(ctx, key)
		require.NoError(t, s.rMock.ExpectationsWereMet())
		assert.NoError(t, err)
		assert.Equal(t, session, expect)
	})
}

func TestDeleteSessionValue(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("ok", func(t *testing.T) {
		s := setUp(t)
		s.db.Close()

		key := "session_id"
		s.rMock.ExpectDel(key).SetVal(1)
		err := s.repo.DeleteSessionValue(ctx, key)
		assert.NoError(t, err)
	})
}
