package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CerateBook creates a new
func CreateBook(w http.ResponseWriter, r *http.Request) {
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
	if err = user.Prepare("register"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
	}

	db, err := database.NewConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	usersRepo := repositories.NewUsersRepository(db)
	user.ID, err = usersRepo.Create(user)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	user.Password = ""

	responses.JSON(w, http.StatusCreated, user)
}

// GetBooks gets all books with an exact title or with a title that contains the req. title
func GetBooks(w http.ResponseWriter, r *http.Request) {
	userNickname := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.NewConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	usersRepo := repositories.NewUsersRepository(db)
	users, err := usersRepo.GetUsersByNicknameLike(userNickname)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// GetBook gets a specific book
func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := auth.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != tokenUserID {
		responses.Error(w, http.StatusForbidden, errors.New("You cannot update an user that is not yours."))
		return
	}

	db, err := database.NewConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	usersRepo := repositories.NewUsersRepository(db)
	user, err := usersRepo.GetUserByID(userID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// UpdateBook updates a specific book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := auth.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != tokenUserID {
		responses.Error(w, http.StatusForbidden, errors.New("You cannot update an user that is not yours."))
		return
	}

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

	if err := user.Prepare("update"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
	}

	db, err := database.NewConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepo := repositories.NewUsersRepository(db)
	if err := userRepo.UpdateUser(userID, user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteBook deletes a specific book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := auth.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != tokenUserID {
		responses.Error(w, http.StatusForbidden, errors.New("You cannot update an user that is not yours."))
		return
	}

	db, err := database.NewConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepo := repositories.NewUsersRepository(db)
	if err := userRepo.DeleteUser(userID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
