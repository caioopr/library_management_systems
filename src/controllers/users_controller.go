package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

// CerateUser create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	if err = user.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
	}

	db, err := database.NewConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	usersRepo := repositories.NewUsersRepository(db)
	user.ID, err = usersRepo.Create(user)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// GetUsers gets all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get users ok"))
}

// GetUser gets a specific user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}

// UpdateUser updates a specific user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update user"))
}

// DeleteUser delete a specific user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete user"))
}
