package types

type User struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type CreateUser struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
}
