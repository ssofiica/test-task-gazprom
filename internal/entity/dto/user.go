package dto

import "github.com/ssofiica/test-task-gazprom/internal/entity"

type User struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Birthday string `json:"birthday,omitempty"`
	Email    string `json:"email"`
}

func NewUser(user *entity.User) *User {
	return &User{
		Id:       user.Id,
		Name:     user.Name,
		Surname:  user.Surname,
		Birthday: user.Birthday,
		Email:    user.Email,
	}
}

func NewUserArray(users []*entity.User) []*User {
	if len(users) == 0 {
		return nil
	}
	usersDTO := make([]*User, len(users))
	for i, user := range users {
		usersDTO[i] = NewUser(user)
	}
	return usersDTO
}
