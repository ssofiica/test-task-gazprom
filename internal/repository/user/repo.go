package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ssofiica/test-task-gazprom/internal/entity"
)

type Repo interface {
	GetAll(ctx context.Context) ([]*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetById(ctx context.Context, id uint64) (*entity.User, error)
	Search(ctx context.Context, name string, surname string) ([]*entity.User, error)
	Subscribe(ctx context.Context, birthdayUserId uint64, subscribingUserId uint64) error
	UnSubscribe(ctx context.Context, birthdayUserId uint64, subscribingUserId uint64) error
	GetTodayBirthdayUsers(ctx context.Context, userId uint64) ([]*entity.User, error)
}

type RepoLayer struct {
	db *sql.DB
}

func NewRepoLayer(dbProps *sql.DB) Repo {
	return &RepoLayer{
		db: dbProps,
	}
}

func (repo *RepoLayer) GetAll(ctx context.Context) ([]*entity.User, error) {
	rows, err := repo.db.QueryContext(ctx,
		`SELECT id, name, surname, email, birthday FROM "user"`)
	if err != nil {
		return nil, err
	}
	users := []*entity.User{}
	for rows.Next() {
		user := entity.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Surname,
			&user.Email, &user.Birthday)
		if err != nil {
			return nil, err
		}
		fmt.Println(user.Birthday)
		users = append(users, &user)
	}
	return users, nil
}

func (repo *RepoLayer) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := entity.User{}
	err := repo.db.QueryRowContext(ctx, `SELECT id, name, surname, birthday FROM "user" WHERE email=$1`, email).Scan(&user.Id, &user.Name, &user.Surname, &user.Birthday)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &entity.User{}, nil
		}
	}
	user.Email = email
	return &user, nil
}

func (repo *RepoLayer) GetById(ctx context.Context, id uint64) (*entity.User, error) {
	user := entity.User{}
	err := repo.db.QueryRowContext(ctx, `SELECT name, surname, email, birthday FROM "user" WHERE id=$1`, id).Scan(&user.Name, &user.Surname, &user.Email, &user.Birthday)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &entity.User{}, nil
		}
	}
	user.Id = id
	return &user, nil
}

func (repo *RepoLayer) Search(ctx context.Context, name string, surname string) ([]*entity.User, error) {
	rows, err := repo.db.QueryContext(ctx, `SELECT id, email, birthday FROM "user" WHERE name=$1 AND surname=$2`, name, surname)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []*entity.User{}, nil
		}
		return nil, err
	}
	users := []*entity.User{}
	for rows.Next() {
		user := entity.User{}
		err := rows.Scan(&user.Id, &user.Email, &user.Birthday)
		if err != nil {
			return nil, err
		}
		user.Name = name
		user.Surname = surname
		users = append(users, &user)
	}
	return users, nil
}

func (repo *RepoLayer) Subscribe(ctx context.Context, birthdayUserId uint64, subscribingUserId uint64) error {
	_, err := repo.db.ExecContext(ctx, `INSERT INTO birthday_subscribing (birthday_user_id, subscribing_user_id) 
		VALUES ($1, $2)`, birthdayUserId, subscribingUserId)
	if err != nil {
		return err
	}
	return nil
}

func (repo *RepoLayer) UnSubscribe(ctx context.Context, birthdayUserId uint64, subscribingUserId uint64) error {
	_, err := repo.db.ExecContext(ctx, `DELETE FROM birthday_subscribing WHERE birthday_user_id=$1 AND subscribing_user_id=$2`, birthdayUserId, subscribingUserId)
	if err != nil {
		return err
	}
	return nil
}

func (repo *RepoLayer) GetTodayBirthdayUsers(ctx context.Context, userId uint64) ([]*entity.User, error) {
	rows, err := repo.db.QueryContext(ctx, `SELECT u.id, u.name, u.surname, u.email, u.birthday FROM "user" AS u 
		JOIN birthday_subscribing AS bs ON u.id=bs.birthday_user_id 
		WHERE bs.subscribing_user_id=$1 AND u.birthday=CURRENT_DATE`, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []*entity.User{}, nil
		}
		return nil, err
	}
	users := []*entity.User{}
	for rows.Next() {
		user := entity.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Surname, &user.Email, &user.Birthday)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

//CURRENT_DATE
