package entity

type User struct {
	Id       uint64
	Name     string
	Surname  string
	BirthDay string
	Email    string
	Password string
}

type Session struct {
	Id    uint64
	Email string
}
