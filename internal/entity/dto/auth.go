package dto

type SignUp struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Birthday string `json:"birthday"`
}

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
