package models

type AuthenticatedUser struct {
	ID    string `json:"user_id"`
	Token string `json:"token"`
}
