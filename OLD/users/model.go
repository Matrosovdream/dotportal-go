package users

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}