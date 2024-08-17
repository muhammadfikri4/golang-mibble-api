package user

type UserBodyRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UsersDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
