package entity

type User struct {
	Id       uint64
	Name     string
	Surname  string
	Birthday string
	Email    string
	Password string
}

type Session struct {
	Id    uint64
	Email string
}
