package entity

type User struct {
	Id       uint64
	Name     string
	Surname  string
	Birthday string
	Email    string
	Password []byte
}

type Session struct {
	Id    string
	Email string
}
