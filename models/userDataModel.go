package models

type LoginData struct {
	Store    string `json:"store"`
	Password string `json:"password"`
}

type SingUpData struct {
	Store    string `json:"store"`
	Password string `json:"password"`
}
