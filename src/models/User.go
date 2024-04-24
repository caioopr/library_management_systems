package models

import (
	"api/src/security"
	"api/src/validation"
	"errors"
	"strings"
	"time"
)

// User represents the user data object
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (user *User) validate(forType string) error {
	if user.Name == "" {
		return errors.New("The name field can not be empty.")
	}
	if user.Nickname == "" {
		return errors.New("The nickname field can not be empty.")
	}
	if user.Email == "" {
		return errors.New("The e-mail field can not be empty.")
	}

	if err := validation.ValidateFormat(user.Email); err != nil {
		return errors.New("The e-mail input has an invalid format.")
	}

	if forType == "register" && user.Password == "" {
		return errors.New("The password field can not be empty.")
	}

	return nil
}

func (user *User) format(prepType string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Email = strings.TrimSpace(user.Email)

	if prepType == "register" {
		hashedPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	return nil
}

// Prepare validates and formats the user obj
func (user *User) Prepare(prepType string) error {
	if err := user.validate(prepType); err != nil {
		return err
	}
	if err := user.format(prepType); err != nil {
		return err
	}

	return nil
}
