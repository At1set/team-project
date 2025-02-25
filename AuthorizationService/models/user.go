package models

type User struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type ProfileResponse struct {
	Login string `json:"login"`
	Name  string `json:"name"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type RegisterRequest struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
