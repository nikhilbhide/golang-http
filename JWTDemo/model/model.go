package model

type Login struct {
	Email string `json:"email"`
	UserID string `json:"userID"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json:"error"`
}