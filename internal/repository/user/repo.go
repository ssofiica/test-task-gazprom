package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ssofiica/test-task-gazprom/internal/entity"
)

type Repo interface {
	GetAll(ctx context.Context) ([]*entity.User, error)
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
		err = rows.Scan(&user.Id, &user.Id, &user.Name, &user.Surname,
			&user.Email, &user.Birthday)
		if err != nil {
			return nil, err
		}
		fmt.Println(user.Birthday)
		users = append(users, &user)
	}
	return users, nil
}
